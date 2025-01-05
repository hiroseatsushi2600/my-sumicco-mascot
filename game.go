package main

import (
	"fmt"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	mascot *Mascot
	skipDraw bool
	introShown bool
}

func (g *Game) Update() error {
	// ゲーム終了
	if g.isQuitRequested() {
		return ebiten.Termination
	}
	if g.isMainActionRequested() {
		g.mascot.ChangePosition()
		g.skipDraw = false
	} else if g.getChangeScaleRequest() == 1 {
		g.mascot.Bigger()
		g.skipDraw = false
	} else if g.getChangeScaleRequest() == -1 {
		g.mascot.Smaller()
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
	slog.Debug("draw", "skipDraw", fmt.Sprint(g.skipDraw))
	// ebiten.SetScreenClearedEveryFrame(false) であるためここで明示的にクリアする
	screen.Clear()
	g.mascot.draw(screen)
	g.introShown = true
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.mascot.mascotR.Bounds().Dx(), g.mascot.mascotR.Bounds().Dy()
}

func (g *Game) isMainActionRequested() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		return true
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return true
	}
	return false
}

func (g *Game) isQuitRequested() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) || inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return true
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		return true
	}
	return false
}

func (g *Game) getChangeScaleRequest() int {
	if inpututil.IsKeyJustPressed(ebiten.KeyB)  {
		return 1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		return -1
	}
	return 0
}
