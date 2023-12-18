package models

// define a custom data type for the scraped data
type Movie struct {
	Show
	Director Person
	Writers  []Person
	Stars    []Person
}
