package greetings

import "fmt"

func Helo(name string) string {

	fmt.Sprint("Привет, %v. Добро пожаловать!", name)
}