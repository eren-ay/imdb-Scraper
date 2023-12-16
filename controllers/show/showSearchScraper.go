package show

import (
	"scraper/imdb/models"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

func ShowScraperByReleaseDate(driver selenium.WebDriver, Shows []models.Show) ([]models.Show, error) {
	// visit the target page
	//https://www.imdb.com/search/title/?release_date=2023-01-01,2023-12-31&sort=release_date,asc
	//https://www.imdb.com/search/title/?title_type=tv_series,feature,tv_movie,tv_miniseries&release_date=1961-01-01,1961-12-31&sort=release_date,asc&num_votes=1,
	err := driver.Get("https://www.imdb.com/search/title/?title_type=tv_series,feature,tv_movie,tv_miniseries&release_date=2010-01-01,2010-12-31&sort=release_date,asc&num_votes=1,")
	if err != nil {
		return nil, err
	}
	time.Sleep(50 * time.Second)
	showMoreBtn, err := driver.FindElement(selenium.ByCSSSelector, ".ipc-see-more__button")
	if err != nil {
		return nil, err
	}

	for i := 0; i < 234; i++ {
		showMoreBtn.Click()
		showMoreBtn.Click()
		time.Sleep(3 * time.Second)
	}
	showMoreBtn.Click()
	showMoreBtn.Click()
	time.Sleep(5 * time.Second)
	// select the product elements
	showElements, err := driver.FindElements(selenium.ByCSSSelector, ".ipc-metadata-list-summary-item__c")

	if err != nil {
		return nil, err
	}

	// iterate over the product elements
	// and extract data from them
	for _, showElement := range showElements {
		nameElement, err := showElement.FindElement(selenium.ByCSSSelector, ".ipc-title-link-wrapper")

		link, err := nameElement.GetAttribute("href")
		//fmt.Printf("%s \n\n\n", link)
		//fmt.Printf("%s \n", parseLinkForId(link))

		name, err := nameElement.Text()
		if err != nil {
			return nil, err
		}

		// add the scraped data to the list
		show := models.Show{}
		// copy the name
		parseIndex := strings.Index(name, ".")
		// example name ="211. Hancock"
		show.Title = name[parseIndex+2:]
		show.ID = parseLinkForId(link)
		Shows = append(Shows, show)
	}
	return Shows, nil
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
