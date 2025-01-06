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
	// slog.Debug("mascot w, h, monitor w, h", "mas w", fmt.Sprint(mascotWidth), "mas h", fmt.Sprint(mascotHeight), "mon w", fmt.Sprint(monitorWidth), "mon h", fmt.Sprint(monitorHeight))

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
		m.characterScale += 0.05
	}
	if m.input.GetRequest() == ScaleDown {
		m.characterScale -= 0.05
	}

	// x座標は原則LRによって決まる
	if m.winPosLR == L {
		m.winPosX = 0
	} else {
		m.winPosX = monitorWidth - mascotWidth
	}

	ebiten.SetWindowPosition(m.winPosX, m.winPosY)
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
	x, y := calcMerikomi(m.winPosLR, img.Bounds().Dx(), img.Bounds().Dy(), m.characterScale)
	op.GeoM.Translate(x, y)
	screen.DrawImage(img, op)
}

func calcMerikomi(winPosLR LR, x int, y int, scale float64) (float64, float64) {
	// マスコットが体半分ほどを画面端に隠す座標を計算する
	hr := 0.4 * scale
	switch winPosLR {
	case L:
		return -float64(x) * hr, float64(y) * hr
	case R:
		return float64(x) * hr, float64(y) * hr
	}
	return 0, 0
}
