package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func envAccountSID() string {
	fmt.Println(godotenv.Unmarshal(".env"))

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
		log.Fatal("error loading .env file")
	}

	return os.Getenv("TWILIO_ACCOUNT_SID")
}

func envServiceSID() string {
	fmt.Println(godotenv.Unmarshal(".env"))

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
		log.Fatal("error loading .env file")
	}

	return os.Getenv("TWILIO_SERVICE_SID")
}

func envAuthToken() string {
	fmt.Println(godotenv.Unmarshal(".env"))

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
		log.Fatal("error loading .env file")
	}

	return os.Getenv("TWILIO_AUTH_TOKEN")
}
