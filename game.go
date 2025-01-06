package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	mascot     *Mascot
	skipDraw   bool
	introShown bool
	input      *Input
}

func NewGame(mascot *Mascot, input *Input) *Game {
	return &Game{
		mascot:     mascot,
		skipDraw:   false,
		introShown: false,
		input:      input,
	}
}

func (g *Game) Update() error {
	// ゲーム終了
	if g.input.GetRequest() == Quit {
		return ebiten.Termination
	}
	if g.input.GetRequest() != 0 {
		g.skipDraw = false
	} else {
		g.skipDraw = true
	}
	return g.mascot.update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.skipDraw && g.introShown {
		return
	}
	// ebiten.SetScreenClearedEveryFrame(false) であるためここで明示的にクリアする
	screen.Clear()
	g.mascot.draw(screen)
	g.introShown = true
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.mascot.mascotR.Bounds().Dx(), g.mascot.mascotR.Bounds().Dy()
}
