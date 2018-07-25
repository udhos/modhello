package main

import (
	"log"

	"github.com/udhos/modhello/modlib"
)

func main() {
	run(1, 2)
}

func run(a, b int) {
	log.Printf("Sum(%d,%d) = %d", a, b, modlib.Sum(a, b))
	log.Printf("Sub(%d,%d) = %d", a, b, modlib.Sub(a, b))
}
