package entity

type User struct {
	ID       int
	Login    string
	Password string
	Token    string
	Option   map[string]any
}
