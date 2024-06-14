package main

import (
	"cleanArchitecture/server"
	"fmt"
	"log"
)

func main() {
	e, err := server.New()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(e.Start(fmt.Sprintf(":%d", 1322)))
}
