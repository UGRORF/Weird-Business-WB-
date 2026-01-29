package main

import (
	"fmt"
	"log"
)

var (
	ordersOnPVZ  = make(map[int]map[int]float64)
	issuedOrders = make(map[int]map[int]float64)
)

func addReceipt(idOrder, idGood int, price float64) {
	if ordersOnPVZ[idOrder] == nil {
		ordersOnPVZ[idOrder] = make(map[int]float64)
	}
	ordersOnPVZ[idOrder][idGood] = price

	log.Printf("На склад поступил товар %d из заказа %d, цена %.f руб.", idGood, idOrder, price)
}

func purchaseOrder(idOrder int) {
	sum := 0.0

	if ordersOnPVZ[idOrder] == nil {
		fmt.Printf("Ошибка: %v\n", errOrderNotExist)
		return
	}

	if issuedOrders[idOrder] != nil {
		fmt.Printf("Ошибка: %v\n", errOrderAlreadyIssued)
		return
	}

	for good, price := range ordersOnPVZ[idOrder] {
		if issuedOrders[idOrder] == nil {
			issuedOrders[idOrder] = make(map[int]float64)
		}
		issuedOrders[idOrder][good] = price

		sum += price
	}

	delete(ordersOnPVZ, idOrder)
	log.Printf("Покупатель забрал заказ %d на сумму %.f руб.", idOrder, sum)
}

func returnGood(idOrder, idGood int) {
	if issuedOrders[idOrder] == nil {
		fmt.Printf("Ошибка: %v\n", errNonExistentOrder)
		return
	}

	price, exist := issuedOrders[idOrder][idGood]
	if !exist {
		fmt.Printf("Ошибка: %v\n", errNonExistentGood)
		return
	}

	if ordersOnPVZ[idOrder] == nil {
		ordersOnPVZ[idOrder] = make(map[int]float64)
	}
	ordersOnPVZ[idOrder][idGood] = price
	delete(issuedOrders[idOrder], idGood)

	if len(issuedOrders[idOrder]) == 0 {
		delete(issuedOrders, idOrder)
	}

	log.Printf("Покупатель вернул товар %d из заказа %d, сумма возврата %.f руб.", idGood, idOrder, price)
}

func closingTheShift() {
	sum := 0.0
	countGoods := 0
	for _, orders := range ordersOnPVZ {
		for _, price := range orders {
			sum += price
			countGoods++
		}
	}
	log.Printf("На складе осталось %d товаров из %d заказов на сумму %.f руб.", countGoods, len(ordersOnPVZ), sum)
}
