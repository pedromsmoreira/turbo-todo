package todo

type todo struct {
	Id          string
	Attributes  *attributes
	DateCreated string
	Version     int64
}

type attributes struct {
	Title   string
	Content string
	Tags    []string
	Status  string
}
