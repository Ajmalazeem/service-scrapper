package api

import (
	"log"

	"bitbucket.org/ajmal_azm/scraperP/scrap"
	"bitbucket.org/ajmal_azm/scraperP/store"
	"bitbucket.org/ajmal_azm/scraperP/web"
)

type ScraperBg interface {
	GeneratePackage()
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
	for i := 0; i < 24; i++ {
		go a.Scrap()
	}
	return &a
}

func (t *Scraperbg) GeneratePackage()  {
	go t.web.CategoriesList()
	go t.web.UrlList()
	go t.web.Searcher()
}

func (t *Scraperbg) Scrap() (err error) {

	for url := range t.web.PackageNameChan() {
		response := scrap.Scraper(url)
		if err := t.webStore.Create(response); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
