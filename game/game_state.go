package game

import (
	"fmt"
	"log"

	"time"

	"github.com/ghtalpo/egb/common/mathutil"
	"github.com/ghtalpo/egb/common/ui"
	"github.com/ghtalpo/egb/game/db"
	"github.com/ghtalpo/egb/game/states"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/juan-medina/goecs"
	"github.com/theckman/go-fsm"
	"github.com/tkanos/gonfig"
)

// status indicates game status, this can be global db
type status struct {
	M             *fsm.Machine
	world         *goecs.World
	currentState  states.State
	configuration Configuration
	logicChannel  chan db.GameMessage
	uiChannel     chan db.GameMessage
}

// newstatus is a constructor for Game
func newstatus() *status {
	return &status{
		M:            &fsm.Machine{},
		world:        goecs.Default(),
		logicChannel: make(chan db.GameMessage, 1),
		uiChannel:    make(chan db.GameMessage, 1),
	}
}

// StateTransitionCallback is a callback from fsm
// used for calling MakeState & OnEnter
// WARNING: does not called in 1st state transition
func (s *status) StateTransitionCallback(newState fsm.State) error {
	if s.currentState != nil {
		s.currentState.OnLeave()
	}
	s.currentState = states.MakeState(newState, gs.M, gs.world, s.logicChannel, s.uiChannel)
	if s.currentState != nil {
		s.currentState.OnEnter()
	}
	return nil
}

var gs = newstatus()

// Configuration says
type Configuration struct {
	Language string
}

// Start ...
func Start() {
	gs.configuration = Configuration{}

	if err := gonfig.GetConf("config.yaml", &gs.configuration); err != nil {
		panic(err)
	}

	db.SetGameLanguage(gs.configuration.Language)

	fmt.Println("egb.init")

	mathutil.RandSeed(time.Now().UnixNano())

	db.LoadDataFiles()

	db.GetMem().Initialize(gs.logicChannel, gs.uiChannel)
	// db.GetSound().Initialize()
	db.GetResources().Initialize()

	gs.world.AddSystem(updateButtonSystem)

	// add initial rule
	gs.M.SetStateTransitionCallback(gs, false)

	if err := gs.M.AddStateTransitionRules("start", "loading", "title", "exit"); err != nil {
		log.Fatal(err)
	}

	// add rest of rules
	gs.M.AddStateTransitionRules("start", "loading")
	gs.M.AddStateTransitionRules("loading", "title")
	gs.M.AddStateTransitionRules("title", "exit")
	gs.M.AddStateTransitionRules("title", "scenario")
	gs.M.AddStateTransitionRules("title", "load_game")
	gs.M.AddStateTransitionRules("load_game", "title")
	gs.M.AddStateTransitionRules("load_game", "play")
	gs.M.AddStateTransitionRules("scenario", "select_player")
	gs.M.AddStateTransitionRules("scenario", "title")
	gs.M.AddStateTransitionRules("select_player", "select_new_player")
	gs.M.AddStateTransitionRules("select_new_player", "select_modes")
	gs.M.AddStateTransitionRules("select_player", "select_modes")
	gs.M.AddStateTransitionRules("select_modes", "exit")
	gs.M.AddStateTransitionRules("select_modes", "scenario")
	gs.M.AddStateTransitionRules("select_modes", "select_player")
	gs.M.AddStateTransitionRules("select_modes", "play")
	gs.M.AddStateTransitionRules("play", "exit")
	gs.M.AddStateTransitionRules("exit") // final state

	// set initial state
	if err := gs.M.StateTransition("start"); err != nil {
		log.Fatal(err)
	}

	gs.M.StateTransition("loading")
}

// Stop is a api for stopping egb logic
func Stop() {
	fmt.Println("egb.Stop called")

	select {
	case gs.logicChannel <- db.GameMessage{"quit": nil}:
		fmt.Println("[GM] sent stop message")
	default:
		fmt.Println("[GM] no message sent")
	}
	close(gs.logicChannel)
	close(gs.uiChannel)
	fmt.Println("channel closed")
}

// OnUpdate called from main.update.
func OnUpdate() error {
	gs.world.Update(1.0 / float32(ebiten.TPS()))
	return gs.currentState.OnUpdate()
}

// OnDraw called from main.draw.
func OnDraw(screen *ebiten.Image) {
	gs.currentState.OnDraw(screen)
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
