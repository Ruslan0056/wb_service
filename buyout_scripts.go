package main

import (
	"fmt"
	"math"
)

func main() {
	a, b, c := 0.5238, 4.1905, 13.857
	amt_keywords, buyout_amt := 4, 100
	proportion := []float64{0.25, 0.25, 0.25, 0.25}
	Print_data(Pike_script(a, b, c, amt_keywords, buyout_amt, proportion))
	fmt.Println("________________")
	Print_data(Raise_script(4.2738, 0.1021, amt_keywords, buyout_amt, proportion))

	// Делаем проверку суммы выкупленных товаров
	//
	//	arr := Raise_script(4.2738, 0.1021, amt_keywords, buyout_amt, proportion)
	//	all := 0.0
	//	for i := 0; i < len(arr); i++ {
	//		sum := 0.0
	//		for j := 0; j < len(arr[i]); j++ {
	//			sum += arr[i][j]
	//		}
	//		fmt.Printf("%.2f\n", sum)
	//		all += sum
	//	}
	//	fmt.Println(all)
}

func Pike_script(a, b, c float64, amt_keywords, buyout_amt int, proportion []float64) [][]float64 {
	var amt_product_by_day [][]float64 //выделить память
	sum_result := 0.0

	for i := 1; i < 8; i++ {
		result := a*math.Pow(float64(i), 2) - b*float64(i) + c
		sum_result += result * float64(buyout_amt) / 100
		var amt_things_by_keyword []float64
		for j := 0; j < amt_keywords; j++ {
			amt_things_by_keyword = append(amt_things_by_keyword, result*float64(buyout_amt)/100*proportion[j])
		}
		amt_product_by_day = append(amt_product_by_day, amt_things_by_keyword)
	}
	return amt_product_by_day
}

func Raise_script(a, b float64, amt_keywords, buyout_amt int, proportion []float64) [][]float64 {
	var amt_product_by_day [][]float64 //выделить память
	sum_result := 0.0
	// a = 4.2738
	// b = 0.1021
	for i := 1; i < 8; i++ {
		result := a * math.Exp(b*float64(i))
		sum_result += result * float64(buyout_amt) / 100
		var amt_things_by_keyword []float64
		for j := 0; j < amt_keywords; j++ {
			amt_things_by_keyword = append(amt_things_by_keyword, result*float64(buyout_amt)/100*proportion[j])
		}
		amt_product_by_day = append(amt_product_by_day, amt_things_by_keyword)
	}
	return amt_product_by_day
}

func Print_data(amt_product_by_day [][]float64) {
	for i := 0; i < len(amt_product_by_day); i++ {
		fmt.Printf("%.2f\n", amt_product_by_day[i])
	}
}
