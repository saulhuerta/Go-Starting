package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 es par")
	} else {
		fmt.Println("7 es impart")
	}

	if 8%4 == 0 {
		fmt.Println("8 es dividible entre 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, " es negativo")
	} else if num < 10 {
		fmt.Println(num, " tiene 1 digito")
	} else {
		fmt.Println(num, " tiene múltiples dígitos")
	}
}
