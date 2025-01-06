package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Input struct{}

func NewInput() *Input {
	return &Input{}
}

type Request int

const (
	Quit Request = iota
	MoveUp
	MoveDown
	MoveLeft
	MoveRight
	ScaleUp
	ScaleDown
	Avoid
	None
)

func (i *Input) GetRequest() Request {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) || inpututil.IsKeyJustPressed(ebiten.KeyEscape) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		return Quit
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		return MoveUp
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		return MoveDown
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		return MoveLeft
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		return MoveRight
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyB) {
		return ScaleUp
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyV) {
		return ScaleDown
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return Avoid
	}
	return None
}
