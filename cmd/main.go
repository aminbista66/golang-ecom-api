package main

import (
	"database/sql"
	"log"

	"github.com/aminbista66/golang-ecom-api/cmd/api"
	"github.com/aminbista66/golang-ecom-api/config"
	"github.com/aminbista66/golang-ecom-api/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMYSQLStorage(mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("DB: Database connected")
}