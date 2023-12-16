package main

import (
	"log"
	"scraper/imdb/controllers/show"
	"scraper/imdb/models"
	"scraper/imdb/pkg/database"

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
	getAllShowFromScrape(driver, Shows)

}

func getAllShowFromScrape(driver selenium.WebDriver, Shows []models.Show) []models.Show {
	Shows, err := show.ShowScraperByReleaseDate(driver, Shows)
	if err != nil {
		log.Fatal("Error:", err)
	}

	database.InsertCollection(database.DB, "Show", "Movie", Shows)
	return Shows
	//fmt.Println(Shows)
}
