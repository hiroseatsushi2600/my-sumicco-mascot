package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	mascot *Mascot
}

func (g *Game) Update() error {
	if g.isQuit() {
		return ebiten.Termination
	}
	return g.mascot.update(g.iSKeyPressed())
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.mascot.draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.mascot.mascotR.Bounds().Dx(), g.mascot.mascotR.Bounds().Dy()
}

func (g *Game) iSKeyPressed() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		return true
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return true
	}
	return false
}

func (g *Game) isQuit() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) || inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return true
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		return true
	}
	return false
}
