package main

import "fmt"

func main() {
	const usdToEuro = 0.85
	const usdToRub = 76.5
	var rub = 115.0
	result := (rub / usdToRub) * usdToEuro
	fmt.Println("result = ", result)
}
