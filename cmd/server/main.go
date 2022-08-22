package main

import (
	"log"

	"github.com/txthinking/socks5"
)

func main() {
	s, err := socks5.NewClassicServer(":1080", "127.0.0.1", "", "", 0, 60)
	if err != nil {
		log.Fatal(err)
	}
	// You can pass in custom Handler
	s.ListenAndServe(nil)
}
