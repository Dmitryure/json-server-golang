package main

// import "fmt"

// var test = []string{"1", "2", "3", "4", "as", "6", "7"}

// func main() {
// 	for val := range generateStr(9) {
// 		fmt.Println(val)
// 	}
// 	fmt.Println(test)
// }

// func generateStr(count int) chan string {
// 	ch := make(chan string)

// 	go func(ch chan string) {
// 		for i := 0; i < count; i++ {
// 			if len(test) == 0 {
// 				ch <- ""
// 			} else {
// 				ch <- string(test[0])
// 				test = append(test[:0], test[1:]...)
// 			}
// 		}
// 		close(ch)
// 	}(ch)
// 	return ch
// }

// func count(start int, end int) chan int {
// 	ch := make(chan int)

// 	go func(ch chan int) {
// 		for i := start; i <= end; i++ {
// 			// Блокирует операцию
// 			ch <- i
// 		}

// 		close(ch)
// 	}(ch)

// 	return ch
// }
