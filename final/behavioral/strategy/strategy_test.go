package strategy

import (
	"testing"
)

// Character предоставляет контекст выполнения той или иной стратегии
// испоьзования оружия (WeaponBehaviorStrategy).
type Character struct {
	WeaponBehaviorStrategy
}

func (c *Character) SetBehavior(b WeaponBehaviorStrategy) {
	c.WeaponBehaviorStrategy = b
}

func (c *Character) Fight() {
	c.UseWeapon()
}

func TestStrategy(t *testing.T) {
	queen := Character{}
	queen.SetBehavior(BowAndArrowBehavior{})
	queen.Fight()

	goblin := Character{}
	goblin.SetBehavior(KnifeBehavior{})
	goblin.Fight()
}
