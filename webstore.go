package main

import (
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WebStore interface {
	Details(DetailsResponse) error
}

type webStore struct {
	db *gorm.DB
}

func (t *webStore) Details(response DetailsResponse) error {
	log.Println(response)
	return t.db.Table("scrap").Clauses(clause.OnConflict{DoNothing: true}).Create(&response).Error

}


func NewWebStore(db *gorm.DB) WebStore {
	return &webStore{
		db: db,
	}
}
