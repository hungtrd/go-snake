package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var goph *ebiten.Image

func init() {
	var err error
	goph, _, err = ebitenutil.NewImageFromFile("vim-go.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	Speed time.Duration
	Board Board
}

func (g *Game) Update() error {
	time.Sleep(g.Speed)

	if ebiten.IsKeyPressed(ebiten.KeyUp) && g.Board.Snake.Dir.y != 1 {
		g.Board.Snake.ChangeDirection(Point{0, -1})
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) && g.Board.Snake.Dir.y != -1 {
		g.Board.Snake.ChangeDirection(Point{0, 1})
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) && g.Board.Snake.Dir.x != 1 {
		g.Board.Snake.ChangeDirection(Point{-1, 0})
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) && g.Board.Snake.Dir.x != -1 {
		g.Board.Snake.ChangeDirection(Point{1, 0})
	}

	if g.Board.Snake.ReachFood(g.Board.Food) {
		g.Board.Snake.Eat(g.Board.Food)
		g.Board.PlaceFood()
		g.Board.Snake.Move()
	} else {
		g.Board.Snake.Move()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, dot := range g.Board.Snake.Dots {
		g.Board.DrawTile(screen, dot, goph)
	}
	g.Board.DrawTile(screen, g.Board.Food, goph)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Board.WidthSize(), g.Board.HeightSize()
}

func main() {
	game := &Game{
		Speed: 100 * time.Millisecond,
		Board: Board{
			Snake:  NewSnake(),
			Food:   Point{3, 4},
			Width:  20,
			Height: 20,
			Tile:   36,
		},
	}
	ebiten.SetWindowSize(game.Board.WidthSize(), game.Board.HeightSize())
	ebiten.SetWindowTitle("Snake with ebitengine")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
