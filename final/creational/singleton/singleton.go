package singleton

import "sync"

// Default difficulty level.
const normal = 2

// game is an interface that describes some video game entity.
type game interface {
	SetDifficulty(difficultyLevel int)
	GetDifficulty() int
	renderLevel() error
	startGame(difficultyLevel int) error
}

// gameSingleton is a Singleton object for game interface.
type gameSingleton struct {
	difficultyLevel int
	sync.Mutex
}

var (
	gameInstance *gameSingleton
	once         sync.Once
)

// GetNewGame start a new game for one time.
// In next calls of GetNewGame will return only game entity
// for future difficulty sets.
func GetNewGame() *gameSingleton {
	if gameInstance == nil {
		// once ensures that the following code is executed only one time.
		once.Do(func() {
			gameInstance = new(gameSingleton)
			gameInstance.startGame(normal)
		})
	}
	return gameInstance
}

// GetDifficulty returns current difficulty level.
func (g *gameSingleton) GetDifficulty() (difficultyLevel int) {
	return g.difficultyLevel
}

// SetDifficulty sets the difficulty level of game.
func (g *gameSingleton) SetDifficulty(difficultyLevel int) {
	g.Lock()
	defer g.Unlock()
	g.difficultyLevel = difficultyLevel
}

// renderLevel renders some game level.
func (g *gameSingleton) renderLevel() error {
	/*
		Massive code for rendering game level.
	*/

	return nil
}

// StartGame starts the new game.
func (g *gameSingleton) startGame(difficultyLevel int) error {
	g.SetDifficulty(difficultyLevel)
	if err := g.renderLevel(); err != nil {
		return err
	}

	/*
		Massive code for run game process.
	*/

	return nil
}
