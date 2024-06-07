package game

import (
	"errors"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

var (
	errImage   = errors.New("failed to load image")
	errTexture = errors.New("failed to load texture")
)

type Game struct {
	SCREEN_WIDTH  int32
	SCREEN_HEIGHT int32
	xCoords       int32
	yCoords       int32
	Ship          raylib.Texture2D
}

func NewGame() *Game {
	return &Game{
		SCREEN_WIDTH:  800,
		SCREEN_HEIGHT: 600,
		xCoords:       2,
		yCoords:       500,
	}
}

func (g *Game) StartWindow() error {
	raylib.InitWindow(g.SCREEN_WIDTH, g.SCREEN_HEIGHT, "space invaders")

	ShipImg := raylib.LoadImage("./assets/Ship.png") // first we load the image
	if ShipImg == nil {
		return errImage
	}

	g.Ship = raylib.LoadTextureFromImage(ShipImg) // then we load its textures
	if g.Ship.ID == 0 {
		return errTexture
	}

	defer raylib.CloseWindow()

	raylib.SetTargetFPS(60)

	for !raylib.WindowShouldClose() { // as long as the window is open
		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.Black)

		raylib.DrawTexture(g.Ship, g.xCoords, g.yCoords, raylib.White) // drawing our ship

		// ship movement
		if raylib.IsKeyDown(raylib.KeyD) {
			if g.Ship.Width+g.xCoords+5 >= g.SCREEN_WIDTH {
			} else {
				g.xCoords += 5
			}

		}

		if raylib.IsKeyDown(raylib.KeyA) {
			if g.xCoords-1 <= 0 {
			} else {
				g.xCoords -= 5
			}
		}

		raylib.EndDrawing()
	}

	return nil
}
