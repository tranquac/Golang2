package main

import (
	"fmt"
	"sync"
)

var x int64 = 0
// Khai báo mutex
var mutex = &sync.Mutex{}

func addOne(wg *sync.WaitGroup) {
	// Lock lại
	//mutex.Lock()
	x = x + 1
	// Unlock
	//mutex.Unlock()

	wg.Done()
}
func main() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go addOne(&w)
	}
	w.Wait()
	fmt.Println("Giá trị của x là: ", x)
}