package main

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	//"gorm.io/gorm/clause"
)

type WebStore interface {
	Create(Model) error
}

type webStore struct {
	db *gorm.DB
	ch chan Model
}

func (t *webStore) Create(response Model) error {
	var err error
	select {
	case t.ch <- response:
	default:
		var a []Model
		for i := 0; i < 50; i++ {
		a =	append(a,<-t.ch)
		}
		err =	t.db.Debug().Table("scrape").Clauses(clause.OnConflict{UpdateAll: true}).Create(&a).Error
		t.ch<-response
	}
	return err 

}

func NewWebStore(db *gorm.DB) WebStore {
	return &webStore{
		db: db,
		ch: make(chan Model, 50),
	}
}
