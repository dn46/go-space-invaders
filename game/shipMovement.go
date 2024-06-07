package game

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) moveShip() {
	if raylib.IsKeyDown(raylib.KeyD) {
		if g.Ship.Width+g.xCoords+5 < g.SCREEN_WIDTH {
			g.xCoords += 5
		}
	}

	if raylib.IsKeyDown(raylib.KeyA) {
		if g.xCoords-5 > 0 {
			g.xCoords -= 5
		}
	}
}
