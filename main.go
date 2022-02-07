package main

import (
	"log"
	"net/http"
	//"time"
	//"net/http"

	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	//"github.com/go-kit/kit/transport/http"
)


func main() {
	
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config")
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%d dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	webStore := NewWebStore(db)
	scrap := NewWebService(webStore)
	log.Println("Listening on", "8000")
	http.ListenAndServe(":8000", MakeHandler(scrap))
}




