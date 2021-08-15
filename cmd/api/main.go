package main

import (
	"github.com/Frankcs96/learning-platform-service/cmd/api/bootstrap"
	"log"
)

func main() {

	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
