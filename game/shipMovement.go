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

	if raylib.IsKeyDown(raylib.KeySpace) {
		current_bullet := NewBullet(g)

		g.Bullets = append(g.Bullets, current_bullet)
	}

	// drawing the bullet
	for index1, current_bullet := range g.Bullets {
		if current_bullet.Draw {
			g.Bullets[index1].Ypos = g.Bullets[index1].Ypos - current_bullet.velocity
			raylib.DrawCircle(current_bullet.Xpos-16, current_bullet.Ypos, current_bullet.radius, current_bullet.Color)

			// if the bullets are out of the screen, stop drawing them
			if current_bullet.Ypos < 0 || current_bullet.Ypos > g.SCREEN_HEIGHT {
				g.Bullets[index1].Draw = false
			}
		}
	}
}
