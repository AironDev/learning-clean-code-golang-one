package config

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Secrets struct {
	RefreshTokenDuration int    `json:"REFRESH_TOKEN_DURATION"`
	AuthTokenDuration    int    `json:"AUTH_TOKEN_DURATION"`
	JWTPublicKey         string `json:"JWT_PUBLIC_KEY"`
	JWTPrivateKey        string `json:"JWT_PRIVATE_KEY"`

	DatabasePort     string `json:"DATABASE_PORT"`
	DatabaseHost     string `json:"DATABASE_HOST"`
	DatabaseUser     string `json:"DATABASE_USER"`
	DatabasePassword string `json:"DATABASE_PASSWORD"`
	DatabaseName     string `json:"DATABASE_NAME"`

	AppPort string `json:"APP_PORT"`
	AppUrl  string `json:"APP_URL"`
}

var ss Secrets

func Init() *Secrets {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file: \n %v", err)
		}
	}

	ss = Secrets{}
	ss.RefreshTokenDuration, _ = getenvInt("REFRESH_TOKEN_DURATION")
	ss.AuthTokenDuration, _ = getenvInt("AUTH_TOKEN_DURATION")
	ss.JWTPublicKey = os.Getenv("JWT_PUBLIC_KEY")
	ss.JWTPrivateKey = os.Getenv("JWT_PRIVATE_KEY")

	ss.DatabasePort = os.Getenv("DATABASE_PORT")
	ss.DatabaseHost = os.Getenv("DATABASE_HOST")
	ss.DatabaseUser = os.Getenv("DATABASE_USER")
	ss.DatabasePassword = os.Getenv("DATABASE_PASSWORD")
	ss.DatabaseName = os.Getenv("DATABASE_NAME")

	ss.AppPort = os.Getenv("APP_PORT")
	ss.AppUrl = os.Getenv("APP_URL")

	return &ss

}

// GetSecrets is used to get value from the Secrets runtime.
func GetSecrets() *Secrets {
	return &ss
}

var ErrEnvVarEmpty = errors.New("getenv: environment variable empty")

func getenvStr(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return v, ErrEnvVarEmpty
	}
	return v, nil
}

func getenvInt(key string) (int, error) {
	s, err := getenvStr(key)
	if err != nil {
		return 0, err
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return v, nil
}
