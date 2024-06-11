package game

import (
	"errors"
	"time"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

var (
	errImage   = errors.New("failed to load image")
	errTexture = errors.New("failed to load texture")
)

type Game struct {
	SCREEN_WIDTH  int32
	SCREEN_HEIGHT int32
	Ship          *Ship
	Enemy         *Enemy
	Bullets       []*Bullet
	Enemies       []*Enemy
	LastShot      time.Time
}

func NewGame() *Game {
	return &Game{
		SCREEN_WIDTH:  800,
		SCREEN_HEIGHT: 600,
		Bullets:       []*Bullet{},
		Enemies:       []*Enemy{},
	}

}

func (g *Game) StartWindow() error {
	raylib.InitWindow(g.SCREEN_WIDTH, g.SCREEN_HEIGHT, "space invaders")

	defer raylib.CloseWindow()

	g.Ship = NewShip() // the ship has to be initialized after the window opens
	g.CreateEnemy()

	raylib.SetTargetFPS(60)

	for !raylib.WindowShouldClose() { // as long as the window is open
		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.Black)

		raylib.DrawTexture(g.Ship.Image, g.Ship.Xpos, g.Ship.Ypos, raylib.White) // drawing our ship
		// raylib.DrawTexture(g.Enemy.EnemyUp, g.Enemy.Xpos, g.Enemy.Ypos, raylib.Blue) // drawing the enemy (test)

		g.DrawEnemy()
		// ship movement
		g.Ship.moveShip(g.SCREEN_WIDTH)
		// ship bullets
		g.FireBullet(g.Ship)
		g.DrawBullet()

		raylib.EndDrawing()
	}

	return nil
}
