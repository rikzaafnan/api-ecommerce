package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	SERVERPORT         = ":9797"
	ENVIRONTMENT_LOCAL = "local"
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
	environtment := os.Getenv("ENVIRONTMENT_ENV")
	jwtToken := os.Getenv("JWT_SECRET_KEY")

	if strings.ToLower(environtment) == strings.ToLower(ENVIRONTMENT_LOCAL) {
		usernameDB = os.Getenv("USERNAME_DATABASE_DB_LOCAL")
		passwordDB = os.Getenv("PASSWORD_DATABASE_DB_LOCAL")
		schemaDB = os.Getenv("SCHEMA_DATABASE_DB_LOCAL")
		hostDB = os.Getenv("HOST_DATABASE_DB_LOCAL")
		portDB = os.Getenv("PORT_DATABASE_DB_LOCAL")
	}

	fmt.Println(environtment)

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
