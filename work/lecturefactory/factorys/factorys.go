package factorys

import (
	"fmt"
	"log"
)

// ENEMY
// ===================================================================
type IEnemy interface {
	Attack()
	GetName() string
}

type Enemy struct {
	ID     int
	Damage int
	HP     int
	Place  int
}

//=======================================================================

type Orc struct {
	Enemy
}

func newOrc() IEnemy {
	return &Orc{
		Enemy: Enemy{
			Damage: 2,
			HP:     3,
			Place:  50,
		},
	}
}

func (o Orc) Attack() {
	fmt.Printf("Orc beats you on %d damage!", o.Damage)
}

func (o Orc) GetName() string {
	return "Orc"
}

type Goblin struct {
	Enemy
}

func newGoblin() IEnemy {
	return &Orc{
		Enemy: Enemy{
			Damage: 1,
			HP:     1,
			Place:  20,
		},
	}
}

func (o Goblin) Attack() {
	fmt.Printf("Goblin beats you on %d damage!", o.Damage)
}

func (o Goblin) GetName() string {
	return "Goblin"
}

type Troll struct {
	Enemy
}

func newTroll() IEnemy {
	return &Orc{
		Enemy: Enemy{
			Damage: 10,
			HP:     30,
			Place:  4,
		},
	}
}

func (o Troll) Attack() {
	fmt.Printf("Troll beats you on %d damage!", o.Damage)
}

func (o Troll) GetName() string {
	return "Troll"
}

// ===================================================================

// GAME
// ===================================================================
type Game struct {
	Difficulty int
}

func NewGame(difficulty int) *Game {
	return &Game{Difficulty: difficulty}
}

func (g Game) StartGame() {
	g.renderLvl()

	enemy := CreateEnemy(g.Difficulty)
	g.SetupEnemies(enemy)
	g.start()
}

func (g Game) renderLvl() {
	fmt.Println("Rendering Nature")
}

func (g Game) SetupEnemies(es IEnemy) {
	fmt.Println("EnemiesSetup")
}

func (g Game) start() {
	fmt.Println("Starting game")
}

// ===================================================================

func CreateEnemy(difficulty int) IEnemy {
	var enemy IEnemy
	switch difficulty {
	case 1:
		enemy = newGoblin()
	case 2:
		enemy = newOrc()
	case 3:
		enemy = newTroll()
	default:
		log.Fatal("There is no such difficulty level %d", difficulty)
	}

	return enemy
}
