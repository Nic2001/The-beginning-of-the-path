package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/patrickmn/go-cache"
)

// Структура представления данных о заказе
type ViewData struct {
	OrderId   string
	OrderInfo string
}

// Запуск сервера
func serverHtmlStart(cacheServ *cache.Cache) {
	//Функция маршрутизации страницы запроса
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//Шаблон страницы запроса
		http.ServeFile(w, r, "templates/index.html")
	})

	//Функция маршрутизации страницы ответа
	http.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {
		//Получение данных из поля orderId
		orderId := r.FormValue("orderId")

		//Забираем данные по id из кэша
		orderData, flag := cacheServ.Get(orderId)

		//Экземпляр структуры представления данных о заказе
		var data ViewData

		//Если такой id в кэше нашёлся, то предоставляем данные
		if flag {

			//Строка со всеми данными о заказе
			var orderInfo string = fmt.Sprintf("%+v", orderData)

			//Заполняем форму для ответа
			data = ViewData{
				OrderId:   orderId,
				OrderInfo: orderInfo,
			}

		} else {
			//Если такого id в кэше нет, то выводим сообщение об этом
			data = ViewData{
				OrderId:   "нет информации",
				OrderInfo: "нет информации",
			}
		}

		//Шаблон страницы ответа
		tmpl, _ := template.ParseFiles("templates/orderID.html")

		//Передаём шаблону данные, генерируем разметку страницы ответа в ответ на запрос
		tmpl.Execute(w, data)

	})

	//Запуска веб приложения по адресу localhost:3000
	http.ListenAndServe(":3000", nil)
}
