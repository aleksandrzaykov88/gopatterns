package strategy

import "fmt"

// WeaponBehaviorStrategy предоставляет интерфейс использования различных видов оружия (конекретных стратегий).
type WeaponBehaviorStrategy interface {
	UseWeapon()
}

type AxeBehavior struct{}

func (a AxeBehavior) UseWeapon() {
	fmt.Println("use axe")
}

type SwordBehavior struct{}

func (s SwordBehavior) UseWeapon() {
	fmt.Println("use sword")
}

type BowAndArrowBehavior struct{}

func (ba BowAndArrowBehavior) UseWeapon() {
	fmt.Println("use bow and arrow")
}

type KnifeBehavior struct{}

func (k KnifeBehavior) UseWeapon() {
	fmt.Println("use knife")
}
