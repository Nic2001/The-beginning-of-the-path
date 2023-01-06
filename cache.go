package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
)

// Подгрузка кэша при запуске приложения
func restartCache(cacheServ *cache.Cache) *cache.Cache {

	fmt.Println("Cache restarting")
	//Создаём новый экземпляр кэша с бессрочным временем жизни
	cacheServ = cache.New(cache.DefaultExpiration, cache.DefaultExpiration)

	//Пытаемся загрузить файл кэша, созданный ранее
	err := cacheServ.LoadFile("cache")

	//Если не получилось - загружаем из бд
	if err != nil {
		fmt.Println("New cache")
		cacheServ = OrdersFromCache(cacheServ)
	}

	return cacheServ
}

// Заполнение кэша принятыми данными (в отдельной функции - чтоб проще использовать в других местах)
func addOrderCache(cacheServ *cache.Cache, order_id string, order_info string) *cache.Cache {

	//Добавляем информацию в кэш
	cacheServ.Set(order_id, order_info, cache.DefaultExpiration)

	return cacheServ
}
