package main

import (
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	"math/rand"
	"time"
)

func main() {
	var token string
	fmt.Printf("Перейдите на https://vkhost.github.io \nНажмите Настройки и обязательно выберите Сообщения\nНажмите получить\nНажмите Разрешить\nСкопируйте все что между 'access_token=' и '&'\nВведите этот код ниже\n")
	fmt.Scanln(&token)
	rand.Seed(time.Now().UnixNano())
	vk := api.NewVK(token)
	res, _ := vk.UsersGet(api.Params{})
	id := res[0].ID
	for i := 0; i < 100000; i++ {
		params := api.Params{
			"user_id":   id,
			"random_id": rand.Uint32(),
			"peer_id":   id,
			"message":   "mm",
		}
		_, e := vk.MessagesSend(params)
		fmt.Println(e)
	}

}
