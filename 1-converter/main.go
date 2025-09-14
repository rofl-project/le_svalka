package main

import "fmt"

const USD_TO_EUR = 0.9
const USD_TO_RUB = 30.0

const RUB = "rub"
const EUR = "eur"
const USD = "usd"

// Сделать считывание ввода пользователя в виде отдельной функции
// Сделать пустую функцию расчета, которая будет принимать 1 число и 2 валюты - исходная и целевая

func convertUsdToRub(amount float64) float64 {
	return amount * USD_TO_RUB
}

func convertRubToUsd(amount float64) float64 {
	return amount / USD_TO_RUB
}

func convertUsdToEur(amount float64) float64 {
	return amount * USD_TO_EUR
}

func convertEurToUsd(amount float64) float64 {
	return amount / USD_TO_EUR
}

func convertRubToEur(amount float64) float64 {
	return convertUsdToEur(convertRubToUsd(amount))
}

func convertEurToRub(amount float64) float64 {
	return convertUsdToRub(convertEurToUsd(amount))
}

func getCurrency(param string) string {
	var currency string = ""
	fmt.Printf("Enter %s currency: usd / rub / eur \n", param)

	for {
		fmt.Scan(&currency)

		if currency == RUB || currency == USD || currency == EUR {
			break
		}

		fmt.Printf("Enter CorrectValue!!! \n%s currency: usd / rub / eur\n", param)
	}

	return currency
}

func getInput() (string, string, float64) {
	var amount float64 = 0
	var fromCurrency string = getCurrency("FROM")
	var toCurrency string = getCurrency("TO")

	fmt.Println("Enter amount: ")
	fmt.Scan(&amount)

	return fromCurrency, toCurrency, amount
}

func uniConvert(fromCurrency string, toCurrency string, amount float64) float64 {

	// var amount float64 = 0
	// var fromCurrency string = ""
	// var toCurrency string = ""
	var convertResult float64 = 0

	switch fromCurrency {
	case USD:
		switch toCurrency {
		case USD:
			convertResult = amount
		case RUB:
			convertResult = convertUsdToRub(amount)
		case EUR:
			convertResult = convertUsdToEur(amount)
		}
	case EUR:
		switch toCurrency {
		case USD:
			convertResult = convertEurToUsd(amount)
		case RUB:
			convertResult = convertEurToRub(amount)
		case EUR:
			convertResult = amount
		}
	case RUB:
		switch toCurrency {
		case USD:
			convertResult = convertRubToUsd(amount)
		case RUB:
			convertResult = amount
		case EUR:
			convertResult = convertRubToEur(amount)
		}
	}

	// fmt.Printf("Cobvert from %s to %s is equal: %.2f \n", fromCurrency, toCurrency, convertResult)

	return convertResult
}

func main() {

	// var amount float64 = 0
	// var fromCurrency string = ""
	// var toCurrency string = ""
	// var convertResult float64 = 0

	fromCurrency, toCurrency, amount := getInput()
	convertResult := uniConvert(fromCurrency, toCurrency, amount)

	fmt.Printf("Convert from %s to %s is equal: %.2f \n", fromCurrency, toCurrency, convertResult)

	// uniConvert

	// fmt.Println("Enter FROM currency: usd / rub / eur")
	// fmt.Scan(&fromCurrency)

	// fmt.Println("Enter TO currency: usd / rub / eur")
	// fmt.Scan(&toCurrency)

	// fmt.Println("Enter AMOUNT of currency: ")
	// fmt.Scan(&amount)

	// var dollars = 100.0
	// var euro = 500.0
	// var rub = 5000.0

	// fmt.Println(dollars, "dolars to Rub = ", convertUsdToRub(dollars))
	// fmt.Println(dollars, "dolars to Euro = ", convertUsdToEur(dollars))

	// fmt.Println(euro, "euro to Rub = ", convertEurToRub(euro))
	// fmt.Println(euro, "euro to Dollars = ", convertEurToUsd(euro))

	// fmt.Println(rub, "rub to Dollars = ", convertRubToUsd(rub))
	// fmt.Println(rub, "rub to Euro = ", convertRubToEur(rub))

}
