package shop

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func OrderCost(reader *bufio.Reader) {
	fmt.Printf("\n--- Расчет стоимости заказа ---")
	fmt.Println("Введите стоимость товара: ")

	pStr, _ := reader.ReadString('\n')
	pStr = strings.TrimSpace(pStr)
	price, _ := strconv.Atoi(pStr)

	fmt.Println("Введите количество ")
	qStr, _ := reader.ReadString('\n')
	qStr = strings.TrimSpace(qStr)
	qty, _ := strconv.Atoi(qStr)

	total := price * qty

	fmt.Printf("Стоимость заказа: %d тенге\n", total)
}
