package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func envACCOUNTSID() string {
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env files %v", err)
	}
	return os.Getenv("TWILIO_ACCOUNT_SID")
}
