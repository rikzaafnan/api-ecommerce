package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	SERVERPORT = ":9797"
)

type LoadEnv struct {
	USERNAMEDB   string
	PASSWORDDB   string
	SCHEMADB     string
	HOSTDB       string
	PORTDB       string
	ENVIRONTMENT string
	JWTTOKEN     string
}

func LoadENV() LoadEnv {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	usernameDB := os.Getenv("USERNAME_DATABASE")
	passwordDB := os.Getenv("PASSWORD_DATABASE")
	schemaDB := os.Getenv("SCHEMA_DATABASE")
	hostDB := os.Getenv("HOST_DATABASE")
	portDB := os.Getenv("PORT_DATABASE")
	environtment := os.Getenv("ENVIRONTMENT")
	jwtToken := os.Getenv("JWT_SECRET_KEY")

	loadENv := LoadEnv{
		USERNAMEDB:   usernameDB,
		PASSWORDDB:   passwordDB,
		SCHEMADB:     schemaDB,
		HOSTDB:       hostDB,
		PORTDB:       portDB,
		ENVIRONTMENT: environtment,
		JWTTOKEN:     jwtToken,
	}

	return loadENv

}
