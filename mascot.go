package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"log/slog"
)

type Mascot struct {
	input          *Input
	monitor        *Monitor
	mascotR        *ebiten.Image
	mascotL        *ebiten.Image
	characterScale float64
	winPosX        int
	winPosY        int
	winPosLR       LR
}

type LR int

const (
	L LR = iota
	R
)

const movePower = 16

func NewMascot(input *Input, mascotR *ebiten.Image, mascotL *ebiten.Image) *Mascot {
	return &Mascot{
		input:          input,
		monitor:        NewMonitor(),
		mascotR:        mascotR,
		mascotL:        mascotL,
		characterScale: 0.25,
		winPosX:        0,
		winPosY:        0,
		winPosLR:       L,
	}
}

func (m *Mascot) update() error {
	mascotWidth := int(float64(m.mascotR.Bounds().Dx()) * m.characterScale)
	mascotHeight := int(float64(m.mascotR.Bounds().Dy()) * m.characterScale)
	monitorWidth, monitorHeight := ebiten.Monitor().Size()

	if m.input.GetRequest() != None {
		slog.Debug("input request", "req", fmt.Sprint(m.input.GetRequest()))
	}
	if m.input.GetRequest() == MoveUp && m.winPosY > 0 {
		m.winPosY -= movePower
	}
	if m.input.GetRequest() == MoveDown && m.winPosY < monitorHeight-mascotHeight {
		m.winPosY += movePower
	}
	if m.input.GetRequest() == MoveLeft {
		if m.winPosLR == R {
			m.winPosLR = L
		} else {
			m.monitor.PreviousMonitor()
		}
	}
	if m.input.GetRequest() == MoveRight {
		if m.winPosLR == L {
			m.winPosLR = R
		} else {
			m.monitor.NextMonitor()
		}
	}
	if m.input.GetRequest() == Avoid {
		if m.winPosLR == R {
			m.winPosLR = L
		} else {
			m.winPosLR = R
		}
	}
	if m.input.GetRequest() == ScaleUp {
		if m.characterScale < 1.35 {
			m.characterScale += 0.05
		}
	}
	if m.input.GetRequest() == ScaleDown {
		if m.characterScale > 0.1 {
			m.characterScale -= 0.05
		}
	}

	// x座標は原則LRによって決まる
	if m.winPosLR == L {
		m.winPosX = 0
	} else {
		m.winPosX = monitorWidth - mascotWidth
	}

	slog.Debug("size", "m.characterScale", fmt.Sprint(m.characterScale))
	ebiten.SetWindowPosition(m.winPosX, m.winPosY)
	ebiten.SetMonitor(m.monitor.currentMonitor)
	return nil
}

func (m *Mascot) draw(screen *ebiten.Image) {
	var img *ebiten.Image
	if m.winPosLR == L {
		img = m.mascotL
	} else {
		img = m.mascotR
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(m.characterScale, m.characterScale)
	x, y := calcMerikomi(m.winPosLR, img.Bounds().Dx(), m.characterScale)
	op.GeoM.Translate(x, y)
	screen.DrawImage(img, op)
}

func calcMerikomi(winPosLR LR, x int, scale float64) (float64, float64) {
	// マスコットが体半分ほどを画面端に隠す座標を計算する
	hr := 0.4 * scale
	switch winPosLR {
	case L:
		return -float64(x) * hr, 0
	case R:
		return float64(x) * hr, 0
	}
	return 0, 0
}
