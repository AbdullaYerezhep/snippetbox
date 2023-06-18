package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string `json:"port"`
	DSN  string `json:"dsn"`
	StaticDir string `json:"static_dir"`
	SecretSession string `json:"secret_session"`
}

var (
	instance *Config
	once     sync.Once
)

// GetConfig returns a singleton instance of Config struct
func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		instance.configParser()
	})
	return instance
}

func (c *Config) configParser() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	host := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("PORT")
	
	c.DSN = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, dbPort, user, password, dbname)
	c.Port = port
}