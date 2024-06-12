package game

import (
	"math/rand"
	"time"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const shotDelay = 500 * time.Millisecond

type Bullet struct {
	Xpos     int32
	Ypos     int32
	velocity int32
	radius   float32
	Draw     bool
	Color    raylib.Color
}

func NewBulletFromShip(s *Ship) *Bullet {
	return &Bullet{
		Xpos:     s.Xpos + 50,
		Ypos:     s.Ypos + 25,
		velocity: 5,
		radius:   7,
		Draw:     true,
		Color:    raylib.White,
	}
}

func NewBulletFromEnemy(e *Enemy) *Bullet {
	return &Bullet{
		Xpos:     e.Xpos + 50,
		Ypos:     e.Ypos + 25,
		velocity: -5, // negative so the bullet moves down
		radius:   7,
		Draw:     true,
		Color:    raylib.Red,
	}
}

func (g *Game) FireBullet(s *Ship) {

	if raylib.IsKeyDown(raylib.KeySpace) && time.Since(g.LastShot) >= shotDelay {
		current_bullet := NewBulletFromShip(s)
		g.Bullets = append(g.Bullets, current_bullet)
		g.LastShot = time.Now()
	}
}

func (g *Game) FireEnemyBullet(e *Enemy) {
	chance := rand.Intn(100) + 1

	if chance <= 5 && time.Since(g.LastShot) >= shotDelay {
		current_bullet := NewBulletFromEnemy(e)
		g.Bullets = append(g.Bullets, current_bullet)
		g.LastShot = time.Now()
	}
}

func (g *Game) DrawBullet() {
	// drawing the bullet
	for index1, current_bullet := range g.Bullets {
		if current_bullet.Draw {
			g.Bullets[index1].Ypos -= current_bullet.velocity
			raylib.DrawCircle(current_bullet.Xpos-16, current_bullet.Ypos, current_bullet.radius, current_bullet.Color)

			// if the bullets are out of the screen, stop drawing them
			if current_bullet.Ypos < 0 || current_bullet.Ypos > g.SCREEN_HEIGHT {
				g.Bullets[index1].Draw = false
			}
		}
	}
}
