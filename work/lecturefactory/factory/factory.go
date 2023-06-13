package factory

import "fmt"

type IEnemy interface {
	Attack()
}

type Enemy struct {
	power int
	name  string
}

/*
В Go невозможно реализовать классический вариант паттерна Фабричный метод,
поскольу в языке отсутствуют возможности ООП, в том числе классы и наследственность.
Несмотря на это, мы все же можем реализовать базовую версию этого паттерна — Простая фабрика.
*/

func NewEnemy(enemyType string) (IEnemy, error) {
	switch enemyType {
	case "orc":
		return newOrc(), nil
	case "goblin":
		return newGoblin(), nil
	default:
		return nil, fmt.Errorf("there is no such enemy type: %s", enemyType)
	}
}

type orc struct {
	Enemy
}

func newOrc() IEnemy {
	return &orc{
		Enemy: Enemy{
			power: 2,
			name:  "orc",
		},
	}
}

func (o orc) Attack() {
	fmt.Println("Orc attacts!")
}

type goblin struct {
	Enemy
	arrows int
}

func newGoblin() IEnemy {
	return &goblin{
		Enemy: Enemy{
			power: 1,
			name:  "goblin",
		},
		arrows: 20,
	}
}

func (g goblin) Attack() {
	fmt.Println("Goblin attacs!")
}

func (g goblin) Shoot() {
	fmt.Println("Goblin shoots")
}
