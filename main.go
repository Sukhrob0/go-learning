// // package main

// // import (
// // 	"container/list"
// // 	"fmt"
// // )

// // // func main() {
// // // 	cycle()
// // // }

// // // func sum(n int) int {
// // //  sum := 0
// // //  for n > 0 {
// // //   sum += n % 10
// // //   n /= 10
// // //  }
// // //  return sum
// // // }

// // // func cycle() {
// // //  for i := 0; i <= 1000; i++ {
// // //   if (i%3 == 0) && (i%5 != 0) && (sum(i) < 10) {
// // //    fmt.Println(i)
// // //   }
// // //  }
// // // }

// // func ArraySearch(list [] int, item int ){

// // 	start := 0
// // 	end := list[len(list)-1]
// // 	mid := start + end
// // 	guess := list[mid]

// // if guess < item{
	
// // 	// start := mid + 1
// // }


// // }



// // func binarySearch(list []int, item int) int {
// //  low := 0
// //  high := len(list) - 1

// //  for low <= high {
// //   mid := (low + high) / 2
// //   guess := list[mid]

// //   if guess == item {
// //    return mid
// //   } else if guess > item {
// //    high = mid - 1
// //   } else {
// //    low = mid + 1
// //   }
// //  }

// //  return -1 // Целевое значение не найдено
// // }

// // func main() {
// //  MyList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 17, 100}
// //  index := binarySearch(MyList, 17)
// //  fmt.Println(index) // Вывод: 16 (первый индекс, где находится 17)
// // }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	// An artificial input source.
//     var input string
//     fmt.Scan(&input)
// 	scanner := bufio.NewScanner(strings.NewReader(input))
// 	// Set the split function for the scanning operation.
// 	scanner.Split(bufio.ScanWords)
// 	// Count the words.
// 	count := 0
// 	for scanner.Scan() {
// 		count++
// 	}
// 	if err := scanner.Err(); err != nil {
// 		fmt.Fprintln(os.Stderr, "reading input:", err)
// 	}
// 	fmt.Printf("%d\n", count)
// }

