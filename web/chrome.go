package web

import (
	"context"
	"log"
	"time"

	//"bitbucket.org/ajmal_azm/scraperP/"bitbucket.org/ajmal_azm/scraperP/api"api"

	//"bitbucket.org/ajmal_azm/scraperP/model"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

type Web interface {
	CategoriesList()
	UrlList()
	Searcher()
	PackageNameChan() chan string
}

type Webs struct {
	playstoreUrl string
	appURLChan   chan string
	headerNodes  []*cdp.Node
	headNodes    []*cdp.Node
}

func NewWeb() Web {
	a := Webs{appURLChan: make(chan string, 1000)}
	return &a
}

func (t *Webs) PackageNameChan() chan string {
	return t.appURLChan
}

func (t *Webs) CategoriesList() {

	t.playstoreUrl = "https://play.google.com"

	categories := []string{"https://play.google.com/store/apps/category/AUTO_AND_VEHICLES",
		"https://play.google.com/store/apps/category/BEAUTY",
		"https://play.google.com/store/apps/category/EVENTS",
		"https://play.google.com/store/apps/category/HOUSE_AND_HOME",
	}

	for j := 0; j < len(categories); j++ {
		clone1, cancel := chromedp.NewContext(context.Background())
		err := chromedp.Run(clone1,
			chromedp.Navigate(categories[j]),
			chromedp.ActionFunc(func(c2 context.Context) error {
				var v int
				a := make(map[int]int)

				for i := 0; i < 20; i++ {
					count := len(a)
					_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(c2)
					time.Sleep(2 * time.Second)
					if err != nil {
						log.Println(err)
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
						log.Println(err)
						return err
					}
				}
				return nil
			}),
			chromedp.Nodes(":is(div.wXUyZd > a)", &t.headNodes, chromedp.ByQueryAll),
			chromedp.ActionFunc(func(c3 context.Context) error {
				for _, node := range t.headNodes {
					u := node.AttributeValue("href")
					appUrl := t.playstoreUrl + u
					//log.Println("[home]", appUrl)
					t.appURLChan <- appUrl
				}

				return nil
			}),
		)
		if err != nil {
			log.Println(err)
		}
		cancel()
	}

}

func (t *Webs) UrlList() {
	t.playstoreUrl = "https://play.google.com"

	categories := []string{"https://play.google.com/store/apps",
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
		clone2, cancel := chromedp.NewContext(context.Background())
		err := chromedp.Run(clone2,
			chromedp.Navigate(categories[j]),
			chromedp.ActionFunc(func(c1 context.Context) error {
				var v int
				a := make(map[int]int)

				for i := 0; i < 20; i++ {

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

								for i := 0; i < 20; i++ {

									count := len(a)
									_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(c3)
									time.Sleep(2 * time.Second)
									if err != nil {
										log.Println(err)
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
										log.Println(err)
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
									//log.Println("[category]", appUrl)
									t.appURLChan <- appUrl
								}
								return nil
							}),
						)
						if err != nil {
							log.Println(err)
							return err
						}
						continue
					}
				}
				return nil
			}),
		)
		if err != nil {
			log.Println(err)
		}
		cancel()
	}

}

func next(a string) string {
	n := len(a) - 1
	out := ""
	add := 1
	for add > 0 || n >= 0 {
		if n < 0 {
			out = "a" + out
			add = 0
		} else {
			out = string(97+(int(a[n])-int('a')+add)%26) + out
			if add == 1 && a[n] != 'z' {
				add = 0
			}
		}
		n--
	}
	return out
}

func (t Webs) Searcher() {
	first := "a"

	for i := 0; i < 10000000; i++ {
		t.SearchKeyword(first)
		val := next(first)
		first = val
	}
}

func (t *Webs) SearchKeyword(searchItem string) {
	t.playstoreUrl = "https://play.google.com"

	search := "https://play.google.com/store/search?q="
	apps := "&c=apps"
	//first := "a"

	// for i:=0;i<10000;i++{
	// 	vale:=next(first)
	// 	t.c<-vale
	// 	first = vale
	// }

	// for i := 0; i<1000; i++ {
	// 	t.c <- next(first)
	// 	val :=<- t.c
	//val := next(first)
	clone3, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	o := search + searchItem + apps
	//first = val
	err := chromedp.Run(clone3,
		chromedp.Navigate(o),
		chromedp.ActionFunc(func(c1 context.Context) error {
			var v int
			a := make(map[int]int)

			for i := 0; i < 20; i++ {

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
			for _, node := range t.headNodes {
				u := node.AttributeValue("href")
				appUrl := t.playstoreUrl + u
				//log.Println("[search]", appUrl)
				t.appURLChan <- appUrl
			}

			return nil
		}),
	)
	if err != nil {
		log.Println(err)
	}

}
