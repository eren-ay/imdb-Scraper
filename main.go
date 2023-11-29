package main

import (
	"fmt"
	"log"
	"scraper/imdb/models"
	"time"

	"github.com/tebeka/selenium"
)

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
	err = driver.Get("https://www.imdb.com/search/title/?release_date=1960-01-01,2042-12-31&sort=release_date,asc/")
	if err != nil {
		log.Fatal("Error:", err)
	}

	showMoreBtn, err := driver.FindElement(selenium.ByCSSSelector, ".ipc-see-more__button")
	if err != nil {
		log.Fatal("Error:", err)
	}

	showMoreBtn.Click()
	showMoreBtn.Click()
	time.Sleep(10 * time.Second)
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
