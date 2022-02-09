package store

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	//"gorm.io/gorm/clause"
	"bitbucket.org/ajmal_azm/scraperP/model"

)

type WebStore interface {
	Create(model.Model) error
	GetPackageNameDetails(req model.GetRequest) (*model.Model, error)
	GetChangeLogDetails(req model.GetRequest) (*[]model.Changelog, error)
}

type webStore struct {
	db *gorm.DB
	ch chan model.Model
}

func (t *webStore) Create(response model.Model) error {
	var err error
	select {
	case t.ch <- response:
	default:
		var a []model.Model
		unique := map[string]struct{}{}
		for i := 0; i < 250; i++ {
			val := <-t.ch
			if _, ok := unique[val.PackageName]; !ok {
				a = append(a, val)
			}
			unique[val.PackageName] = struct{}{}
		}

		err = t.db.Table("scrape").Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "package_name"}},
			UpdateAll: true,
		}).Create(&a).Error
		t.ch <- response
	}
	return err
}

func (t *webStore) GetPackageNameDetails(req model.GetRequest) (*model.Model, error) {
	result := model.Model{}
	err := t.db.Table("scrape").Where("package_name = ?", req.PackageName).Find(&result).Error

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (t *webStore) GetChangeLogDetails(req model.GetRequest) (*[]model.Changelog, error) {
	result := []model.Changelog{}
	err := t.db.Table("scrape_logs").Where("package_name = ?", req.PackageName).Find(&result).Limit(-1).Error

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func NewWebStore(db *gorm.DB) WebStore {
	return &webStore{
		db: db,
		ch: make(chan model.Model, 250),
		
	}
}
