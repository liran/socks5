package main

import (
	"fmt"
	"log"
	"os"

	"github.com/liran/socks5"
)

func main() {
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	port := os.Getenv("PORT")
	if port == "" {
		port = "1080"
	}
	s, err := socks5.NewClassicServer(fmt.Sprintf(":%s", port), username, password, 0, 60)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("server listen at %s", port)
	s.ListenAndServe(nil)
}
