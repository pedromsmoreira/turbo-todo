package todos

type todoData struct {
	Id          string
	DateCreated string
	Title       string
	Content     string
	Tags        []string
	Status      string
	Version     int64
}
