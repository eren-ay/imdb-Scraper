package main

import (
	"fmt"
	"log"
	"scraper/imdb/models"
	"scraper/imdb/pkg/database"
	"time"

	"github.com/tebeka/selenium"
)

type Date struct {
	day   string
	month string
	year  string
}

// define a custom data type for the scraped data

func main() {

	// where to store the scraped data
	var Shows []models.Show

	// initialize a Chrome browser instance on port 4444
	service, err := selenium.NewChromeDriverService("./chromedriver", 4444)
	if err != nil {
		log.Fatal("Error:", err)
	}

	defer service.Stop()

	// configure the browser options
	caps := selenium.Capabilities{}

	// create a new remote client with the specified options
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// maximize the current window to avoid responsive rendering
	err = driver.MaximizeWindow("")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// visit the target page
	//https://www.imdb.com/search/title/?release_date=2023-01-01,2023-12-31&sort=release_date,asc
	//https://www.imdb.com/search/title/?title_type=tv_series,feature,tv_movie,tv_miniseries&release_date=1961-01-01,1961-12-31&sort=release_date,asc&num_votes=1,
	err = driver.Get("https://www.imdb.com/search/title/?title_type=tv_series,feature,tv_movie,tv_miniseries&release_date=1962-01-01,1964-12-31&sort=release_date,asc&num_votes=1,")
	if err != nil {
		log.Fatal("Error:", err)
	}
	//HoyXN

	showMoreBtn, err := driver.FindElement(selenium.ByCSSSelector, ".ipc-see-more__button")
	if err != nil {
		log.Fatal("Error:", err)
	}

	for i := 0; i < 117; i++ {
		showMoreBtn.Click()
		showMoreBtn.Click()
		time.Sleep(2 * time.Second)
	}
	showMoreBtn.Click()
	showMoreBtn.Click()
	time.Sleep(2 * time.Second)
	// select the product elements
	showElements, err := driver.FindElements(selenium.ByCSSSelector, ".ipc-metadata-list-summary-item__c")

	if err != nil {
		log.Fatal("Error:", err)
	}

	// iterate over the product elements
	// and extract data from them
	for _, showElement := range showElements {
		nameElement, err := showElement.FindElement(selenium.ByCSSSelector, ".ipc-title-link-wrapper")

		link, err := nameElement.GetAttribute("href")
		fmt.Printf("%s \n\n\n", link)
		fmt.Printf("%s \n", parseLinkForId(link))

		name, err := nameElement.Text()
		if err != nil {
			log.Fatal("Error:", err)
		}

		// add the scraped data to the list
		show := models.Show{}
		show.Title = name
		show.ID = parseLinkForId(link)
		Shows = append(Shows, show)
	}
	database.InsertCollection(database.DB, "Show", "Movie", Shows)
	fmt.Println(Shows)
}

func parseLinkForId(link string) string {
	id := ""
	for i := 1; i < len(link); i++ {
		if link[i] == '/' {
			for j := i + 1; j < len(link); j++ {
				if link[j] == '/' {
					return id
				}
				id = id + string(link[j])
			}
		}
	}
	return ""
}
