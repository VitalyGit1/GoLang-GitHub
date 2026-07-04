package main

import "fmt"

func main() {
	var number int
	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Ошибка ввода.")
		return
	}

	if number < 0 {
		fmt.Println("This number is less than zero")
	}

	if number > 0 {
		fmt.Println("This number is greater than zero")
	}

}
