package main

import (
	"design_patterns/behavioral/strategy/shapes"
	"log"
	"os"
)

func main() {
	var output = "text"

	activeStrategy, err := shapes.NewPrinter(output)
	if err != nil {
		log.Fatal(err)
	}

	switch output {
	case "text":
		activeStrategy.SetWriter(os.Stdout)
		err := activeStrategy.Print()
		if err != nil {
			log.Fatal(err)
		}
	case "image":
		w, err := os.Create("./image.jpg")
		if err != nil {
			log.Fatal(err)
		}
		defer w.Close()
		activeStrategy.SetWriter(w)

		err = activeStrategy.Print()
		if err != nil {
			log.Fatal(err)
		}
	}
}
