package main

import "fmt"

func main() {
	const usdToEuro = 0.85
	const usdToRub = 76.5
	eurToRub := usdToRub / usdToEuro
	input()
	fmt.Println("eurToRub = ", eurToRub)
}
func input() (float64, string, string) {
	count, from, to := 0.0, "usd", "eur"
	fmt.Println("Введите число:")
	fmt.Scan(&count)
	fmt.Println("Введите исходную валюту:")
	fmt.Scan(&from)
	fmt.Println("Введите целеву/ валюту:")
	fmt.Scan(&to)
	return count, from, to
}
