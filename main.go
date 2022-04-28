package main

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	//"github.com/go-kit/kit/transport/http"

	"bitbucket.org/ajmal_azm/scraperP/api"
	"bitbucket.org/ajmal_azm/scraperP/store"
	"bitbucket.org/ajmal_azm/scraperP/web"
)

func main() {

	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	web := web.NewWeb()
	webStore := store.NewWebStore(db)
	api.NewScraperBg(webStore, web)
	scrap := api.NewWebService(webStore)
	log.Println("Listening on", "8000")
	http.ListenAndServe(":8000", api.MakeHandler(scrap))
	//log.Println(http.ListenAndServe(":8000", api.MakeHandler(scrap)))
	// <-time.NewTicker(48*time.Hour).C
	// log.Println(time.Now())

}
