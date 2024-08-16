package main

import (
	"bufio"
	"fmt"
	"go/scanner"
	"os"
	"strings"
)

// // func main(){
// //     readerWords("n")
// // }

// // func readerWords(n string)(x int){
// //     scannerIN:= bufio.NewScanner(os.Stdin)
// //     scannerIN.Scan()
// //     input := scannerIN.Text()
// //     scanner := bufio.NewScanner(strings.NewReader(input))
// //     // scanner.Split(bufio.ScanWords)
// //     count := 0
// // 	for scanner.Scan() {
// // 		count++
// // 	}
// // 	if err := scanner.Err(); err != nil {
// // 		fmt.Fprintln(os.Stderr, "reading input:", err)
// // 	}
// // 	fmt.Printf("%d\n", count)
// //     return x
// // }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main(){
// 	fmt.Println(readerWords("n") + readerWordsn("v"))

// }

// func readerWords(n string)(x int){
//     scannerIN:= bufio.NewScanner(os.Stdin)
//     scannerIN.Scan()
//     input := scannerIN.Text()
//     scanner := bufio.NewScanner(strings.NewReader(input))
//     scanner.Split(bufio.ScanWords)
//     var  count = 0
// 	for scanner.Scan() {
// 		count++

// 	}
// 	fmt.Printf("%d\n", count )
//     return x
// }
// func readerWordsn(n string)(x int){
//     scannerIN:= bufio.NewScanner(os.Stdin)
//     scannerIN.Scan()
//     input := scannerIN.Text()
//     scanner := bufio.NewScanner(strings.NewReader(input))
//     var  count = 0
// 	for scanner.Scan() {
// 		count++

// 	}
// 	fmt.Printf("%d\n", count )
//     return x
// }

func main(){
	
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	size := scanner
	elements := make([]int, size)
	println(elements)
}

