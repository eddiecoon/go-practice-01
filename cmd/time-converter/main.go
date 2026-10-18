package main

import "fmt"

func main() {
	var min int
	fmt.Print("Введите количество минут: ")

	_, err := fmt.Scan(&min)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести целое число.")
		return
	}
	if min < 0 {
		fmt.Println("Ошибка: количество минут не должно быть отрицательным.")
		return
	}
	const minutesPerHour = 60
	hours := min / minutesPerHour
	minutes := min % minutesPerHour
	fmt.Printf("Полных часов: %d\n", hours)
	fmt.Printf("Оставшихся минут: %d\n", minutes)
}
