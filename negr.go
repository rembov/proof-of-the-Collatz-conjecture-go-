package main

import "fmt"

func negr(goida int) (int, int) {

	k := 0
	for goida != 1 {
		if goida%2 == 0 {
			goida /= 2
			k += 1
			fmt.Println(goida, k)
		} else {
			goida = 3*goida + 1
			k += 1
			fmt.Println(goida, k)
		}
	}
	return goida, k
}

func main() {
	n := 0
	for n < 1 {
		print("Введите натуральное число: ")
		fmt.Scan(&n)
	}
	res, steps := negr(n)
	fmt.Printf("Результат: %d\nКоличество шагов: %d", res, steps)
}
