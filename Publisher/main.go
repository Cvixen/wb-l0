package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"strconv"
)

const (
	clusterName = "test-cluster"
	clientName  = "publisher"
	channel     = "orders"
)

/**
Создает 10 экземпляров ордеров и отправляет в натс.
*/

func main() {
	order := CreateStruct()
	sc, err := stan.Connect(clusterName, clientName)
	defer sc.Close()
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		orderId := i
		order.OrderUid = fmt.Sprintf("b563feb7b2b84b6test" + strconv.Itoa(orderId))
		jsonData, err := json.Marshal(order)
		if err != nil {
			log.Fatal(err)
		}
		err = sc.Publish(channel, jsonData)
		if err != nil {
			log.Fatalf("Error during publish: %v\n", err)
		}
	}
	log.Println("Publisher is done")
}
