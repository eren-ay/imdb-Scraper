package models

type ShowDetails struct {
	CastMembers []Person
	Awards      string
	Poster      []byte //image file mongogridfs
}
