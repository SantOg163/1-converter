package main

import "fmt"

const usdToEur = 0.85
const usdToRub = 76.5
const eurToRub = usdToRub / usdToEur

func main() {
	countFrom, from, to := input()
	countTo := convert(countFrom, from, to)
	fmt.Println("eurToRub = ", eurToRub)
	fmt.Println(countTo)
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

func convert(amount float64, fromCurrency string, toCurrency string) float64 {
	if fromCurrency == "rub" {
		if toCurrency == "eur" {
			return amount / eurToRub
		}
		if toCurrency == "usd" {
			return amount / usdToRub
		}
	}

	if fromCurrency == "usd" {
		if toCurrency == "eur" {
			return amount * usdToEur
		}
		if toCurrency == "rub" {
			return amount * usdToRub
		}
	}

	if fromCurrency == "eur" {
		if toCurrency == "usd" {
			return amount / usdToEur
		}
		if toCurrency == "rub" {
			return amount * eurToRub
		}
	}

	return amount
}
