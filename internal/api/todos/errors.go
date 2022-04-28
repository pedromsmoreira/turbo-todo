package todos

type DbResourceNotFound struct {
	Message string
}

func NewDbResourceNotFound(msg string) *DbResourceNotFound {
	return &DbResourceNotFound{
		Message: msg,
	}
}

func (rnf *DbResourceNotFound) Error() string {
	return rnf.Message
}
