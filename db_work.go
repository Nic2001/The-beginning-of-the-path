package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/patrickmn/go-cache"
)

// Строка для коннектинга с бд
var ConnectStrDataBase string = "host=localhost port=5432 user=postgres password=123456 dbname=postgres sslmode=disable"

// Функция добавления заказа в бд
func addOrderDB(idData string, data string) {
	fmt.Println("Add order data in data base")

	//Открываем бд с помощью строки коннектинга
	db, err := sql.Open("postgres", ConnectStrDataBase)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//Выполняем запрос
	db.Exec(fmt.Sprintf(`INSERT INTO "postgres" ("order_id", "order_info") VALUES ('%s', '%s')`, idData, data))

}

// Выгрузка данных из бд в кэш
func OrdersFromCache(cacheServ *cache.Cache) *cache.Cache {

	fmt.Println("Load cache from db")
	//Открываем бд
	db, err := sql.Open("postgres", ConnectStrDataBase)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//Количество строк
	count := 0
	db.QueryRow(`SELECT count(*) FROM postgres`).Scan(&count)

	//Пустые строки для id заказа и информации о нём
	ord_id := ""
	ord_info := ""

	//В цикле заполняем кэш значениями из бд
	for i := 1; i <= count; i++ {
		db.QueryRow(fmt.Sprintf(`SELECT order_id FROM postgres WHERE order_key = '%d'`, i)).Scan(&ord_id)
		db.QueryRow(fmt.Sprintf(`SELECT order_info FROM postgres WHERE order_key = '%d'`, i)).Scan(&ord_info)
		cacheServ.Set(ord_id, ord_info, cache.DefaultExpiration)
	}

	//Сохраняем файл кэша
	cacheServ.SaveFile("cache")

	return cacheServ
}
