package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		log.Println("Error opening .env file: ", err)
		return err
	}

	return nil
}
