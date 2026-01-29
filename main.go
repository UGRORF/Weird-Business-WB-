package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	errCommand            = errors.New("некорректное действие")             //— если в input прочитали некорректную строку с действием
	errReceiptCountParam  = errors.New("некорректное поступление")          // — если пришла строка receipt с некорректным количеством параметров
	errReceiptIDOrder     = errors.New("некорректный поступивший заказ")    //— если для receipt указан не числовой ID заказа
	errReceiptIDGood      = errors.New("некорректный поступивший товар")    //— если для receipt указан не числовой ID товара
	errReceiptPrice       = errors.New("некорректная цена поступления")     //— если для receipt указана не числовая цена
	errPurchaseCountParam = errors.New("некорректная покупка")              //— если пришла строка purchase с некорректным количеством параметров
	errPurchaseIDOrder    = errors.New("некорректный заказ покупателя")     //— если для purcahse указан не числовой ID заказа
	errOrderNotExist      = errors.New("несуществующий заказ покупателя")   //— если пытаются купить несуществующий товар на ПВЗ
	errIncorrectReturn    = errors.New("некорректный возврат")              //— если пришла строка return с некорректным количеством параметров
	errReturnIDOrder      = errors.New("некорректный возвращаемый заказ")   //— если для return указан не числовой ID заказа
	errReturnIDGood       = errors.New("некорректный возвращаемый товар")   //— если для return указан не числовой ID товара
	errNonExistentOrder   = errors.New("несуществующий возвращаемый заказ") //— если пытаются вернуть товар из заказа, который или не выкупался, или никогда не был в ПВЗ
	errNonExistentGood    = errors.New("несуществующий возвращаемый товар") //— если пытаются вернуть товар, который или не выкупался, или никогда не существовал в заказе
)

func main() {
	input := []string{
		"receipt 101 2001 1000",
		"receipt 101 2002 2000",
		"receipt 102 2003 4000",
		"purchase 102",
		"receipt 103 2004 8000",
		"purchase 101",
		"receipt 104 2005 16000",
		"return 101 2002",
		"return 103 2004",
	}

	handle(input)
}

func handle(input []string) {
	for _, request := range input {
		command := strings.Split(request, " ")

		switch command[0] {
		case "receipt":
			if len(command) != 4 {
				fmt.Printf("Ошибка: %v\n", errReceiptCountParam)
				break
			}
			idOrder, err := strconv.Atoi(command[1])
			if err != nil {
				fmt.Printf("Ошибка: %v\n", errReceiptIDOrder)
				break
			}

			idGood, err := strconv.Atoi(command[2])
			if err != nil {
				fmt.Printf("Ошибка: %v\n", errReceiptIDGood)
				break
			}

			price, err := strconv.ParseFloat(command[3], 64)
			if err != nil {
				fmt.Printf("Ошибка: %v\n", errReceiptPrice)
				break
			}

			addReceipt(idOrder, idGood, price)
		case "purchase":
			fmt.Println("purchase")
		case "return":
			fmt.Println("return")
		default:
			fmt.Printf("Ошибка: %v\n", errCommand)
		}
	}
}
