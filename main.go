package main

import "github.com/tj/anydns/server"
import "github.com/tj/anydns/config"
import "github.com/tj/go-gracefully"
import "github.com/tj/docopt"
import "log"
import "os"

var Version = "0.0.1"

const Usage = `
  Usage:
    anydns
    anydns -h | --help
    anydns --version

  Options:
    -h, --help       output help information
    -v, --version    output version

`

func main() {
	_, err := docopt.Parse(Usage, nil, true, Version, false)
	if err != nil {
		log.Fatalf("[error] %s", err)
	}

	c, err := config.Read(os.Stdin)
	if err != nil {
		log.Fatalf("[error] parsing config: %s", err)
	}

	s := server.New(c)

	log.Printf("[info] starting anydns %s", Version)
	err = s.Start()
	if err != nil {
		log.Fatalf("[error] %s", err)
	}

	log.Printf("[info] binding on %s", s.Config.Bind)
	gracefully.Shutdown()

	log.Printf("[info] stopping")
	err = s.Stop()
	if err != nil {
		log.Fatalf("[error] %s", err)
	}

	log.Printf("[info] bye :)")
}
