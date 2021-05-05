package main

import "fmt"

type PostgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (c PostgresConfig) Dialect() string {
	return "postgres"
}

func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "111",
		Name:     "mywebapp_dev",
	}
}

func (c PostgresConfig) ConnectionInfo() string {
	if c.Password == "" {
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Name)
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Name)
}

type Config struct {
	Port int
	Env  string
}

func DefaultConfig() Config {
	return Config{
		Port: 3000,
		Env:  "dev",
	}
}

func (c Config) IsProd() bool {
	return c.Env == "prod"
}

// isProd := false

// Port = 3000

// fmt.Println("Starting server on :3000...")
// http.ListenAndServe(":3000", csrfMw(userMw.Apply(r)))

// # models/users.go

// const userPwPepper = "secret-random-string"
// const hmacSecretKey = "secret-hmac-key"

// # models/services.go

// db, err := gorm.Open("postgres", connectionInfo)
// 	if err != nil {
// 		return nil, err
// 	}
// 	db.LogMode(true)
