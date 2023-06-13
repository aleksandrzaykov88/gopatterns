package factorym

import "fmt"

type IEnemy interface {
	Attack()
}

type Orc struct{}

func (o Orc) Attack() {
	fmt.Println("Orc attacts!")
}

type Goblin struct{}

func (g Goblin) Attack() {
	fmt.Println("Goblin attacs!")
}

type EnemyFactory interface {
	CreateEnemy() IEnemy
}

type GoblinFactory struct{}

func (gc GoblinFactory) CreateEnemy() IEnemy {
	return Goblin{}
}

type OrcFactory struct{}

func (oc OrcFactory) CreateEnemy() IEnemy {
	return Orc{}
}
