package main

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	//"gorm.io/gorm/clause"
)

type WebStore interface {
	Create(Model) error
	GetP(req GetRequest) (*Model, error)
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
		unique := map[string]struct{}{}
		for i := 0; i < 50; i++ {
			val := <-t.ch
			if _, ok := unique[val.PackageName]; !ok {
				a = append(a, val)
			}
			unique[val.PackageName] = struct{}{}
		}

		err = t.db.Debug().Table("scrape").Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "package_name"}},
			UpdateAll: true,
		}).Create(&a).Error
		t.ch <- response
	}
	return err

}

func (t *webStore) GetP(req GetRequest) (*Model, error) {
	result := Model{}
	err := t.db.Table("scrape").Where("package_name = ?", req.PackageName).Find(&result).Error

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func NewWebStore(db *gorm.DB) WebStore {
	return &webStore{
		db: db,
		ch: make(chan Model, 50),
	}
}
