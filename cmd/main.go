package main

import (
	"log"

	"github.com/edbertjk/server-api/cmd/api"
)

func main(){
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil{
		log.Fatal(err)
	}
}