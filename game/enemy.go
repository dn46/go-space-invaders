package game

import (
	"log"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

var (
	numberEnemies = 7
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
		Xpos:      10,
		Ypos:      100,
		imageDown: true,
		draw:      true,
		EnemyUp:   enemyUp,
		EnemyDown: enemyDown,
	}
}

func (g *Game) CreateEnemy() {
	for i := 0; i < numberEnemies; i++ {
		current_enemy := NewEnemy()
		current_enemy.Xpos += int32(i) * 75
		g.Enemies = append(g.Enemies, current_enemy)
	}
}

func (g *Game) DrawEnemy() {
	for _, current_enemy := range g.Enemies {
		if current_enemy.draw {
			if current_enemy.imageDown { // draw the down image
				raylib.DrawTexture(current_enemy.EnemyDown, current_enemy.Xpos, current_enemy.Ypos, raylib.White)
			} else { // draw the up image
				raylib.DrawTexture(current_enemy.EnemyUp, current_enemy.Xpos, current_enemy.Ypos, raylib.White)
			}
		}
	}
}
