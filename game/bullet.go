package game

import (
	"math/rand"
	"time"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const playerShotDelay = 500 * time.Millisecond
const enemyShotDelay = 500 * time.Millisecond

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

	if raylib.IsKeyDown(raylib.KeySpace) && time.Since(g.LastPlayerShot) >= playerShotDelay {
		current_bullet := NewBulletFromShip(s)
		g.PlayerBullets = append(g.PlayerBullets, current_bullet)
		g.LastPlayerShot = time.Now()
	}
}

func (g *Game) FireEnemyBullet(e *Enemy) {
	chance := rand.Intn(100) + 1

	if chance <= 5 && time.Since(g.LastEnemyShot) >= enemyShotDelay {
		current_bullet := NewBulletFromEnemy(e)
		g.EnemyBullets = append(g.EnemyBullets, current_bullet)
		g.LastEnemyShot = time.Now()
	}
}

func (g *Game) CheckBulletEnemyCollision() {
	i := 0
	for i < len(g.PlayerBullets) {
		bullet := g.PlayerBullets[i]
		collision := false
		j := 0
		for j < len(g.Enemies) {
			enemy := g.Enemies[j]
			if raylib.CheckCollisionRecs(raylib.NewRectangle(float32(bullet.Xpos), float32(bullet.Ypos), bullet.radius, bullet.radius), enemy.Rectangle) {
				// collision detected, decrease health and handle the score
				g.Enemies[j].Health -= 1
				g.Score++
				if g.Enemies[j].Health <= 0 {
					// remove the enemy
					g.Enemies = append(g.Enemies[:j], g.Enemies[j+1:]...)
				} else {
					j++
				}
				// remove the bullet
				g.PlayerBullets = append(g.PlayerBullets[:i], g.PlayerBullets[i+1:]...)
				collision = true
				break
			} else {
				j++
			}
		}
		if !collision {
			i++
		}
	}
}

func (g *Game) CheckBulletShipCollision() {
	i := 0
	for i < len(g.EnemyBullets) {
		bullet := g.EnemyBullets[i]
		if raylib.CheckCollisionRecs(raylib.NewRectangle(float32(bullet.Xpos), float32(bullet.Ypos), bullet.radius, bullet.radius), g.Ship.Rectangle) {
			// collision detected, decrease health
			g.Ship.Health -= 1
			// remove the bullet (otherwise it will keep decreasing health in this space)
			g.EnemyBullets = append(g.EnemyBullets[:i], g.EnemyBullets[i+1:]...)
		} else {
			i++
		}
	}
}

func (g *Game) DrawBullet() {
	// drawing the player's bullet
	for index1, current_bullet := range g.PlayerBullets {
		if current_bullet.Draw {
			g.PlayerBullets[index1].Ypos -= current_bullet.velocity
			raylib.DrawCircle(current_bullet.Xpos-16, current_bullet.Ypos, current_bullet.radius, current_bullet.Color)

			// if the bullets are out of the screen, stop drawing them
			if current_bullet.Ypos < 0 || current_bullet.Ypos > g.SCREEN_HEIGHT {
				g.PlayerBullets[index1].Draw = false
			}
		}
	}

	// Check for bullet-enemy collisions after all bullets have been drawn
	g.CheckBulletEnemyCollision()

	// drawing the enemy's bullets
	for index1, current_bullet := range g.EnemyBullets {
		if current_bullet.Draw {
			g.EnemyBullets[index1].Ypos -= current_bullet.velocity
			raylib.DrawCircle(current_bullet.Xpos-16, current_bullet.Ypos, current_bullet.radius, current_bullet.Color)

			// if the bullets are out of the screen, stop drawing them
			if current_bullet.Ypos < 0 || current_bullet.Ypos > g.SCREEN_HEIGHT {
				g.EnemyBullets[index1].Draw = false
			}
		}
	}

	// Check for bullet-ship collisions after all bullets have been drawn
	g.CheckBulletShipCollision()
}
