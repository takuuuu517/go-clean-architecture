package main

import (
	"cleanArchitecture/server"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	e, err := server.New()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))))
}
