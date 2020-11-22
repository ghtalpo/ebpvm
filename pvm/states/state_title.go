package states

import (
	"image"
	"image/color"
	"log"

	"github.com/ghtalpo/ebpvm/pvm/db"
	"github.com/ghtalpo/ebpvm/pvm/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/juan-medina/goecs"
	"github.com/theckman/go-fsm"
)

// StateTitle implements State interface and for title menu
type StateTitle struct {
	m            *fsm.Machine
	world        *goecs.World
	logicChannel chan<- db.GameMessage
	uiChannel    <-chan db.GameMessage
	eids         []goecs.EntityID
}

// NewStateTitle is a constructor
func NewStateTitle(m *fsm.Machine, world *goecs.World, lch chan db.GameMessage, uch chan db.GameMessage) *StateTitle {
	s := StateTitle{
		m:            m,
		world:        world,
		logicChannel: lch,
		uiChannel:    uch,
	}

	return &s
}

// OnUpdate for updating UI
func (s *StateTitle) OnUpdate() error {
	return nil
}

// OnDraw for draw scene
func (s *StateTitle) OnDraw(screen *ebiten.Image) {
	drawBackground(s.world, screen)
	drawUI(s.world, screen)
}

// OnEnter for pre-process before entering state
func (s *StateTitle) OnEnter() {
	mem := db.GetMem()
	s.eids = append(s.eids,
		s.world.AddEntity(ui.NewBackground(
			0,
			0,
			ui.LoadImage("_resources/image/bg.png"),
		)),
		s.world.AddEntity(ui.NewButton(
			image.Rect(29*8, 9*8*2, (80-29)*8, (10*8+3)*2),
			db.GetResources().GetTopMenu(0),
			func(b *ui.Button) {
				if err := s.m.StateTransition("scenario"); err != nil {
					log.Fatal(err)
				}
			},
		)),
		s.world.AddEntity(ui.NewButton(
			image.Rect(29*8, 12*8*2, (80-29)*8, (13*8+3)*2),
			db.GetResources().GetTopMenu(1),
			func(b *ui.Button) {
				if err := s.m.StateTransition("load_game"); err != nil {
					log.Fatal(err)
				}
			},
		)),
		s.world.AddEntity(ui.NewText(
			40*8,
			6*8*2,
			mem.GetLocalized("eb_pvm"),
			color.White,
			true,
		)),
	)
}

// OnLeave for post-process after leaving state
func (s *StateTitle) OnLeave() {
	removeEntities(s.world, s.eids)
	s.eids = nil
}
