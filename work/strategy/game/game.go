package game

// Chatacter предоставляет контекст выполнения той или иной стратегии.
type Character struct {
	WeaponBehaviorStrategy
}

func (c *Character) SetBehavior(b WeaponBehaviorStrategy) {
	c.WeaponBehaviorStrategy = b
}

func (c *Character) Fight() {
	c.UseWeapon()
}
