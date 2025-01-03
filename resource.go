package main

import (
	_ "image/png"
	"log"
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	images "sumicco-mascot/assets/images"
)

type Resource struct {
	mascotR *ebiten.Image
	mascotL *ebiten.Image
}

func NewResource() *Resource {

	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(images.GopherR_png))
	if err != nil {
		log.Fatal(err)
	}
	mascot1 := ebiten.NewImageFromImage(img)
	img, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(images.GopherL_png))
	if err != nil {
		log.Fatal(err)
	}
	mascot2 := ebiten.NewImageFromImage(img)

	return &Resource{
		mascotR: mascot1,
		mascotL: mascot2,
	}
}

func (r *Resource) GetMascotR() *ebiten.Image {
	return r.mascotR
}

func (r *Resource) GetMascotL() *ebiten.Image {
	return r.mascotL
}
