package main

import (
	"flag"
	"log"
	"os"

	"github.com/kyf/klein/session"
)

var (
	confpath string
)

func main() {
	flag.StringVar(&confpath, "config", "", "session server config path")
	flag.Parse()

	if confpath == "" {
		flag.Usage()
		os.Exit(0)
	}

	svr, err := session.NewServer(confpath)
	if err != nil {
		log.Fatal(err)
	}
	svr.Run()
}
