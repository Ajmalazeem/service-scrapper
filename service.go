package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)


type WebService interface {
	urlP (b chan string)(error)
	Details(b chan string)  error
}

type webService struct{
	webStore WebStore
}

var ErrEmpty = errors.New("empty string")


func (t *webService) urlP(b chan string)(error){
	catogories := [55]string{"https://play.google.com/store/apps","https://play.google.com/store/apps/top","https://play.google.com/store/apps/new","https://play.google.com/store/apps/category/ART_AND_DESIGN","https://play.google.com/store/apps/stream/baselist_featured_arcore","https://play.google.com/store/apps/category/BOOKS_AND_REFERENCE","https://play.google.com/store/apps/category/BUSINESS","https://play.google.com/store/apps/category/COMICS","https://play.google.com/store/apps/category/COMMUNICATION","https://play.google.com/store/apps/category/DATING","https://play.google.com/store/apps/stream/vr_top_device_featured_category","https://play.google.com/store/apps/category/EDUCATION","https://play.google.com/store/apps/category/ENTERTAINMENT","https://play.google.com/store/apps/category/FINANCE","https://play.google.com/store/apps/category/FOOD_AND_DRINK","https://play.google.com/store/apps/category/HEALTH_AND_FITNESS","https://play.google.com/store/apps/category/LIBRARIES_AND_DEMO","https://play.google.com/store/apps/category/LIFESTYLE","https://play.google.com/store/apps/category/MAPS_AND_NAVIGATION","https://play.google.com/store/apps/category/MEDICAL","https://play.google.com/store/apps/category/MUSIC_AND_AUDIO","https://play.google.com/store/apps/category/NEWS_AND_MAGAZINES","https://play.google.com/store/apps/category/PERSONALIZATION","https://play.google.com/store/apps/category/PHOTOGRAPHY","https://play.google.com/store/apps/category/PRODUCTIVITY","https://play.google.com/store/apps/category/SHOPPING","https://play.google.com/store/apps/category/SOCIAL","https://play.google.com/store/apps/category/TOOLS","https://play.google.com/store/apps/category/TRAVEL_AND_LOCAL","https://play.google.com/store/apps/category/VIDEO_PLAYERS","https://play.google.com/store/apps/category/ANDROID_WEAR","https://play.google.com/store/apps/category/WATCH_FACE","https://play.google.com/store/apps/category/WEATHER","https://play.google.com/store/apps/category/GAME","https://play.google.com/store/apps/category/GAME_ACTION","https://play.google.com/store/apps/category/GAME_ADVENTURE","https://play.google.com/store/apps/category/GAME_ARCADE","https://play.google.com/store/apps/category/GAME_BOARD","https://play.google.com/store/apps/category/GAME_CARD","https://play.google.com/store/apps/category/GAME_CASINO","https://play.google.com/store/apps/category/GAME_CASUAL","https://play.google.com/store/apps/category/GAME_EDUCATIONAL","https://play.google.com/store/apps/category/GAME_MUSIC","https://play.google.com/store/apps/category/GAME_PUZZLE","https://play.google.com/store/apps/category/GAME_RACING","https://play.google.com/store/apps/category/GAME_ROLE_PLAYING","https://play.google.com/store/apps/category/GAME_SIMULATION","https://play.google.com/store/apps/category/GAME_SPORTS","https://play.google.com/store/apps/category/GAME_STRATEGY","https://play.google.com/store/apps/category/GAME_TRIVIA","https://play.google.com/store/apps/category/GAME_WORD","https://play.google.com/store/apps/category/FAMILY","https://play.google.com/store/apps/category/FAMILY?age=AGE_RANGE1","https://play.google.com/store/apps/category/FAMILY?age=AGE_RANGE2","https://play.google.com/store/apps/category/FAMILY?age=AGE_RANGE3"}
	v:= "https://play.google.com"
	q :="https://play.google.com/store/apps/editors_choice"

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// var u string
	
	
	// t := make(chan string)

	//var b[]string
	var headerNodes []*cdp.Node
	err := chromedp.Run(ctx, 
		chromedp.Navigate(`https://play.google.com/store/apps`),
		chromedp.Nodes(":is(#action-dropdown-children-Categories > div > ul > li > ul > li.KZnDLd > a)", &headerNodes, chromedp.ByQueryAll),//catogories
		chromedp.ActionFunc(func(c context.Context) error {  
			for i, node := range headerNodes {
				a :=node.AttributeValue("href")
				//v:= "https://play.google.com"
				url:= v+a
				log.Println(url)
				log.Println(i)
				
				
				err := chromedp.Run(ctx,
					chromedp.Navigate(url),//cluster
					chromedp.Nodes(":is(div.wXUyZd > a)", &headerNodes, chromedp.ByQueryAll),
					chromedp.ScrollIntoView(`footer`),
					// chromedp.WaitVisible(":is(footer > div)"),
					chromedp.ActionFunc(func(c context.Context) error {
					for i, node := range headerNodes {
						//log.Printf("-------------------- %s ------------------ \n", node.NodeName)
						
						//log.Printf("Node: %+v \n", u)
							log.Println(i)
							log.Println("sending")
							u :=node.AttributeValue("href")
							j:= v + u
							log.Printf("Node: %+v \n", j)
							b <- j
							log.Println("sent")
							i++
						
						}
					// close(b)
					return nil
					}),
					
				)
				if err != nil {
							log.Fatal(err)
							return err
					}

				i++
			}
			// close(b)
			return nil
		}),

	)
	
	if err != nil {
		log.Fatal(err)
	}
	//catagories======================================================================================
	
	for j:=0;j<len(catogories);j++{
		log.Println(catogories[j])
	err := chromedp.Run(ctx,
		chromedp.Navigate(catogories[j]),
		chromedp.Nodes(":is(div.xwY9Zc > a)", &headerNodes, chromedp.ByQueryAll),
		chromedp.ScrollIntoView(`footer`),
			chromedp.ActionFunc(func(c context.Context) error {  
			for i, node := range headerNodes {
				a :=node.AttributeValue("href")
				
				url:= v+a
				log.Println(url)
				
				if url!=q{
				
				err := chromedp.Run(ctx,
							chromedp.Navigate(url),//cluster
							chromedp.Nodes(":is(div.wXUyZd > a)", &headerNodes, chromedp.ByQueryAll),
							chromedp.ScrollIntoView(`footer`),
							chromedp.ActionFunc(func(c context.Context) error {
							for i, node := range headerNodes {
								//log.Printf("-------------------- %s ------------------ \n", node.NodeName)
								
								//log.Printf("Node: %+v \n", u)
									log.Println(i)
									log.Println("sending")
									u :=node.AttributeValue("href")
									j:= v + u
									log.Printf("Node: %+v \n", j)
									b <- j
									log.Println("sent")
									i++
								
								}
							// close(b)
							return nil
							}),
							
						)
						if err != nil {
									log.Fatal(err)
									return err
								}
						continue
						}
	
	
				// t <- a
				// log.Printf("Node: %+v \n", a)
				i++
				
				}
			return nil                                                        
			}),
						)
	if err != nil {
		log.Fatal(err)
	}
	}
	close(b)

	// log.Println(u)
	// b <- u

	return err
}

func NewWebService(webStore WebStore) WebService{
	a:= webService{webStore: webStore}
	b := make(chan string)
	var wg sync.WaitGroup

	wg.Add(2)

	
	go func(){
		defer wg.Done()
		
		a.urlP(b)
	
		
	}()
	
	go func(chan string){
			defer wg.Done()
		
			a.Details(b)

			
	}(b)	
	

	
	wg.Wait()

	return &a
}


func (k *webService) Details(b chan string)(err error) {
	for url := range b {
	
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println(err)
	}

	response := DetailsResponse{}

	packageName := url
	packageName = strings.TrimPrefix(packageName, "https://play.google.com/store/apps/details?id=")
	response.PackageName = packageName
	
	
	response.Url = url

	doc.Find(".AHFaub").Each(func(i int, v *goquery.Selection) {

		response.APPNAME = v.Text()
		log.Println(response.APPNAME)		
	})

	doc.Find(".R8zArc").Each(func(i int, f *goquery.Selection) {

		if i == 0 {
			response.DEVELOPERNAME = f.Text()
			log.Println(response.DEVELOPERNAME)
			return
		}

	})

	doc.Find("div > div.hkhL9e > div > img").Each(func(i int, s *goquery.Selection){
		var ok bool
		response.ImageUrl, ok = s.Attr("src")
		if ok{
			log.Println(response.ImageUrl)
		}
	})

	doc.Find("div.BHMmbe").Each(func(i int, t *goquery.Selection) {
		
		response.Rating = t.Text()
		log.Println(response.Rating)
	})

	doc.Find("div > div.D0ZKYe > div > div.sIskre > div.bSIuKf").Each(func(i int, b *goquery.Selection) {
		var c bool = true
		//var j string
		
		log.Println(b.Text())
		c = b.Text() == "Contains Ads·Offers in-app purchases" //"Offers in-app purchases"
		if c {
			
			response.ContainAds = c
			response.InAppPurchase = c

		} else {
			if b.Text() == "Offers in-app purchases" {
				c= true
				response.InAppPurchase = c
			
			} else if b.Text() == "Contains Ads"{
				c = true
				response.ContainAds = c
			}else{
				return
			}
		}
	})
	//---------------------ratedpeoplecount------------------------
	doc.Find("span.hzfjkd").Each(func(i int, s *goquery.Selection){
		response.RatedPeopleCount = s.Next().Text()
		log.Println(response.RatedPeopleCount)
	})

	//------------updated-----------------------------------------------
	doc.Find(" div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(1) > div").Each(func(i int, s *goquery.Selection){
		log.Println(s.Text())
		if s.Text()=="Updated"{

			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(1) > span > div > span").Each(func(i int, t *goquery.Selection) {
			
				response.Updated = t.Text()
				log.Println(response.Updated)
			})
		}else{
			return 
		}
	})	

	//------------------------size---------------------------------
	doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(2) > div").Each(func(i int, s *goquery.Selection){
		log.Println(s.Text())
		if s.Text()=="Size"{
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(2) > span > div > span").Each(func(i int, t *goquery.Selection) {

				response.Size = t.Text()
				log.Println(response.Size)
			})
		}
	})

	//------------------------------installs---------------------------
	doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(3) > div").Each(func(i int, s *goquery.Selection){
		log.Println(s.Text())
		if s.Text()=="Installs"{

			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(3) > span > div > span").Each(func(i int, t *goquery.Selection) {
				
				response.Installs = t.Text()
				log.Println(response.Installs)
			})
		}else {
			return
		}
	})


	//------------------------current version-----------------------
	doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(4) > div").Each(func(i int, s *goquery.Selection){
		log.Println(s.Text())
		if s.Text()=="Current Version"{

			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(4) > span > div > span").Each(func(i int, t *goquery.Selection) {

				response.CurrentVersion = t.Text()
				log.Println(response.CurrentVersion)
				
			})
		}else {
			return
		}
	})	
	
	//------------------------------android-----------------------
	doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(5) > div").Each(func(i int, s *goquery.Selection){
		log.Println(s.Text())
		if s.Text()=="Requires Android"{
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(5) > span > div > span").Each(func(i int, t *goquery.Selection) {

				response.AndroidVersion = t.Text()
				log.Println(response.AndroidVersion)
			})
		}else {
			return
		}
	})	

	//---------------------------content rating----------------------------------
	doc.Find("div > div:nth-child(6) > div.BgcNfc").Each(func(i int, s *goquery.Selection){
		log.Println(s.Text())
		if s.Text()=="Content Rating"{

			doc.Find("div:nth-child(6) > span > div > span.htlgb").Each(func(i int, t *goquery.Selection) {

				response.ContentRating = t.Text()
				log.Println(response.ContentRating)
			})
		}else {
			return
		}
	})
	// --------------------------------------------------------------------------------
	//#1
	doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(7) > div").Each(func(i int, s *goquery.Selection){
		switch s.Text() {
		case "Interactive Elements":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(7) > span > div > span").Each(func(i int, t *goquery.Selection) {
				log.Println(s.Text())
				response.InteractiveElements = t.Text()
				log.Println(response.InteractiveElements)
			})
		case "In-app Products":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(7) > span > div > span").Each(func(i int, t *goquery.Selection) {
				log.Println(s.Text())
				response.InAppProducts = t.Text()
				log.Println(response.InAppProducts)
			})
		case "Offered By":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(7) > span > div > span").Each(func(i int, t *goquery.Selection) {
				log.Println(s.Text())
				response.OfferedBy = t.Text()
				log.Println(response.OfferedBy)
			})
		case "Developer":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(7) > span > div > span").Each(func(i int, t *goquery.Selection) {
				log.Println(s.Text())
				response.Developer = t.Text()
				log.Println(response.Developer)
			})
		}
	})

	
	//---------------------------------------------------------------------------------------
	// #2
	doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(8) > div").Each(func(i int, s *goquery.Selection){
		switch s.Text() {
		case "Interactive Elements":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(8) > span > div > span").Each(func(i int, t *goquery.Selection) {
				log.Println(s.Text())
				response.InteractiveElements = t.Text()
				log.Println(response.InteractiveElements)
			})
		case "In-app Products":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(8) > span > div > span").Each(func(i int, t *goquery.Selection) {
				log.Println(s.Text())
				response.InAppProducts = t.Text()
				log.Println(response.InAppProducts)
			})
		case "Offered By":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(8) > span > div > span").Each(func(i int, t *goquery.Selection) {
				log.Println(s.Text())
				response.OfferedBy = t.Text()
				log.Println(response.OfferedBy)
			})
		case "Developer":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(8) > span > div > span").Each(func(i int, t *goquery.Selection) {
				log.Println(s.Text())
				response.Developer = t.Text()
				log.Println(response.Developer)
			})
		}
	})




	//---------------------------------------------------------------------------------------
	// #3
	doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(9) > div").Each(func(i int, s *goquery.Selection){
		switch s.Text() {
		case "Interactive Elements":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(9) > span > div > span").Each(func(i int, t *goquery.Selection) {
				log.Println(s.Text())
				response.InteractiveElements = t.Text()
				log.Println(response.InteractiveElements)
			})
		case "In-app Products":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(9) > span > div > span").Each(func(i int, t *goquery.Selection) {
				log.Println(s.Text())
				response.InAppProducts = t.Text()
				log.Println(response.InAppProducts)
			})
		case "Offered By":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(9) > span > div > span").Each(func(i int, t *goquery.Selection) {
				log.Println(s.Text())
				response.OfferedBy = t.Text()
				log.Println(response.OfferedBy)
			})
		case "Developer":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(9) > span > div > span").Each(func(i int, t *goquery.Selection) {
				log.Println(s.Text())
				response.Developer = t.Text()
				log.Println(response.Developer)
			})
		}
	})

	//---------------------------------------------------------------------------------------
	// #4
	doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(10) > div").Each(func(i int, s *goquery.Selection){
		switch s.Text() {
		case "Interactive Elements":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(10) > span > div > span.htlgb").Each(func(i int, t *goquery.Selection) {
				log.Println(s.Text())
				response.InteractiveElements = t.Text()
				log.Println(response.InteractiveElements)
			})
		case "In-app Products":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(10) > span > div > span.htlgb").Each(func(i int, t *goquery.Selection) {
				log.Println(s.Text())
				response.InAppProducts = t.Text()
				log.Println(response.InAppProducts)
			})
		case "Offered By":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(10) > span > div > span.htlgb").Each(func(i int, t *goquery.Selection) {
				log.Println(s.Text())
				response.OfferedBy = t.Text()
				log.Println(response.OfferedBy)
			})
		case "Developer":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(10) > span > div > span.htlgb").Each(func(i int, t *goquery.Selection) {
				log.Println(s.Text())
				response.Developer = t.Text()
				log.Println(response.Developer)
			})
		}
	})

	//---------------------------------------------------------------------------------------
	// #5
	doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(11) > div").Each(func(i int, s *goquery.Selection){
		switch s.Text() {
		case "Interactive Elements":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(11) > span > div > span").Each(func(i int, t *goquery.Selection) {
				log.Println(s.Text())
				response.InteractiveElements = t.Text()
				log.Println(response.InteractiveElements)
			})
		case "In-app Products":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(11) > span > div > span").Each(func(i int, t *goquery.Selection) {
				log.Println(s.Text())
				response.InAppProducts = t.Text()
				log.Println(response.InAppProducts)
			})
		case "Offered By":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(11) > span > div > span").Each(func(i int, t *goquery.Selection) {
				log.Println(s.Text())
				response.OfferedBy = t.Text()
				log.Println(response.OfferedBy)
			})
		case "Developer":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(11) > span > div > span").Each(func(i int, t *goquery.Selection) {
				log.Println(s.Text())
				response.Developer = t.Text()
				log.Println(response.Developer)
			})
		}
	})

	//---------------------------------------------------------------------------------------
	// #6
	doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(12) > div").Each(func(i int, s *goquery.Selection){
		switch s.Text() {
		case "Interactive Elements":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(12) > span > div > span").Each(func(i int, t *goquery.Selection) {
				log.Println(s.Text())
				response.InteractiveElements = t.Text()
				log.Println(response.InteractiveElements)
			})
		case "In-app Products":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(12) > span > div > span").Each(func(i int, t *goquery.Selection) {
				log.Println(s.Text())
				response.InAppProducts = t.Text()
				log.Println(response.InAppProducts)
			})
		case "Offered By":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(12) > span > div > span").Each(func(i int, t *goquery.Selection) {
				log.Println(s.Text())
				response.OfferedBy = t.Text()
				log.Println(response.OfferedBy)
			})
		case "Developer":
			doc.Find("div.W4P4ne > div.JHTxhe.IQ1z0d > div > div:nth-child(12) > span > div > span").Each(func(i int, t *goquery.Selection) {
				log.Println(s.Text())
				response.Developer = t.Text()
				log.Println(response.Developer)
			})
		}
	})
	//-------------------------------------------------------------------




	//pass
	if err := k.webStore.Details(response) ; err!= nil{
		log.Fatal(err)
	}
}

 return err
}





//-------------------------------------------------------------------------------

	

	