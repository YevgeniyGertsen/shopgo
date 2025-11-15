package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/YevgeniyGertsen/shopgo/shop"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("1) Добавить товар в корзину")
	fmt.Println("2) расчет стоимости заказа")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	ch, _ := strconv.Atoi(input)

	if ch == 0 {
		shop.AddToCart(reader)
	} else {
		shop.OrderCost(reader)
	}
}
