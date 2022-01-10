package main

import (
	"fmt"
	"sync"
	"time"
)

func test1(wg *sync.WaitGroup) {
	fmt.Println("Chạy hàm test 1")
	// Oánh dấu đã done
	wg.Done()
}

func test2(wg *sync.WaitGroup) {
	fmt.Println("Chạy hàm test 2")
	// Oánh dấu đã done
	wg.Done()
}

func main() {
	// Khai báo 1 wait group
	var wg sync.WaitGroup
	// Set giá trị counter = 2
	wg.Add(2)

	go test1(&wg)
	go test2(&wg)

	// Block main goroutines chạy tiếp, khi nào counter = 0 thì mới cho chạy q
	wg.Wait()

	fmt.Println("Good bye")
}