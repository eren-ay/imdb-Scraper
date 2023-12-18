package models

// define a custom data type for the scraped data
type Show struct {
	ID    string `bson:"ID"`
	Title string `bson:"Title"`
	ShowDetails
}
