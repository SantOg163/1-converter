package main

import "fmt"

func main() {
	const usdToEuro = 0.85
	const usdToRub = 76.5
	eurToRub := usdToRub / usdToEuro
	fmt.Println("eurToRub = ", eurToRub)
}
