package singleton

import "sync"

type singleton struct{}

var (
	instance *singleton
	once     sync.Once
)

// GetInstanse создаёт (единожды) и возвращает instanse.
func GetInstanse() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}
