package main

import (
	"log"

	"github.com/udhos/modhello/modlib/lib"
)

func main() {
	run(1, 2)
}

func run(a, b int) {
	log.Printf("Sum(%d,%d) = %d", a, b, lib.Sum(a, b))
}
