package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"log/slog"
)

type MascotPosition int

const (
	TopLeft MascotPosition = iota
	TopRight
	BottomLeft
	BottomRight
)

type Mascot struct {
	mascotR        *ebiten.Image
	mascotL        *ebiten.Image
	characterScale float64
	position       MascotPosition
	winPosX        int
	winPosY        int
}

func NewMascot(mascotR *ebiten.Image, mascotL *ebiten.Image) *Mascot {

	return &Mascot{
		mascotR:        mascotR,
		mascotL:        mascotL,
		characterScale: 0.25,
		position:       BottomLeft,
		winPosX:        0,
		winPosY:        0,
	}
}

func (m *Mascot) update() error {
	mascotWidth := int(float64(m.mascotR.Bounds().Dx()) * m.characterScale)
	mascotHeight := int(float64(m.mascotR.Bounds().Dy()) * m.characterScale)
	monitorWidth, monitorHeight := ebiten.Monitor().Size()
	slog.Debug("mascot w, h, monitor w, h", "mas w", fmt.Sprint(mascotWidth), "mas h", fmt.Sprint(mascotHeight), "mon w", fmt.Sprint(monitorWidth), "mon h", fmt.Sprint(monitorHeight))

	switch m.position {
	case TopLeft:
		m.winPosX = 0
		m.winPosY = 0
	case TopRight:
		m.winPosX = monitorWidth - mascotWidth
		m.winPosY = 0
	case BottomLeft:
		m.winPosX = 0
		m.winPosY = monitorHeight - mascotHeight
	case BottomRight:
		m.winPosX = monitorWidth - mascotWidth
		m.winPosY = monitorHeight - mascotHeight
	}

	ebiten.SetWindowPosition(m.winPosX, m.winPosY)
	return nil
}

func (m *Mascot) draw(screen *ebiten.Image) {
	var img *ebiten.Image
	switch m.position {
	case TopLeft, BottomLeft:
		img = m.mascotL
	case TopRight, BottomRight:
		img = m.mascotR
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(m.characterScale, m.characterScale)
	x, y := getZure(m.position, img.Bounds().Dx(), img.Bounds().Dy(), m.characterScale)
	op.GeoM.Translate(x, y)
	screen.DrawImage(img, op)
}

func (m *Mascot) ChangePosition() {
	m.position = (m.position + 1) % 4
}

func (m *Mascot) Bigger() {
	m.characterScale += 0.05
}

func (m *Mascot) Smaller() {
	m.characterScale -= 0.05
}

func getZure(pos MascotPosition, x int, y int, scale float64) (float64, float64) {
	hr := 0.4 * scale
	switch pos {
	case TopLeft, BottomLeft:
		return -float64(x) * hr, 0
	case TopRight, BottomRight:
		return float64(x) * hr, 0
	}
	return 0, 0
}
