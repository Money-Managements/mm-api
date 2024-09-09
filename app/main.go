package main

import (
	"money-manager/app/db"
	"money-manager/app/server"
)

func main() {
	db := db.ConnectDB()
	server.StartServer(db)
}
