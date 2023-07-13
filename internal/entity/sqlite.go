package entity

import (
	"encoding/json"
)

type User struct {
	ID       int64   `db:"id" gorm:"primaryKey;AUTO_INCREMENT;not null;"`
	Login    string  `db:"login"`
	Password string  `db:"password"`
	Token    *string `db:"token"`
	Option   []byte  `db:"option"`
}

type ExteranlSQL interface {
	Get(conditional string, values []any) (*User, error)
	GetList() ([]*User, error)
	Create(*User) error
	Update(*User) error
	Delete(*User) error
}

func SetOption(option map[string]any) ([]byte, error) {

	if option == nil {
		return nil, nil
	}

	b, err := json.Marshal(option)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GetOption(b []byte) (map[string]any, error) {

	if len(b) == 0 ||
		string(b) == "{}" ||
		string(b) == "null" {
		return nil, nil
	}

	option := make(map[string]any)

	err := json.Unmarshal(b, &option)
	if err != nil {
		return nil, err
	}
	return option, nil
}
