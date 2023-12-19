package showdetails

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func scraperCreate() *colly.Collector {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})
	return c

}

func PeopleInfoScraper(peopleUrl string) {
	c := scraperCreate()

	c.OnHTML(".gelzOk", func(e *colly.HTMLElement) {
		name := e.ChildText(".dxeMrU")
		jobs := e.ChildTexts(".cdJsTz > .ipc-inline-list__item")
		birthDate := e.ChildTexts(".dzVFxE")
		fmt.Println(name)
		fmt.Println(jobs[1])
		fmt.Printf("%s ", birthDate[1])
	})

	c.Visit(peopleUrl)
}
