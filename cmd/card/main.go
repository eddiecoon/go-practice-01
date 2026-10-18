package main

import "fmt"

func main() {
	participant := "Иванов Иван Иванович"
	var group string = "ТЭП-03-26"
	const course = "Бэкенд-разработка на языке Go"
	completed := 2
	interest := true

	fmt.Println("Участник", participant)
	fmt.Printf("Группа: %s\n", group)
	fmt.Printf("Курс: %s\n", course)
	fmt.Printf("Завершено работ: %d\n", completed)
	fmt.Printf("Интерес к бэкенду: %.t\n", interest)
}
