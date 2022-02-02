package main

import (
	"context"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

type WebService interface {
	GeneratePackage() error
	Scrap() error
}

type webService struct {
	webStore 		WebStore
	b				chan string
	playstoreUrl 	string
	headerNodes 	[]*cdp.Node
	headNodes		[]*cdp.Node

}

func NewWebService(webStore WebStore) WebService {
	a := webService{webStore: webStore, b: make(chan string)}
	var wg sync.WaitGroup

	wg.Add(4)

	go func() {
		defer wg.Done()
		a.GeneratePackage()
	}()
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			a.Scrap()
		}()
	}
	wg.Wait()
	return &a
}

func (t *webService) GeneratePackage() error {
	t.playstoreUrl = "https://play.google.com"

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(3)
	go func(ctx context.Context) {
		defer wg.Done()
		t.categoriesList(ctx)
	}(ctx)
	go func(ctx context.Context) {
		defer wg.Done()
		t.urlList(ctx)
	}(ctx)
	go func(ctx context.Context) {
		defer wg.Done()
		t.searchKeyword(ctx)
	}(ctx)
	
	wg.Wait()
	close(t.b)
	return nil
}

func (t *webService) categoriesList(ctx context.Context){
	clone, cancel := chromedp.NewContext(ctx)
    defer cancel()
	categories := []string{	"https://play.google.com/store/apps/category/AUTO_AND_VEHICLES",
							"https://play.google.com/store/apps/category/BEAUTY",
							"https://play.google.com/store/apps/category/EVENTS",
							"https://play.google.com/store/apps/category/HOUSE_AND_HOME",
						}
	for j := 0; j < len(categories); j++ {
		err := chromedp.Run(clone,
						chromedp.Navigate(categories[j]),
						chromedp.ActionFunc(func(c2 context.Context) error {
							var v int
							a := make(map[int]int)

							for i := 0; i < 10; i++ {

								count := len(a)
								_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(c2)
								time.Sleep(2 * time.Second)
								if err != nil {
									return err
								}
								if exp != nil {
									return exp
								}

								err = chromedp.Run(c2,
									chromedp.Evaluate(`document.querySelector('.ZmHEEd').childNodes.length;`, &v),
								)
								a[v] = i
								if len(a) == count {
									break
								}

								if err != nil {
									return err
								}
							}
							return nil
						}),
						chromedp.Nodes(":is(div.wXUyZd > a)", &t.headNodes, chromedp.ByQueryAll),
						chromedp.ActionFunc(func(c3 context.Context) error {
							for _, node := range t.headNodes {
								//log.Println("categories list url==", j)
								u := node.AttributeValue("href")
								appUrl := t.playstoreUrl + u
								t.b <- appUrl
							}
							return nil
						}),
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	
}


func (t *webService) urlList(ctx context.Context){
	clone2, cancel := chromedp.NewContext(ctx)
    defer cancel()
	categories := []string{	"https://play.google.com/store/apps",
							"https://play.google.com/store/apps/top",
							"https://play.google.com/store/apps/new",
							"https://play.google.com/store/apps/category/ART_AND_DESIGN",
							"https://play.google.com/store/apps/stream/baselist_featured_arcore",
							"https://play.google.com/store/apps/category/BOOKS_AND_REFERENCE",
							"https://play.google.com/store/apps/category/BUSINESS",
							"https://play.google.com/store/apps/category/COMICS",
							"https://play.google.com/store/apps/category/COMMUNICATION",
							"https://play.google.com/store/apps/category/DATING",
							"https://play.google.com/store/apps/stream/vr_top_device_featured_category",
							"https://play.google.com/store/apps/category/EDUCATION",
							"https://play.google.com/store/apps/category/ENTERTAINMENT",
							"https://play.google.com/store/apps/category/FINANCE",
							"https://play.google.com/store/apps/category/FOOD_AND_DRINK",
							"https://play.google.com/store/apps/category/HEALTH_AND_FITNESS",
							"https://play.google.com/store/apps/category/LIBRARIES_AND_DEMO",
							"https://play.google.com/store/apps/category/LIFESTYLE",
							"https://play.google.com/store/apps/category/MAPS_AND_NAVIGATION",
							"https://play.google.com/store/apps/category/MEDICAL",
							"https://play.google.com/store/apps/category/MUSIC_AND_AUDIO",
							"https://play.google.com/store/apps/category/NEWS_AND_MAGAZINES",
							"https://play.google.com/store/apps/category/PERSONALIZATION",
							"https://play.google.com/store/apps/category/PHOTOGRAPHY",
							"https://play.google.com/store/apps/category/PRODUCTIVITY",
							"https://play.google.com/store/apps/category/SHOPPING",
							"https://play.google.com/store/apps/category/SOCIAL",
							"https://play.google.com/store/apps/category/TOOLS",
							"https://play.google.com/store/apps/category/TRAVEL_AND_LOCAL",
							"https://play.google.com/store/apps/category/VIDEO_PLAYERS",
							"https://play.google.com/store/apps/category/ANDROID_WEAR",
							"https://play.google.com/store/apps/category/WATCH_FACE",
							"https://play.google.com/store/apps/category/WEATHER",
							"https://play.google.com/store/apps/category/GAME",
							"https://play.google.com/store/apps/category/GAME_ACTION",
							"https://play.google.com/store/apps/category/GAME_ADVENTURE",
							"https://play.google.com/store/apps/category/GAME_ARCADE",
							"https://play.google.com/store/apps/category/GAME_BOARD",
							"https://play.google.com/store/apps/category/GAME_CARD",
							"https://play.google.com/store/apps/category/GAME_CASINO",
							"https://play.google.com/store/apps/category/GAME_CASUAL",
							"https://play.google.com/store/apps/category/GAME_EDUCATIONAL",
							"https://play.google.com/store/apps/category/GAME_MUSIC",
							"https://play.google.com/store/apps/category/GAME_PUZZLE",
							"https://play.google.com/store/apps/category/GAME_RACING",
							"https://play.google.com/store/apps/category/GAME_ROLE_PLAYING",
							"https://play.google.com/store/apps/category/GAME_SIMULATION",
							"https://play.google.com/store/apps/category/GAME_SPORTS",
							"https://play.google.com/store/apps/category/GAME_STRATEGY",
							"https://play.google.com/store/apps/category/GAME_TRIVIA",
							"https://play.google.com/store/apps/category/GAME_WORD",
							"https://play.google.com/store/apps/category/FAMILY",
							"https://play.google.com/store/apps/category/FAMILY?age=AGE_RANGE1",
							"https://play.google.com/store/apps/category/FAMILY?age=AGE_RANGE2",
							"https://play.google.com/store/apps/category/FAMILY?age=AGE_RANGE3",
						}
	q := "https://play.google.com/store/apps/editors_choice"

	for j := 0; j < len(categories); j++ {
		//log.Println("url list==", j)
		err := chromedp.Run(clone2,
			chromedp.Navigate(categories[j]),
			chromedp.ActionFunc(func(c1 context.Context) error {
				var v int
				a := make(map[int]int)

				for i := 0; i < 10; i++ {

					count := len(a)
					_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(c1)
					time.Sleep(2 * time.Second)
					if err != nil {
						return err
					}
					if exp != nil {
						return exp
					}

					err = chromedp.Run(c1,
						chromedp.Evaluate(`document.querySelector('.ZmHEEd').childNodes.length;`, &v),
					)
					a[v] = i
					if len(a) == count {
						break
					}

					if err != nil {
						return err
					}

				}
				return nil
			}),

			chromedp.Nodes(":is(div.xwY9Zc > a)", &t.headerNodes, chromedp.ByQueryAll),
			chromedp.ActionFunc(func(c2 context.Context) error {
				for _, node := range t.headerNodes {
					cluster := node.AttributeValue("href")
					url := t.playstoreUrl + cluster
					// log.Println("cluster list url list==",n)

					if url != q {
						err := chromedp.Run(c2,
							chromedp.Navigate(url),
							chromedp.ActionFunc(func(c3 context.Context) error {
								var v int
								a := make(map[int]int)

								for i := 0; i < 10; i++ {

									count := len(a)
									_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(c3)
									time.Sleep(2 * time.Second)
									if err != nil {
										return err
									}
									if exp != nil {
										return exp
									}

									err = chromedp.Run(c3,
										chromedp.Evaluate(`document.querySelector('.ZmHEEd').childNodes.length;`, &v),
									)
									a[v] = i
									if len(a) == count {
										break
									}

									if err != nil {
										return err
									}

								}
								return nil
							}),
							chromedp.Nodes(":is(div.wXUyZd > a)", &t.headNodes, chromedp.ByQueryAll),
							chromedp.ActionFunc(func(c4 context.Context) error {
								for _, node := range t.headNodes {
									u := node.AttributeValue("href")
									appUrl := t.playstoreUrl + u
									// log.Println("appurl url list==",q)
									t.b <- appUrl
								}
								return nil
							}),
						)
						if err != nil {
							log.Fatal(err)
							return err
						}
						continue
					}
				}
				return nil
			}),
		)
		if err != nil {
			log.Fatal(err)
		}
	}
 }


 func (t *webService) searchKeyword(ctx context.Context){
	clone3, cancel := chromedp.NewContext(ctx)
    defer cancel()
	search := "https://play.google.com/store/search?q="
	apps := "&c=apps"
	keyword := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for i := 0; i < len(keyword); i++ {
		o := search + keyword[i] + apps
		log.Println("search key==",i)
		err := chromedp.Run(clone3,
			chromedp.Navigate(o),
			chromedp.ActionFunc(func(c1 context.Context) error {
				var v int
				a := make(map[int]int)

				for i := 0; i < 10; i++ {

					count := len(a)
					_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(c1)
					time.Sleep(2 * time.Second)
					if err != nil {
						return err
					}
					if exp != nil {
						return exp
					}

					err = chromedp.Run(c1,
						chromedp.Evaluate(`document.querySelector('.ZmHEEd').childNodes.length;`, &v),
					)
					a[v] = i
					if len(a) == count {
						break
					}

					if err != nil {
						return err
					}

				}
				return nil
			}),
			chromedp.Nodes(":is(div.wXUyZd > a)", &t.headNodes, chromedp.ByQueryAll),
			chromedp.ActionFunc(func(c2 context.Context) error {
				for r, node := range t.headNodes {
					u := node.AttributeValue("href")
					appUrl := t.playstoreUrl + u
					log.Println("appurl search list==",r)
					t.b <- appUrl

				}
				return nil
			}),
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	 
 }



func (t *webService) Scrap() (err error) {
	
	for url := range t.b {

		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Println(err)
		}

		response := Model{}

		packageName := url
		packageName = strings.TrimPrefix(packageName, "https://play.google.com/store/apps/details?id=")
		response.PackageName = packageName

		response.Url = url

		doc.Find(".AHFaub").Each(func(i int, v *goquery.Selection) {
			response.AppName = v.Text()
			//log.Println(response.AppName)
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
		if err := t.webStore.Create(response); err != nil {
			log.Fatal(err)
		}
	}

	return err
}

//-------------------------------------------------------------------------------
