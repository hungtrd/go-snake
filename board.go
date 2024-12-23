package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	Snake  *Snake
	Food   Point
	Width  int
	Height int
	Tile   int
}

func (b *Board) DrawTile(screen *ebiten.Image, x, y int, image *ebiten.Image) {
	geom := ebiten.GeoM{}
	geom.Translate(float64(x*b.Tile), float64(y*b.Tile))
	options := &ebiten.DrawImageOptions{
		GeoM: geom,
	}
	screen.DrawImage(image, options)
}

func (b *Board) PlaceFood() {
	for {
		b.Food = Point{rand.Intn(b.Width), rand.Intn(b.Height)}

		valid := true
		for _, dot := range b.Snake.Dots {
			if dot == b.Food {
				valid = false
			}
		}

		if valid {
			return
		}
	}
}

func (b *Board) WidthSize() int {
	return b.Width * b.Tile
}

func (b *Board) HeightSize() int {
	return b.Height * b.Tile
}
