package main

import (
	"log"

	"gihtub.com/go-backend/go-ecommerce/cmd/api"
	"gihtub.com/go-backend/go-ecommerce/config"
	"gihtub.com/go-backend/go-ecommerce/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

	db , err := db.NewMySQLStorage(mysql.Config{
		User : config.Envs.DBUser,
		Passwd : config.Envs.DBPassword,
		Addr : config.Envs.DBAddress,
		DBName : config.Envs.DBName,
		Net   : "tcp",
		AllowNativePasswords : true,
		ParseTime : true,
	})

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080",db)
	if err := server.Run(); err != nil {
       log.Fatal(err)
	}

}