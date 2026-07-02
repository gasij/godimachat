package main

import (
	"context"
	"fmt"
	"godima/internal/config"
	"godima/internal/database"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	pool, err := database.Connect(context.Background(), cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	fmt.Println("БД подключена!")
}
