package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func envACCOUNTSID() string {
	println(godotenv.Unmarshal(".env"))
	// Load the .env file in the current directory
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env files %v", err)
	}
	//use os.Getenv to retrieve an environment variable
	return os.Getenv("TWILIO_ACCOUNT_SID")
}
