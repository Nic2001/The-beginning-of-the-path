package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"github.com/patrickmn/go-cache"
	"time"
)

var cacheServ *cache.Cache //Кэш

func main() {

	cacheServ = restartCache(cacheServ) //Подгружаем кэш

	go subscribe() //Подписка на канал

	go serverHtmlStart(cacheServ) //Запускаем сервер

	fmt.Scanln() //Ждём ввода id для поиска
}

func subscribe() { //Функция подписки на канал
	//Подключение
	fmt.Println("Connect channel")
	nc, err := nats.Connect("nick@localhost")
	if err != nil {
		println("Not connect channel")
		time.Sleep(2 * time.Second)
		subscribe()
	}
	ec, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	//Принимающий канал
	recvCh := make(chan orderModel)

	//Синхронизация каналов
	ec.BindRecvChan("order", recvCh)

	//Вытаскиваем данные из канала и разбираем их
	order := <-recvCh
	order_id := order.OrderUID
	order_data, err := json.Marshal(order)
	if err != nil {
		fmt.Println(err)
	}
	order_info := string(order_data) //Так проще
	//order_info := fmt.Sprintf("Трек-номер: '%s', запись: '%s', Доставка: имя: '%s', телефон: '%s', почтовый индекс: '%s' город: '%s' адрес: '%s' регион: '%s' электронная почта: '%s' Оплата: транзакция: '%s' идентификатор: '%s' валюта: '%s' провайдер: '%s' сумма: '%s' платёж: '%s' банк: '%s' стоимость доставки: '%s' количество товаров: '%s' таможенный сбор: '%s' Предметы: '%s' Локация: '%s' внутренняя подпись: '%s' идентификатор: '%s' сервис доставки: '%s' ключ: '%s' id: '%s' дата заказа: '%s' ещё что то: '%s'", order.TrackNumber, order.Entry, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDt, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee, order.Items, order.Local, order.InternalSignature, order.CustomerID, order.DeliveryService, order.Shardkey, order.SmID, order.DateCreated, order.OofShard)

	//Добавляем в кэш и в бд
	addOrderCache(cacheServ, order_id, order_info)
	addOrderDB(order_id, order_info)

	//fmt.Println(order_id, order_info)

	//Отключение
	nc.Close()
	ec.Close()
}
