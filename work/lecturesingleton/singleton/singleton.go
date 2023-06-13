package singleton

import (
	"strconv"
	"sync"
)

type memoryStorage struct {
	Cache map[string]string
	sync.Mutex
}

var ms *memoryStorage
var once sync.Once

func GetConnecton() *memoryStorage {
	once.Do(func() {
		ms = &memoryStorage{}
		ms.Cache = make(map[string]string)
	})

	return ms
}

func (ms *memoryStorage) Write(i int) {
	ms.Lock()
	defer ms.Unlock()
	ms.Cache["test"] = strconv.Itoa(i)
}
