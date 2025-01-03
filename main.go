package main

import (
	"log"
	
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	res := NewResource()

	ebiten.SetWindowSize(res.GetMascotR().Bounds().Dx(), res.mascotR.Bounds().Dy())
	ebiten.SetWindowTitle("human lost mascot2")
	ebiten.SetWindowDecorated(false)
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetWindowFloating(true)
	ebiten.SetTPS(6)

	op := &ebiten.RunGameOptions{
		ScreenTransparent: true,
	}

	game := &Game{
		mascot: NewMascot(res.GetMascotR(), res.GetMascotL()),
	}

	if err := ebiten.RunGameWithOptions(game, op); err != nil {
		log.Fatal(err)
	}
}
