package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Orders struct {
	Orders []Order `json:"orders"`
}

type Order struct {
	DAO      string  `json:"dao"`
	Price    float64 `json:"price"`
	Quantity float64 `json:quantity`
	APR      float64 `json:apr`
}

func showOrders(orders Orders) {

	fmt.Println("\nORDER HISTORY")
	fmt.Println("______________")

	for i := 0; i < len(orders.Orders); i++ {
		fmt.Printf("\nDAO: %12s\n", orders.Orders[i].DAO)
		fmt.Printf("Price: %10.4f\n", orders.Orders[i].Price)
		fmt.Printf("Quantity: %7.4f\n", orders.Orders[i].Quantity)
		fmt.Printf("_________________\n")
	}
}

func calcProjections(order Order) {

	fmt.Println("\nPROJECTIONS BY 5-DAY APR")
	fmt.Println("________________________")

	day := 0
	rebase := 1
	dao := order.DAO
	price := order.Price
	quantity := order.Quantity
	apr := order.APR

	value := price * quantity

	updatedQuantity := quantity
	newQuan := quantity * 2

	fmt.Printf("\nDay %-5d %10.4f\n", day, price*updatedQuantity)
	for i := 0; i < 1095; i++ {

		fmt.Printf("\nRebase %-5d %12.2f\n", rebase, updatedQuantity)

		rebase += 1
		updatedQuantity += updatedQuantity * apr
		newQuan += newQuan * apr

		if rebase == 4 {
			day += 1
			rebase = 1
			fmt.Printf("\nDay %-5d %10.4f\n", day, price*updatedQuantity)
		}

	}
	fmt.Println("________________________")

	updatedValue := price * updatedQuantity
	value2 := price * newQuan

	fmt.Printf("_________________________________________\n")
	fmt.Printf("\nDAO: %12s\n", dao)
	fmt.Printf("Price: %10.4f\n", price)
	fmt.Printf("Quantity: %7.2f    ->    %5.2f   ->    %5.2f\n", quantity, updatedQuantity, newQuan)
	fmt.Printf("Value: %10.2f    ->    %5.2f   ->    %5.2f\n", value, updatedValue, value2)
	fmt.Printf("\n_________________________________________\n")

}

func showProjections(orders Orders) {

	option := 1
	for option > 0 {

		fmt.Println("\nEnter option:")
		fmt.Print("____________\n\n1. ICE\n2. JADE\n3. ROME\n0. Exit\n____________\n: ")
		fmt.Scanln(&option)

		fmt.Print(len(orders.Orders))

		if option > 0 && option <= len(orders.Orders) {
			calcProjections(orders.Orders[option-1])
		}

	}
	fmt.Println("\nBack to menu...")
}

func main() {

	// 1. Reading the JSON file
	ordersFile, err := os.Open("./data/orders.json")
	if err != nil {
		log.Fatal("Error: Cannot open file - ", err)
	}
	defer ordersFile.Close()

	// 2. Unmarshal the JSON

	// read opened ordersFile as a byte array
	byteValue, _ := ioutil.ReadAll(ordersFile)

	// initialize orders array
	orders := Orders{}

	// unmarshal byteArray with ordersFile into orders
	err = json.Unmarshal(byteValue, &orders)

	option := 1
	for option > 0 {

		fmt.Println("\nEnter option:")
		fmt.Print("________________\n\n1. Order History\n2. Projections\n0. Exit\n________________\n: ")
		fmt.Scanln(&option)

		switch option {
		case 0:
			fmt.Println("\nGoodbye!")
		case 1:
			showOrders(orders)
		case 2:
			showProjections(orders)
		default:
			fmt.Println("\nError: Invalid option")
		}
	}
}
