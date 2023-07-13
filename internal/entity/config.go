package entity

import "time"

type Config struct {
	FirstUser struct {
		Login    string `yaml:"login"`
		Password string `yaml:"password" `
	} `yaml:"first_user"`

	HTTP struct {
		Mode    string `yaml:"mode"`
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
		Timeout struct {
			Write time.Duration `yaml:"write" default:"1m"`
			Read  time.Duration `yaml:"read" default:"1m"`
		} `yaml:"timeout"`
	} `yaml:"server_http"`

	JWT struct {
		TTL         time.Duration `yaml:"ttl"`
		TokenClaims string        `yaml:"token_claims"`
	}

	SQL struct {
		Dir  string `yaml:"dir"`
		Name string `yaml:"name"`
	} `yaml:"sql`
}
