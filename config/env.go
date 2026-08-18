package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBDSN      string
	SWGoH      SWGoHConfig
	AuthConfig AuthConfig
	Mode       string
}

type SWGoHConfig struct {
	FetchCharacterURL string
	FetchPlayerURL    string
}

type AuthConfig struct {
	Salt      string
	JWTSecret string
	JWTExpiry int
}

var (
	Cfg     Config
	envOnce sync.Once
)

func LoadEnv() Config {
	envOnce.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found, using system default envs")
		}

		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			os.Getenv("BACKEND_DB_HOST"),
			os.Getenv("BACKEND_DB_USER"),
			os.Getenv("BACKEND_DB_PASSWORD"),
			os.Getenv("BACKEND_DB_NAME"),
			os.Getenv("BACKEND_DB_PORT"),
			os.Getenv("BACKEND_DB_SSLMODE"),
		)
		jwtExpiry, err := strconv.Atoi(os.Getenv("AUTH_JWT_EXPIRY"))
		if err != nil {
			log.Println("Invalid AUTH_JWT_EXPIRY, using 24")
			jwtExpiry = 24
		}

		Cfg = Config{
			Port:  os.Getenv("BACKEND_PORT"),
			DBDSN: dsn,
			SWGoH: SWGoHConfig{
				FetchCharacterURL: os.Getenv("SWGOH_FETCH_CHARACTER_URL"),
				FetchPlayerURL:    os.Getenv("SWGOH_FETCH_PLAYER_URL"),
			},
			AuthConfig: AuthConfig{
				Salt:      os.Getenv("AUTH_SALT"),
				JWTSecret: os.Getenv("AUTH_JWT_SECRET"),
				JWTExpiry: jwtExpiry,
			},
			Mode: os.Getenv("MODE"),
		}
	})

	return Cfg
}
