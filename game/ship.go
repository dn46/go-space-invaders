package game

import (
	"log"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

type Ship struct {
	Xpos           int32
	Ypos           int32
	Health         int
	Image          raylib.Texture2D
	Rectangle      raylib.Rectangle
	IsInvincible   bool
	InvincibleTime float32
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
		Xpos:           2,
		Ypos:           500,
		Image:          shipTexture,
		Health:         3,
		Rectangle:      raylib.NewRectangle(2, 500, float32(shipTexture.Width), float32(shipTexture.Height)),
		IsInvincible:   false,
		InvincibleTime: 0,
	}
}

func (s *Ship) Hit() {
	if !s.IsInvincible {
		s.Health--
		s.IsInvincible = true
		s.InvincibleTime = 2
	}
}

func (s *Ship) UpdateShip(dt float32) {
	if s.IsInvincible {
		s.InvincibleTime -= dt
		if s.InvincibleTime <= 0 {
			s.IsInvincible = false
		}
	}
}

func (s *Ship) IsDestroyed() bool {
	return s.Health <= 0
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
