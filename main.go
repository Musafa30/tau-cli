package main

import (
	"log"
	"os"

	"github.com/taubyte/tau/cli"
	"github.com/taubyte/tau/i18n"
)

func main() {
	err := cli.Run(os.Args...)
	if err != nil {
		log.Fatal(i18n.AppCrashed(err))
	}
}
