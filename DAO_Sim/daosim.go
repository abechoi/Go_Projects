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
		fmt.Println("\nDAO: " + orders.Orders[i].DAO)
		fmt.Println("Price: ", orders.Orders[i].Price)
		fmt.Println("Quantity: ", orders.Orders[i].Quantity)
		fmt.Println("______________")
	}
}

func calcProjections(order Order) {

	fmt.Println("\nPROJECTIONS BY 5-DAY ROI")
	fmt.Println("_________________________")

	day := 0
	quantity := order.Quantity
	roi := order.ROI

	fmt.Println(quantity)

	for i := 0; i < 73; i++ {

		fmt.Printf("\nDay %-5d %f\n", day, quantity)

		day += 5
		quantity += quantity * roi
	}

}

func showProjections(orders Orders) {

	option := 1
	for option > 0 {

		fmt.Println("\nEnter option:")
		fmt.Print("____________\n\n1. ODAO\n2. OHM\n3. TIME\n4. ROME\n0. Exit\n____________\n: ")
		fmt.Scanln(&option)

		switch option {
		case 0:
			fmt.Println("\nBack to menu...")
		case 1:
			calcProjections(orders.Orders[option])
		case 2:
			calcProjections(orders.Orders[option])
		case 3:
			calcProjections(orders.Orders[option])
		case 4:
			calcProjections(orders.Orders[option])
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
		fmt.Print("____________\n\n1. Order History\n2. Projections\n0. Exit\n____________\n: ")
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
