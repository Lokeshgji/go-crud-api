package main

import (
	"github.com/Lokeshgji/go-crud-api/db"
	"github.com/Lokeshgji/go-crud-api/router"
)

func main() {
	db.InitPostgresDB()
	router.InitRouter().Run()
}
