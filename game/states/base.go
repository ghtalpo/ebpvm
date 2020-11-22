package states

import (
	"fmt"
	"log"

	"github.com/ghtalpo/egb/common/ui"
	"github.com/ghtalpo/egb/game/db"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/juan-medina/goecs"
	"github.com/theckman/go-fsm"
)

// State is a interface for 5 essential functions for state transition
type State interface {
	OnUpdate() error
	OnDraw(screen *ebiten.Image)
	OnEnter()
	OnLeave()
}

// MakeState make
func MakeState(newState fsm.State, m *fsm.Machine, world *goecs.World, logicChannel chan db.GameMessage, uiChannel chan db.GameMessage) State {
	switch newState {
	case "loading":
		return NewStateLoading(m, world, logicChannel, uiChannel)
	case "title":
		return NewStateTitle(m, world, logicChannel, uiChannel)
	}
	panic(fmt.Sprintf("unhandled state: %s", newState))
}

// Common functions
func resetGeom(op *ebiten.DrawImageOptions) {
	resetGeomMirror(op, false)
}

func resetGeomMirror(op *ebiten.DrawImageOptions, bMirror bool) {
	op.GeoM.Reset()
	if bMirror {
		op.GeoM.Scale(-1, 1)
	} else {
		op.GeoM.Scale(1, 1)
	}
}

// setupGeom is reset, scale, translate geoM
func setupGeom(op *ebiten.DrawImageOptions, x int, y int) {
	setupGeomMirror(op, x, y, false)
}

// setupGeom ...
func setupGeomMirror(op *ebiten.DrawImageOptions, x int, y int, bMirror bool) {
	resetGeomMirror(op, bMirror)
	op.GeoM.Translate(float64(x), float64(y))
}

func drawUI(world *goecs.World, screen *ebiten.Image) {
	for it := world.Iterator(); it != nil; it = it.Next() {
		// get the values
		ent := it.Value()
		if ent.Contains(ui.TextType) {
			uiText := ent.Get(ui.TextType).(*ui.Text)
			uiText.Draw(screen)
		}
		if ent.Contains(ui.ButtonType) {
			uiButton := ent.Get(ui.ButtonType).(*ui.Button)
			uiButton.Draw(screen)
		}
	}
}

func removeEntities(world *goecs.World, eids []goecs.EntityID) {
	for _, eid := range eids {
		if err := world.Remove(eid); err != nil {
			log.Fatal(err)
		}
	}
}

func drawBackground(world *goecs.World, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	for it := world.Iterator(ui.BackgroundType); it != nil; it = it.Next() {
		// get the values
		ent := it.Value()
		uiBackground := ent.Get(ui.BackgroundType).(*ui.Background)
		uiBackground.Draw(screen, op)
	}
}
