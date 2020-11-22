package pvm

import (
	"fmt"
	"log"
	"math/rand"

	"time"

	"github.com/ghtalpo/ebpvm/pvm/db"
	"github.com/ghtalpo/ebpvm/pvm/states"
	"github.com/ghtalpo/ebpvm/pvm/ui"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/juan-medina/goecs"
	"github.com/theckman/go-fsm"
	"github.com/tkanos/gonfig"
)

// GameState indicates game status, this can be global db
type GameState struct {
	M     *fsm.Machine
	world *goecs.World
}

// NewGameState is a constructor for Game
func NewGameState() *GameState {
	return &GameState{
		M:     &fsm.Machine{},
		world: goecs.Default(),
	}
}

// StateTransitionCallback is a callback from fsm
// used for calling MakeState & OnEnter
// WARNING: does not called in 1st state transition
func (*GameState) StateTransitionCallback(newState fsm.State) error {
	if currentState != nil {
		currentState.OnLeave()
	}
	currentState = states.MakeState(newState, gameState.M, gameState.world, logicChannel, uiChannel)
	if currentState != nil {
		currentState.OnEnter()
	}
	return nil
}

// Configuration says
type Configuration struct {
	Language string
}

var (
	gameState     = NewGameState()
	currentState  states.State
	logicChannel  chan db.GameMessage = make(chan db.GameMessage, 1)
	uiChannel     chan db.GameMessage = make(chan db.GameMessage, 1)
	configuration Configuration
)

// Start is
func Start() {
	configuration = Configuration{}

	if err := gonfig.GetConf("config.yaml", &configuration); err != nil {
		panic(err)
	}

	db.SetGameLanguage(configuration.Language)

	fmt.Println("pvm.init")

	rand.Seed(time.Now().UnixNano())

	db.LoadDataFiles()

	db.GetMem().Initialize(logicChannel, uiChannel)
	// db.GetSound().Initialize()
	db.GetResources().Initialize()

	gameState.world.AddSystem(updateButtonSystem)

	// add initial rule
	gameState.M.SetStateTransitionCallback(gameState, false)

	if err := gameState.M.AddStateTransitionRules("start", "loading", "title", "exit"); err != nil {
		log.Fatal(err)
	}

	// add rest of rules
	gameState.M.AddStateTransitionRules("start", "loading")
	gameState.M.AddStateTransitionRules("loading", "title")
	gameState.M.AddStateTransitionRules("title", "exit")
	gameState.M.AddStateTransitionRules("title", "scenario")
	gameState.M.AddStateTransitionRules("title", "load_game")
	gameState.M.AddStateTransitionRules("load_game", "title")
	gameState.M.AddStateTransitionRules("load_game", "play")
	gameState.M.AddStateTransitionRules("scenario", "select_player")
	gameState.M.AddStateTransitionRules("scenario", "title")
	gameState.M.AddStateTransitionRules("select_player", "select_new_player")
	gameState.M.AddStateTransitionRules("select_new_player", "select_modes")
	gameState.M.AddStateTransitionRules("select_player", "select_modes")
	gameState.M.AddStateTransitionRules("select_modes", "exit")
	gameState.M.AddStateTransitionRules("select_modes", "scenario")
	gameState.M.AddStateTransitionRules("select_modes", "select_player")
	gameState.M.AddStateTransitionRules("select_modes", "play")
	gameState.M.AddStateTransitionRules("play", "exit")
	gameState.M.AddStateTransitionRules("exit") // final state

	// set initial state
	if err := gameState.M.StateTransition("start"); err != nil {
		log.Fatal(err)
	}

	gameState.M.StateTransition("loading")
}

// Stop is a api for stopping pvm logic
func Stop() {
	fmt.Println("pvm.Stop called")

	select {
	case logicChannel <- db.GameMessage{"quit": nil}:
		fmt.Println("[GM] sent stop message")
	default:
		fmt.Println("[GM] no message sent")
	}
	close(logicChannel)
	close(uiChannel)
	fmt.Println("channel closed")
}

// OnUpdate called from main.update.
func OnUpdate() error {
	gameState.world.Update(1.0 / float32(ebiten.MaxTPS()))
	return currentState.OnUpdate()
}

// OnDraw called from main.draw.
func OnDraw(screen *ebiten.Image) {
	currentState.OnDraw(screen)
}

func updateButtonSystem(world *goecs.World, delta float32) error {
	for it := world.Iterator(ui.ButtonType); it != nil; it = it.Next() {
		// get the values
		ent := it.Value()
		uiButton := ent.Get(ui.ButtonType).(*ui.Button)
		uiButton.Update()
	}

	return nil
}
