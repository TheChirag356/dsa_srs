package main

import (
	"os"
	"log"

	"github.com/TheChirag356/dsa_srs/ui"
)

func init() {
	f, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)
}

func main() {
	ui.StartTUI()
}
