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
	Price    float32 `json:"price"`
	Quantity float32 `json:quantity`
	ROI      float32 `json:roi`
}

func showOrders(orders Orders) {

	fmt.Println("\nORDER HISTORY")
	fmt.Println("______________")

	for i := 0; i < len(orders.Orders); i++ {
		fmt.Printf("\nDAO: %12s\n", orders.Orders[i].DAO)
		fmt.Printf("Price: %10.2f\n", orders.Orders[i].Price)
		fmt.Printf("Quantity: %7.2f\n", orders.Orders[i].Quantity)
		fmt.Printf("_________________\n")
	}
}

func calcProjections(order Order) {

	fmt.Println("\nPROJECTIONS BY 5-DAY ROI")
	fmt.Println("________________________")

	day := 0
	dao := order.DAO
	price := order.Price
	quantity := order.Quantity
	roi := order.ROI

	value := price * quantity
	updatedQuantity := quantity

	for i := 0; i < 73; i++ {

		fmt.Printf("\nDay %-5d %12.2f\n", day, updatedQuantity)

		day += 5
		updatedQuantity += updatedQuantity * roi
	}
	fmt.Println("________________________")

	updatedValue := price * updatedQuantity

	fmt.Printf("_________________________________________\n")
	fmt.Printf("\nDAO: %12s\n", dao)
	fmt.Printf("Price: %10.2f\n", price)
	fmt.Printf("Quantity: %7.2f    ->    %5.2f\n", quantity, updatedQuantity)
	fmt.Printf("Value: %10.2f    ->    %5.2f\n", value, updatedValue)
	fmt.Printf("\n_________________________________________\n")

}

func showProjections(orders Orders) {

	option := 1
	for option > 0 {

		fmt.Println("\nEnter option:")
		fmt.Print("____________\n\n1. ODAO\n2. OHM\n3. TIME\n4. ROME\n5. BBY\n0. Exit\n____________\n: ")
		fmt.Scanln(&option)

		switch option {
		case 0:
			fmt.Println("\nBack to menu...")
		case 1:
			calcProjections(orders.Orders[option-1])
		case 2:
			calcProjections(orders.Orders[option-1])
		case 3:
			calcProjections(orders.Orders[option-1])
		case 4:
			calcProjections(orders.Orders[option-1])
		case 5:
			calcProjections(orders.Orders[option-1])
		default:
			fmt.Println("\nError: Invalid option")
		}
	}
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
