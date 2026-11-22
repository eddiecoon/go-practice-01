package main

import "fmt"

func main() {
	var a, b, c = 20, 25, 35
	fmt.Println("Сумма чисел: ", a+b+c)
	av := (float64(a) + float64(b) + float64(c)) / 3
	fmt.Printf("Среднее арифметическое: %.2f\n", av)

}
