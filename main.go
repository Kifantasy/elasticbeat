package main

import (
	"github.com/brandslisten/elasticbeat/beater"

	"github.com/elastic/beats/libbeat/beat"
	"os"
)

var Name = "elasticbeat"

func main() {
	if err := beat.Run(Name, "", beater.New()); err != nil {
		os.Exit(1)
	}
}
