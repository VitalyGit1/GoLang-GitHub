// Дана строка. Выведите в консоль длину этой строки.

package main

import "fmt"

func main() {
	var text string
	fmt.Print("Enter the text without spaces: ")
	_, err := fmt.Scanln(&text)
	if err != nil {
		fmt.Println("You didn't enter with spaces")
		return
	}
	fmt.Println("count:", len([]rune(text)))
}
