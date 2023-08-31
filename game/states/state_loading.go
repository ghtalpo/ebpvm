package states

import (
	"fmt"
	"image/color"
	"log"

	"github.com/ghtalpo/egb/common/ui"
	"github.com/ghtalpo/egb/game/db"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/juan-medina/goecs"
	"github.com/theckman/go-fsm"
)

const (
	maxCount int = 2
)

// StateLoading implements State interface and for loading scene
type StateLoading struct {
	m            *fsm.Machine
	world        *goecs.World
	count        int
	logicChannel chan<- db.GameMessage
	uiChannel    <-chan db.GameMessage
	eids         []goecs.EntityID
	tick         int
}

// NewStateLoading is a constructor
func NewStateLoading(m *fsm.Machine, world *goecs.World, lch chan db.GameMessage, uch chan db.GameMessage) *StateLoading {
	return &StateLoading{
		m:            m,
		world:        world,
		count:        maxCount,
		logicChannel: lch,
		uiChannel:    uch,
		tick:         ebiten.TPS(),
	}
}

// OnUpdate for updating UI
func (s *StateLoading) OnUpdate() error {
	s.tick++
	if ebiten.TPS() <= s.tick {
		s.tick = 0
		switch s.count {
		case 2:
			db.CheckDatFiles()
			s.countDown()
		case 1:
			// do something else
			s.countDown()
		case 0:
			if err := s.m.StateTransition("title"); err != nil {
				log.Fatal(err)
			}
		}
	}

	return nil
}

// OnDraw for draw scene
func (s *StateLoading) OnDraw(screen *ebiten.Image) {
	drawUI(s.world, screen)
}

// OnEnter for pre-process before entering state
func (s *StateLoading) OnEnter() {
	s.eids = append(s.eids,
		s.world.AddEntity(ui.NewText(
			35*8,
			12*8*2,
			fmt.Sprintf(db.GetMem().GetLocalized("now_loading_d"), s.count),
			color.White,
			true,
		)),
	)
}

// OnLeave for post-process after leaving state
func (s *StateLoading) OnLeave() {
	removeEntities(s.world, s.eids)
	s.eids = nil
}

func (s *StateLoading) countDown() {
	s.count--
	eidText := s.eids[0]
	entity := s.world.Get(eidText)
	text := entity.Get(ui.TextType).(*ui.Text)
	text.Text = fmt.Sprintf(db.GetMem().GetLocalized("now_loading_d"), s.count)
}
