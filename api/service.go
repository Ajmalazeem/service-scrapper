package api

import (
	"log"

	//"sync"

	//"github.com/chromedp/cdproto/cdp"

	"bitbucket.org/ajmal_azm/scraperP/model"
	"bitbucket.org/ajmal_azm/scraperP/scrap"
	"bitbucket.org/ajmal_azm/scraperP/store"
	"bitbucket.org/ajmal_azm/scraperP/web"
	//"bitbucket.org/ajmal_azm/scraperP/web"
)

type WebService interface {
	GeneratePackage() error
	Scrap() error
	GetPackageNameDetails(req model.GetRequest) (*model.Model, error)
	GetChangeLogDetails(req model.GetRequest) (*[]model.Changelog, error)
}

type WebServices struct {
	web      web.Web
	webStore store.WebStore
	b        chan string
	//a []string

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

func NewWebService(webStore store.WebStore, web web.Web) WebService {
	a := WebServices{webStore: webStore, web: web, b: make(chan string)}
	// var wg sync.WaitGroup

	// wg.Add(4)

	go func() {
		// defer wg.Done()
		a.GeneratePackage()
	}()
	for i := 0; i < 25; i++ {
		go func() {
			// defer wg.Done()
			a.Scrap()
		}()
	}
	// wg.Wait()
	return &a
}

func (t *WebServices) GeneratePackage() error {
	go t.web.CategoriesList()
	go t.web.UrlList()
	go t.web.Searcher()

	return nil
}

func (t *WebServices) Scrap() (err error) {

	for url := range t.web.PackageNameChan() {
		response := scrap.Scraper(url)
		if err := t.webStore.Create(response); err != nil {
			log.Fatal(err)
		}
	}
	return err
}

//-------------------------------------------------------------------------------
