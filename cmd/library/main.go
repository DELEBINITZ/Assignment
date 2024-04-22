package main

import (
	"log"

	"github.com/akshay/libraryAssign/internal/library"
)

func main() {
	app := library.InitializeApp()
	if err := app.Run(":80"); err != nil {
		log.Fatal(err)
	}
}
