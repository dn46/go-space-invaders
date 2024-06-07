package main

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

var (
	SCREEN_WIDTH  int32 = 800
	SCREEN_HEIGHT int32 = 600
)

func main() {
	raylib.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "raylib [core] example - basic window")
	defer raylib.CloseWindow()

	raylib.SetTargetFPS(60)

	for !raylib.WindowShouldClose() { // as long as the window is open
		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.RayWhite)
		raylib.DrawText("Congrats! You created your first window", 190, 200, 20, raylib.LightGray)

		raylib.EndDrawing()
	}
}
