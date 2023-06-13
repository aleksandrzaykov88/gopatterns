package main

import (
	"fmt"
	"lecturesingleton/singleton"
	"sync"
)

func main() {
	ms1 := singleton.GetConnecton()
	ms2 := singleton.GetConnecton()

	var wg sync.WaitGroup

	for i := 0; i <= 3000; i++ {
		wg.Add(2)
		go func() {
			ms1.Write(i)
			wg.Done()
		}()

		go func() {
			ms2.Write(i)
			wg.Done()
		}()

	}
	wg.Wait()

	fmt.Println("ms1:", &ms1)
	fmt.Println("ms1:", &ms2)

}
