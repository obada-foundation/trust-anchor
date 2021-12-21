package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, " :: TRUST-ANCHOR  :: ", 0)

	if err := run(logger); err != nil {

	}
}

func run(logger *log.Logger) error {
	return nil
}
