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
	DAO      string `json:"dao"`
	Price    uint32 `json:"price"`
	Quantity uint32 `json:quantity`
}

func showOrders(orders Orders) {

	fmt.Println("\nORDER HISTORY")
	fmt.Println("______________")

	for i := 0; i < len(orders.Orders); i++ {
		fmt.Println("\nDAO: " + orders.Orders[i].DAO)
		fmt.Println("Price: ", orders.Orders[i].Price)
		fmt.Println("Quantity: ", orders.Orders[i].Quantity)
		fmt.Println("______________")
		fmt.Println("TEST")
	}
}

func main() {

	// 1. Reading the JSON file
	ordersFile, err := os.Open("./data/orders.json")
	if err != nil {
		log.Fatal("Error: cannot open file - ", err)
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
		fmt.Print("____________\n\n0. Exit\n1. Order History\n2. Continue\n____________\n: ")
		fmt.Scanln(&option)

		switch option {
		case 0:
			fmt.Println("\nGoodbye")
		case 1:
			showOrders(orders)
		case 2:
			fmt.Println("\nOption 2")
		default:
			fmt.Println("\nError: option out of bounds, please pick a number between 0 - 2.")
		}
	}
}
