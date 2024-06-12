package game

import (
	"log"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

type Ship struct {
	Xpos      int32
	Ypos      int32
	Health    int
	Image     raylib.Texture2D
	Rectangle raylib.Rectangle
}

func NewShip() *Ship {

	// load the ship image
	ShipImg := raylib.LoadImage("./assets/Ship.png")
	if ShipImg.Width == 0 || ShipImg.Height == 0 {
		log.Fatal(errImage)
	}

	shipTexture := raylib.LoadTextureFromImage(ShipImg)
	if shipTexture.ID == 0 {
		log.Fatal(errTexture)
	}

	raylib.UnloadImage(ShipImg) // freeing the image after being loaded to a texture

	return &Ship{
		Xpos:      2,
		Ypos:      500,
		Image:     shipTexture,
		Rectangle: raylib.NewRectangle(2, 500, float32(shipTexture.Width), float32(shipTexture.Height)),
	}
}

func (s *Ship) moveShip(screenWidth int32) {

	if raylib.IsKeyDown(raylib.KeyD) {
		if s.Image.Width+s.Xpos+5 < screenWidth {
			s.Xpos += 5
		}
	}

	if raylib.IsKeyDown(raylib.KeyA) {
		if s.Xpos-5 > 0 {
			s.Xpos -= 5
		}
	}

	// update the ships rectangle
	s.Rectangle.X = float32(s.Xpos)
	s.Rectangle.Y = float32(s.Ypos)
}
