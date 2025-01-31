package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	godotenv.Load()
}

func GetMongoURI() string {
	return os.Getenv("MONGO_URI")
}

func GetRedisURI() string {
	return os.Getenv("REDIS_URI")
}

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}


func GetPort() string {
	return os.Getenv("PORT")
}