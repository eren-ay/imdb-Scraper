package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type ReleaseDate struct {
	Year  int
	Month int
	Day   int
}

type Show struct {
	imgSrc string
	title  string
}

var url string = "imdb.com"

func main() {

	startDate := ReleaseDate{
		Year:  1950,
		Month: 01,
		Day:   01,
	}
	endDate := ReleaseDate{
		Year:  2042,
		Month: 12,
		Day:   31,
	}

	if searchByReleaseDate(startDate, endDate) {
		fmt.Println("Hello, World!")
	}

}

func searchByReleaseDate(startReleaseDate ReleaseDate, endReleaseDate ReleaseDate) bool {

	c := colly.NewCollector(
		colly.AllowedDomains("imdb.com", "www.imdb.com"),
	)

	//for the show details page
	//dataCollector := c.Clone()

	//each show one div and classname is = .mode-advanced
	c.OnHTML(".mode-advanced", func(e *colly.HTMLElement) {

		//showPageUrl := e.ChildAttr("div.lister-item-image > a", "href")
		//showPageUrl = movie, anime or tvshow id  example = <a href="/title/tt4589218/?ref_=adv_li_i">…</a>
		//showPageUrl = e.Request.AbsoluteURL(showPageUrl)

		// data collector visits showpage
		// no need to visit show detail page ,necessary information is available on search page
		//dataCollector.Visit(showPageUrl)

		//image source

		showImgSrc := e.ChildAttr("div.lister-item-image > a > img", "src")
		showTitle := e.ChildText("div.lister-item-content > h3.lister-item-header > a")
		tmpShow := Show{}
		tmpShow.title = showTitle
		tmpShow.imgSrc = showImgSrc
		fmt.Printf("%s\n", tmpShow.title)
	})
	// going next page
	/*c.OnHTML("a.lister-page-next", func(e *colly.HTMLElement) {
		nextPage := e.Request.AbsoluteURL(e.Attr("href"))
		c.Visit(nextPage)
	})*/

	startUrl := fmt.Sprintf("https://www.imdb.com/search/title/?release_date=%d-%d-%d,%d-%d-%d", startReleaseDate.Year, startReleaseDate.Month, startReleaseDate.Day, endReleaseDate.Year, endReleaseDate.Month, endReleaseDate.Day)
	c.Visit(startUrl)

	return true

}
