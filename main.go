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

type wb_service struct {
}

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

	//fmt.Println(product_price, target_revenue, amt_keywords, keywords)

	weekly_buyout := target_revenue / 4 * 100
	buyout_amt := weekly_buyout / product_price //количество выкупов на неделю
	var amt_product_by_day []float64
	sum_result := 0.0
	for i := 1; i < 8; i++ {
		result := 0.2208 * (math.Pow(float64(i), -0.384))
		amt_product_by_day = append(amt_product_by_day, result*float64(buyout_amt)/100)
		sum_result += result * float64(buyout_amt) / 100

	}

	fmt.Println(amt_product_by_day)
	fmt.Println(int(sum_result))
}
