package window

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

var (
	SCREEN_WIDTH  int32 = 800
	SCREEN_HEIGHT int32 = 600

	xCoords int32 = 2
	yCoords int32 = 500
)

func StartWindow() {
	raylib.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "raylib [core] example - basic window")

	ShipImg := raylib.LoadImage("./assets/Ship.png") // first we load the image
	Ship := raylib.LoadTextureFromImage(ShipImg)     // then we load its textures

	defer raylib.CloseWindow()

	raylib.SetTargetFPS(60)

	for !raylib.WindowShouldClose() { // as long as the window is open
		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.Black)

		raylib.DrawTexture(Ship, xCoords, yCoords, raylib.White) // drawing our ship

		// ship movement
		if raylib.IsKeyDown(raylib.KeyD) {
			xCoords += 5
		}

		if raylib.IsKeyDown(raylib.KeyA) {
			xCoords -= 5
		}

		// raylib.DrawText("Congrats! You created your first window", 190, 200, 20, raylib.LightGray)

		raylib.EndDrawing()
	}
}
