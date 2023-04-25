package main

import (
	"fmt"
)

func main() {
	fmt.Println("Три числа.")

	//переменная создана для удобства смены условий программы.
	numRequired := 5
	fmt.Println("Эта программа определит, есть ли среди ваших трех чисел, число больше", numRequired)

	fmt.Println("Введите первое число:")
	var numOne int
	fmt.Scan(&numOne)

	fmt.Println("Введите второе число:")
	var numTwo int
	fmt.Scan(&numTwo)

	fmt.Println("Введите третье число:")
	var numThree int
	fmt.Scan(&numThree)

	if numOne > numRequired {
		fmt.Println("Среди введённых чисел есть число больше", numRequired)
		fmt.Println("Это число", numOne)
	} else if numTwo > numRequired {
		fmt.Println("Среди введённых чисел есть число больше", numRequired)
		fmt.Println("Это число", numTwo)
	} else if numThree > numRequired {
		fmt.Println("Среди введённых чисел есть число больше", numRequired)
		fmt.Println("Это число", numThree)
	} else {
		fmt.Println("Среди введённых чисел нет числа больше", numRequired)
	}

}
