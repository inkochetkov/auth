package entity

type User struct {
	Login    string
	Password string
	Token    string
	Option   map[string]any
}
