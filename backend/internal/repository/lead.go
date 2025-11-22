package repository

type Lead struct {
	Id           int
	Organization string
	Stack        []string
	Role         string
	Contact      string
	ConsultantId int
}
