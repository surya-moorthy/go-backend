package main

import (
	"log"

	"gihtub.com/go-backend/go-ecommerce/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080",nil)
	if err := server.Run(); err != nil {
       log.Fatal(err)
	}

}