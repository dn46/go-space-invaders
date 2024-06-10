package game

import (
	"log"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

type Ship struct {
	Xpos  int32
	Ypos  int32
	Image raylib.Texture2D
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
		Xpos:  2,
		Ypos:  500,
		Image: shipTexture,
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
}
