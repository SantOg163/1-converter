package main

import "fmt"

var usdTo = map[string]float64{
	"rub":  76.5,
	"euro": 0.85,
	"usd":  1.0,
}

const usdToEur = 0.85
const usdToRub = 76.5
const eurToRub = usdToRub / usdToEur

func main() {
	from := inputCurrency("Введите исходную валюту:")
	countFrom := inputCount()
	to := inputCurrency("Введите целевую валюту:")

	countTo := convert(countFrom, from, to)

	fmt.Println("eurToRub = ", eurToRub)
	fmt.Println(countTo)
}
func inputCount() float64 {
	count := 0.0
	for {
		fmt.Println("Введите число:")
		_, err := fmt.Scan(&count)
		if err == nil {
			break
		} else {
			fmt.Println("Введите корректные данные!")
		}
	}
	return count
}
func inputCurrency(text string) string {
	currency := ""
	for {
		fmt.Println(text)
		_, err := fmt.Scan(&currency)
		if err == nil && (currency == "rub" || currency == "usd" || currency == "eur") {
			break
		} else {
			fmt.Println("Введите корректные данные!")
		}
	}
	return currency
}

func convert(amount float64, fromCurrency string, toCurrency string) float64 {

	if fromCurrency == toCurrency {
		return amount
	}

	countInUsd := amount / usdTo[fromCurrency]

	// switch fromCurrency {
	// case "rub":
	// 	countInUsd = amount / usdToRub
	// case "eur":
	// 	countInUsd = amount / usdToEur
	// default:
	// 	countInUsd = amount
	// }

	// switch toCurrency {
	// case "rub":
	// 	return countInUsd * usdToRub
	// case "eur":
	// 	return countInUsd * usdToEur
	// default:
	// 	return countInUsd
	// }
	return  countInUsd * usdTo[toCurrency]
}
