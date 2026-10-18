package main

import "fmt"

func main() {
	participant := "Джессе"
	var group string = "ТЭП-03-26"
	const course = "Бэкенд-разработка на языке Go"
	completed := 1
	interest := true

	fmt.Println("Участник", participant)
	fmt.Printf("Группа: %s\n", group)
	fmt.Printf("Курс: %s\n", course)
	fmt.Printf("Завершено работ: %d\n", completed)
	fmt.Printf("Интерес к бэкенду: %.t\n", interest)

	completed = 2
	fmt.Printf("После изменения: %d\n", completed)
}
