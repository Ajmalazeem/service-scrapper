package scrap

import (
	"log"
	"net/http"

	"strings"

	"bitbucket.org/ajmal_azm/scraperP/model"
	"github.com/PuerkitoBio/goquery"
)

func Scraper(url string) model.Model {

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println(err)
	}

	response := model.Model{}

	packageName := url
	packageName = strings.TrimPrefix(packageName, "https://play.google.com/store/apps/details?id=")
	response.PackageName = packageName

	response.Url = url

	doc.Find(".AHFaub").Each(func(i int, v *goquery.Selection) {
		response.AppName = v.Text()
		 log.Println(response.AppName)
	})

	doc.Find(".R8zArc").Each(func(i int, f *goquery.Selection) {

		if i == 0 {
			response.DeveloperName = f.Text()
			return
		}

	})

	doc.Find("div > div.hkhL9e > div > img").Each(func(i int, s *goquery.Selection) {
		var ok bool
		response.ImageUrl, ok = s.Attr("src")
		if !ok {
			log.Fatal(err)
		}
	})

	doc.Find("div.BHMmbe").Each(func(i int, t *goquery.Selection) {

		response.Rating = t.Text()
	})

	doc.Find("div > div.D0ZKYe > div > div.sIskre > div.bSIuKf").Each(func(i int, b *goquery.Selection) {
		var c bool = true
		c = b.Text() == "Contains Ads·Offers in-app purchases" //"Offers in-app purchases"
		if c {

			response.ContainAds = c
			response.InAppPurchase = c

		} else {
			if b.Text() == "Offers in-app purchases" {
				c = true
				response.InAppPurchase = c

			} else if b.Text() == "Contains Ads" {
				c = true
				response.ContainAds = c
			} else {
				return
			}
		}
	})
	//---------------------ratedpeoplecount------------------------
	doc.Find("span.hzfjkd").Each(func(i int, s *goquery.Selection) {
		response.RatedPeopleCount = s.Next().Text()
	})

	//------------updated-----------------------------------------------
	doc.Find(" div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(1) > div").Each(func(i int, s *goquery.Selection) {
		if s.Text() == "Updated" {

			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(1) > span > div > span").Each(func(i int, t *goquery.Selection) {

				response.Updated = t.Text()
			})
		} else {
			return
		}
	})

	//------------------------size---------------------------------
	doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(2) > div").Each(func(i int, s *goquery.Selection) {

		if s.Text() == "Size" {
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(2) > span > div > span").Each(func(i int, t *goquery.Selection) {

				response.Size = t.Text()
			})
		}
	})

	//------------------------------installs---------------------------
	doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(3) > div").Each(func(i int, s *goquery.Selection) {

		if s.Text() == "Installs" {

			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(3) > span > div > span").Each(func(i int, t *goquery.Selection) {

				response.Installs = t.Text()

			})
		} else {
			return
		}
	})

	//------------------------current version-----------------------
	doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(4) > div").Each(func(i int, s *goquery.Selection) {

		if s.Text() == "Current Version" {

			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(4) > span > div > span").Each(func(i int, t *goquery.Selection) {

				response.CurrentVersion = t.Text()

			})
		} else {
			return
		}
	})

	//------------------------------android-----------------------
	doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(5) > div").Each(func(i int, s *goquery.Selection) {
		if s.Text() == "Requires Android" {
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(5) > span > div > span").Each(func(i int, t *goquery.Selection) {

				response.AndroidVersion = t.Text()
			})
		} else {
			return
		}
	})

	//---------------------------content rating----------------------------------
	doc.Find("div > div:nth-child(6) > div.BgcNfc").Each(func(i int, s *goquery.Selection) {
		if s.Text() == "Content Rating" {

			doc.Find("div:nth-child(6) > span > div > span.htlgb").Each(func(i int, t *goquery.Selection) {

				response.ContentRating = t.Text()
			})
		} else {
			return
		}
	})
	// --------------------------------------------------------------------------------
	//#1
	doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(7) > div").Each(func(i int, s *goquery.Selection) {
		switch s.Text() {
		case "Interactive Elements":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(7) > span > div > span").Each(func(i int, t *goquery.Selection) {

				response.InteractiveElements = t.Text()

			})
		case "In-app Products":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(7) > span > div > span").Each(func(i int, t *goquery.Selection) {

				response.InAppProducts = t.Text()

			})
		case "Offered By":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(7) > span > div > span").Each(func(i int, t *goquery.Selection) {

				response.OfferedBy = t.Text()

			})
		case "Developer":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(7) > span > div > span").Each(func(i int, t *goquery.Selection) {

				response.Developer = t.Text()

			})
		}
	})

	//---------------------------------------------------------------------------------------
	// #2
	doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(8) > div").Each(func(i int, s *goquery.Selection) {
		switch s.Text() {
		case "Interactive Elements":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(8) > span > div > span").Each(func(i int, t *goquery.Selection) {

				response.InteractiveElements = t.Text()

			})
		case "In-app Products":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(8) > span > div > span").Each(func(i int, t *goquery.Selection) {

				response.InAppProducts = t.Text()

			})
		case "Offered By":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(8) > span > div > span").Each(func(i int, t *goquery.Selection) {

				response.OfferedBy = t.Text()

			})
		case "Developer":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(8) > span > div > span").Each(func(i int, t *goquery.Selection) {

				response.Developer = t.Text()

			})
		}
	})

	//---------------------------------------------------------------------------------------
	// #3
	doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(9) > div").Each(func(i int, s *goquery.Selection) {
		switch s.Text() {
		case "Interactive Elements":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(9) > span > div > span").Each(func(i int, t *goquery.Selection) {

				response.InteractiveElements = t.Text()

			})
		case "In-app Products":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(9) > span > div > span").Each(func(i int, t *goquery.Selection) {

				response.InAppProducts = t.Text()

			})
		case "Offered By":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(9) > span > div > span").Each(func(i int, t *goquery.Selection) {

				response.OfferedBy = t.Text()

			})
		case "Developer":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(9) > span > div > span").Each(func(i int, t *goquery.Selection) {

				response.Developer = t.Text()

			})
		}
	})

	//---------------------------------------------------------------------------------------
	// #4
	doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(10) > div").Each(func(i int, s *goquery.Selection) {
		switch s.Text() {
		case "Interactive Elements":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(10) > span > div > span.htlgb").Each(func(i int, t *goquery.Selection) {

				response.InteractiveElements = t.Text()

			})
		case "In-app Products":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(10) > span > div > span.htlgb").Each(func(i int, t *goquery.Selection) {

				response.InAppProducts = t.Text()

			})
		case "Offered By":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(10) > span > div > span.htlgb").Each(func(i int, t *goquery.Selection) {

				response.OfferedBy = t.Text()

			})
		case "Developer":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(10) > span > div > span.htlgb").Each(func(i int, t *goquery.Selection) {

				response.Developer = t.Text()

			})
		}
	})

	//---------------------------------------------------------------------------------------
	// #5
	doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(11) > div").Each(func(i int, s *goquery.Selection) {
		switch s.Text() {
		case "Interactive Elements":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(11) > span > div > span").Each(func(i int, t *goquery.Selection) {

				response.InteractiveElements = t.Text()

			})
		case "In-app Products":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(11) > span > div > span").Each(func(i int, t *goquery.Selection) {

				response.InAppProducts = t.Text()
			})
		case "Offered By":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(11) > span > div > span").Each(func(i int, t *goquery.Selection) {
				response.OfferedBy = t.Text()
			})
		case "Developer":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(11) > span > div > span").Each(func(i int, t *goquery.Selection) {
				response.Developer = t.Text()
			})
		}
	})

	//---------------------------------------------------------------------------------------
	// #6
	doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(12) > div").Each(func(i int, s *goquery.Selection) {
		switch s.Text() {
		case "Interactive Elements":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(12) > span > div > span").Each(func(i int, t *goquery.Selection) {
				response.InteractiveElements = t.Text()
			})
		case "In-app Products":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(12) > span > div > span").Each(func(i int, t *goquery.Selection) {
				response.InAppProducts = t.Text()
			})
		case "Offered By":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(12) > span > div > span").Each(func(i int, t *goquery.Selection) {
				response.OfferedBy = t.Text()
			})
		case "Developer":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(12) > span > div > span").Each(func(i int, t *goquery.Selection) {
				response.Developer = t.Text()
			})
		}
	})
	//-------------------------------------------------------------------

	//pass

	return response

}
