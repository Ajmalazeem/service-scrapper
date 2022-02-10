package api

import (
	"bitbucket.org/ajmal_azm/scraperP/model"
	"bitbucket.org/ajmal_azm/scraperP/store"
)


type WebService interface {
	GetPackageNameDetails(req model.GetRequest) (*model.Model, error)
	GetChangeLogDetails(req model.GetRequest) (*[]model.Changelog, error)
}

type WebServices struct {
	webStore store.WebStore
}

func (t *WebServices) GetPackageNameDetails(req model.GetRequest) (*model.Model, error) {
	d, err := t.webStore.GetPackageNameDetails(req)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (t *WebServices) GetChangeLogDetails(req model.GetRequest) (*[]model.Changelog, error) {
	d, err := t.webStore.GetChangeLogDetails(req)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func NewWebService(webStore store.WebStore) WebService {
	a := WebServices{webStore: webStore}
	return &a
}


