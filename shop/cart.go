package shop

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func AddToCart(reader *bufio.Reader) {
	fmt.Println("\n--- Добавление товара в корзину ---")

	fmt.Println("Введите название товара: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Println("Введите цену товара: ")
	priceStr, _ := reader.ReadString('\n')
	priceStr = strings.TrimSpace(priceStr)
	price, _ := strconv.Atoi(priceStr)

	fmt.Println("Введите количество товаров: ")
	qtyStr, _ := reader.ReadString('\n')
	qtyStr = strings.TrimSpace(qtyStr)
	qty, _ := strconv.Atoi(qtyStr)

	fmt.Printf("\nВы добавили в корзину: %s, количество: %d, цена за штуку: %d тенге",
		title, price, qty)
}
