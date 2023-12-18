package models

type Episode struct {
	TVShow
	Director Person
	Writers  []Person
	Stars    []Person
}
