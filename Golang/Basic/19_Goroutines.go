// package main

// import (
//     "fmt"
//     "time"
// )

// func f(from string) {
//     for i := 0; i < 3; i++ {
//         fmt.Println(from, ":", i)
//     }
// }

// func main() {

//     f("direct")

//     go f("goroutine")

//     go func(msg string) {
//         fmt.Println(msg)
//     }("going")

//     time.Sleep(time.Second)
//     fmt.Println("done")
// }

package main

import (
	"fmt"
	"time"
)

func test1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Trứng rán cần mỡ")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Yêu không cần cớ")
}

func test2() {
	time.Sleep(150 * time.Millisecond)
	fmt.Println("Bắp cần bơ")
	time.Sleep(150 * time.Millisecond)
	fmt.Println("Cần cậu cơ")
}

func main() {
	go test1()
	go test2()
	time.Sleep(2 * time.Second)
	fmt.Println("Good bye")
}

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func hienCoXinhGaiKhong() {
// 	fmt.Println("Hien rat xinh gai")
// }
// func main() {
// 	go hienCoXinhGaiKhong()
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Good bye")
// }
