package entity

import "time"

type Config struct {
	FirstUser struct {
		Login    string `yaml:"login"`
		Password string `yaml:"password" `
	} `yaml:"first_user"`

	HTTP Server `yaml:"server_http"`

	JWT struct {
		TTL         time.Duration `yaml:"ttl"`
		TokenClaims string        `yaml:"token_claims"`
	}

	SQL struct {
		PathSQL      string        `yaml:"path_sql"`
		PathSQLName  string        `yaml:"path_sql_name"`
		Timeout      time.Duration `yaml:"timeout"`
		MaxOpenConns int           `yaml:"max_open_conns"`
		MaxIdleConns int           `yaml:"max_idle_conns"`
	}
}

type Server struct {
	Mode    string  `yaml:"mode"`
	Host    string  `yaml:"host"`
	Port    string  `yaml:"port"`
	Timeout Timeout `yaml:"timeout"`
}

type Timeout struct {
	Write time.Duration `yaml:"write" default:"1m"`
	Read  time.Duration `yaml:"read" default:"1m"`
}
