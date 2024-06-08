package game

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

type Bullet struct {
	Xpos     int32
	Ypos     int32
	velocity int32
	radius   float32
	Draw     bool
	Color    raylib.Color
}

func NewBullet(g *Game) *Bullet {
	return &Bullet{
		Xpos:     g.xCoords + 50,
		Ypos:     g.yCoords + 25,
		velocity: 5,
		radius:   10,
		Draw:     true,
		Color:    raylib.White,
	}
}
