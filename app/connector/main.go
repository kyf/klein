package main

import (
	"github.com/kyf/klein/connector"

	"flag"
	"log"
	"os"
)

var (
	confpath string
)

func main() {
	flag.StringVar(&conpath, "config", "", "connector config path")
	flag.Parse()

	if confpath == "" {
		flag.Usage()
		os.Exit(0)
	}

	svr := connector.NewConnector()
	err := svr.Init(confpath)
	if err != nil {
		log.Fatal(err)
	}
	svr.Run()
}
