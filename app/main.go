package main

import (
	"money-manager/app/db"
	"money-manager/app/server"
)

func main() {
	db.ConnectDB()
	server.StartServer()
}
