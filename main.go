package main

import (
	"TransportManagementService/db"
	"TransportManagementService/server"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("🚀 Запуск InitDB...")
	err := db.InitDB("root", "vdySqAwCIwMHUfdUyqaQlBOBlCrZovdD", "centerbeam.proxy.rlwy.net:3885", "railway")

	if err != nil {
		log.Println("Не удалось подключиться к базе данных")
		return
	}

	log.Println("БД успешно подключена")
	server.Start()
}
