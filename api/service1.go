package api

import (
	"log"

	"bitbucket.org/ajmal_azm/scraperP/scrap"
	"bitbucket.org/ajmal_azm/scraperP/store"
	"bitbucket.org/ajmal_azm/scraperP/web"
)

type ScraperBg interface {
	GeneratePackage() error
	Scrap() error
}

type Scraperbg struct {
	web      web.Web
	webStore store.WebStore
	b        chan string
}

func NewScraperBg(webStore store.WebStore, web web.Web) ScraperBg {
	a := Scraperbg{webStore: webStore, web: web, b: make(chan string)}
	go a.GeneratePackage()
	for i := 0; i < 30; i++ {
		go a.Scrap()
	}
	return &a
}

func (t *Scraperbg) GeneratePackage() error {
	go t.web.CategoriesList()
	go t.web.UrlList()
	go t.web.Searcher()

	return nil
}

func (t *Scraperbg) Scrap() (err error) {

	for url := range t.web.PackageNameChan() {
		response := scrap.Scraper(url)
		if err := t.webStore.Create(response); err != nil {
			log.Println(err)
		}
	}
	return err
}
