package main

import (
	"money-manager/app/db"
	"money-manager/app/server"
	"money-manager/app/service"
)

func main() {
	db := db.ConnectDB()
	service.SetServices(db)
	server.StartServer()
}
