package main

import (
	"fmt"
	"log"
	"time"

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
		log.Fatal("cannot load config")
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%d dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	web := web.NewWeb()
	webStore := store.NewWebStore(db)
	api.NewWebService(webStore, web)
	log.Println("Listening on", "8080")

	// go func(){
	// defer wg.Done()
	//http.ListenAndServe(":8080", api.MakeHandler(scrap))
	<-time.NewTicker(time.Hour).C
	log.Println(time.Now())
	// }()
	// wg.Wait()
}
