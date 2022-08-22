package main

import (
	"flag"
	"log"
	"os"

	"github.com/liran/socks5"
)

func main() {
	var publicIP string
	flag.StringVar(&publicIP, "host", "127.0.0.1", "")
	flag.Parse()
	if publicIP == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	s, err := socks5.NewClassicServer(":1080", publicIP, "", "", 0, 60)
	if err != nil {
		log.Fatal(err)
	}
	// You can pass in custom Handler
	s.ListenAndServe(nil)
}
