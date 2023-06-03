// Калькулятор, в который вносишь непобходимую выручку в неделю для появления в топе товаров. Калькулятор рассчитывает какое количество товара надо покупать в день. Возможна реализация различных сценариев
// Логика калькулятора
// Входные данные
// Целевая выручка в месяц
// Стоимость 1шт. собственного товара без СПП
// По каким запросам делать выкупы. Выбор количества ключевых запросов
// Размерная сетка
// Органические выкупы
// Выбор сценария выкупа: нарастающая, убывающая, лесенка, внешняя реклама (посмотреть еще). Предлагать ориентироваться на тренд
// Вывод
// Юнит экономика?
// Количество выкупаемых товаров по дням недели

// Сервер. Клиент-серверное взаимодействие
// БД
// Визуал

package main

import (
	"fmt"
	"math"
)

// type wb_service struct {
// 	product_price int
// 	target_revenue int
// 	amt_keywords int
// 	keywords []string
// 	proportion []float64
// }

func main() {

	var product_price, target_revenue, amt_keywords int
	var keywords []string
	var proportion []float64
	fmt.Print("Введите цену одной единицы товара: ")
	fmt.Scanln(&product_price)
	fmt.Print("Введите целевую выручку: ")
	fmt.Scanln(&target_revenue)
	fmt.Print("Количество ключевых слов: ")
	fmt.Scanln(&amt_keywords)

	//Ввод исх данных по ключевым словам в запросе
	for i := 0; i < amt_keywords; i++ {
		var a string
		var b float64
		fmt.Println("Введите ключевое слово ", i+1)
		fmt.Scanln(&a)
		fmt.Println("Введите долю выкупа по ключевому слову ", i+1)
		fmt.Scanln(&b)
		keywords = append(keywords, a)
		proportion = append(proportion, b)
	}
	// Вывод исходных данных
	//fmt.Println(product_price, target_revenue, amt_keywords, keywords)

	// Расчет количетва выкупов
	weekly_buyout := target_revenue / 4 * 100
	buyout_amt := weekly_buyout / product_price //количество выкупов на неделю
	var amt_product_by_day [][]int              //выделить память
	// option := 0
	// fmt.Println("Выберите сценарий:")
	// fmt.Println("1 - Низходящая регрессия;")
	// fmt.Println("2 - Возрастающая регрессия;")
	// fmt.Println("3 - Волнобразный сценарий;")
	// fmt.Println("4 - Сценарий с пиками по краям.")
	// fmt.Scanln(&option)

	// Вычиления выкупов по дням недели
	sum_result := 0.0

	for i := 1; i < 8; i++ {
		result := 0.5238*math.Pow(float64(i), 2) - 4.1905*float64(i) + 13.857
		sum_result += result * float64(buyout_amt) / 100
		var amt_things_by_keyword []int
		for j := 0; j < amt_keywords; j++ {
			amt_things_by_keyword = append(amt_things_by_keyword, int(result*float64(buyout_amt)/100*proportion[j]))
		}
		amt_product_by_day = append(amt_product_by_day, amt_things_by_keyword)
	}
	// switch option {

	// case 1:

	// 	for i := 1; i < 8; i++ {
	// 		result := 0.2208 * (math.Pow(float64(i), -0.384))
	// 		amt_product_by_day = append(amt_product_by_day, result*float64(buyout_amt)/100)
	// 		sum_result += result * float64(buyout_amt) / 100

	// 	}

	// case 2:
	// 	for i := 1; i < 8; i++ {
	// 		result := 4.2738 * math.Exp(0.1021*float64(i))
	// 		amt_product_by_day = append(amt_product_by_day, result*float64(buyout_amt)/100)
	// 		sum_result += result * float64(buyout_amt) / 100
	// 	}

	// case 3:
	// 	for i := 1; i < 8; i++ {
	// 		result := -0.1111*math.Pow(float64(i), 3) + 1.3333*math.Pow(float64(i), 2) - 4.6984*float64(i) + 8.5714 // низходящий
	// 		//result2 := 0.0333 * math.Pow(i,5) - 0.6667 * math.Pow(i,4) + 4.8333 * math.Pow(i,3) - 15.333 * math.Pow(i,2) - 20.133 * i + 4
	// 		amt_product_by_day = append(amt_product_by_day, result*float64(buyout_amt)/100)
	// 		sum_result += result * float64(buyout_amt) / 100
	// 	}
	// case 4:
	// }

	fmt.Println(amt_product_by_day)

	fmt.Println(int(sum_result))

	// Вычисления выкупов по размерам

	//размерная сетка
	// size_chart := make(map[string]float64)
	// amt_size := 5

	// for i := 0; i < amt_size; i++ {
	// 	fmt.Println("Введите размер")
	// 	size := ""
	// 	fmt.Scanln(&size)
	// 	fmt.Println("Ввведите долю выкупа")
	// 	proportion_size := 0.0
	// 	fmt.Scanln(&proportion_size)
	// 	size_chart[size] = proportion_size
	// }

	// fmt.Println(size_chart)

}
