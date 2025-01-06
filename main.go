package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	res := NewResource()

	ebiten.SetWindowSize(res.GetMascotR().Bounds().Dx(), res.mascotR.Bounds().Dy())
	ebiten.SetWindowTitle("sumicco mascot")
	ebiten.SetWindowDecorated(false)
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetWindowFloating(true)
	ebiten.SetTPS(8)
	ebiten.SetScreenClearedEveryFrame(false)
	if len(os.Args) > 1 {
        switch os.Args[1] {
        case "-d":
            slog.SetLogLoggerLevel(slog.LevelDebug)
        }
    }
	slog.Info("start", "args", os.Args)
	slog.Debug("debug")

	op := &ebiten.RunGameOptions{
		ScreenTransparent: true,
		// 奥ゆかしさ重点
		InitUnfocused: true,
		SkipTaskbar: true,
	}

    input := NewInput()
    game := NewGame(NewMascot(input, res.GetMascotR(), res.GetMascotL()), input)

    if err := ebiten.RunGameWithOptions(game, op); err != nil {
		log.Fatal(err)
	}
}
