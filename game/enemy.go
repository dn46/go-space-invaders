package game

import (
	"log"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

type Enemy struct {
	Xpos      int32
	Ypos      int32
	imageDown bool // this is to change between images
	draw      bool // dead enemies should not be drawn
	EnemyUp   raylib.Texture2D
	EnemyDown raylib.Texture2D
}

func NewEnemy() *Enemy {

	EnemyImg1 := raylib.LoadImage("./assets/enemy1.png")
	if EnemyImg1 == nil {
		log.Fatal(errImage)
	}

	EnemyImg2 := raylib.LoadImage("./assets/enemy2.png")
	if EnemyImg2 == nil {
		log.Fatal(errImage)
	}

	enemyUp := raylib.LoadTextureFromImage(EnemyImg1)
	if enemyUp.ID == 0 {
		log.Fatal(errTexture)
	}

	enemyDown := raylib.LoadTextureFromImage(EnemyImg2)
	if enemyDown.ID == 0 {
		log.Fatal(errTexture)
	}

	raylib.UnloadImage(EnemyImg1)
	raylib.UnloadImage(EnemyImg2)

	return &Enemy{
		Xpos:      2,
		Ypos:      100,
		imageDown: false,
		draw:      false,
		EnemyUp:   enemyUp,
		EnemyDown: enemyDown,
	}
}
