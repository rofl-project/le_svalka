package main

import "fmt"

const USD_TO_EUR = 0.9
const USD_TO_RUB = 30.0

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


func main()  {

	var dollars = 100.0
	var euro = 500.0
	var rub = 5000.0

	fmt.Println(dollars, "dolars to Rub = ", convertUsdToRub(dollars))
	fmt.Println(dollars, "dolars to Euro = ", convertUsdToEur(dollars))

	fmt.Println(euro, "euro to Rub = ", convertEurToRub(euro))
	fmt.Println(euro, "euro to Dollars = ", convertEurToUsd(euro))

	fmt.Println(rub, "rub to Dollars = ", convertRubToUsd(rub))
	fmt.Println(rub, "rub to Euro = ", convertRubToEur(rub))

}