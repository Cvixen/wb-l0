package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/nats-io/stan.go"
	"log"
	"net/http"
	"wb-l0/Model"
	"wb-l0/Postgresql"
	"wb-l0/Utils"
)

const (
	clusterName = "test-cluster"
	clientName  = "subscriber"
	channel     = "orders"
)

/**
Записывает в кэш из бд. Подписывается на натс и полученные данные проверяет. В случае успешной проверки,
добавляет в бд и в кэш. После отправляет в handler значение кэша. Запуск сервера http.
*/
func main() {
	cache := make(map[string]Model.Order)
	Postgresql.FromDbToCash(&cache)
	handler := Model.NewHandler(cache)
	sc, _ := stan.Connect(clusterName, clientName)
	sub, _ := sc.Subscribe(channel, func(msg *stan.Msg) {
		var order Model.Order
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Print(err)
			return
		}
		err = Utils.ValidateStruct(&order)
		if err != nil {
			log.Print(err)
			return
		}
		Postgresql.OrderPutDb(&order, &cache)
		cache[order.OrderUid] = order
		handler = Model.NewHandler(cache)

	})
	router := httprouter.New()
	handler.Register(router)
	log.Fatal(http.ListenAndServe(":1234", router))
	sub.Unsubscribe()

}
