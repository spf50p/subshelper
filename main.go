package main

import (
	"log"

	"github.com/spf50p/subshelper/cmd"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	cmd.Execute()
}
