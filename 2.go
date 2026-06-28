package main

import "fmt"

// Определяем структуру Person
type Person struct {
	Name string
	Age  int
}
n n n n
// Создаём слайс людей
var people []Person

// Функция для добавления человека в слайс
func addPerson(name string, age int) {
	person := Person{
		Name: name,
		Age:  age,
	}
	people = append(people, person)
}

// Функция для вывода списка людей
func printPeople() {
	for _, person := range people {
		fmt.Printf("Имя: %s, Возраст: %d\n", person.Name, person.Age)
	}
}

func main() {
	addPerson("Алексей", 30)
	addPerson("Мария", 25)
	addPerson("Дмитрий", 40)

	printPeople()
}
