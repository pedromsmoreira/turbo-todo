package models

type Todo struct {
	Id          string
	Attributes  *Attributes
	DateCreated string
	Version     int64
}

type Attributes struct {
	Title   string
	Content string
	Tags    []string
	Status  string
}
