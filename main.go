package main

import (
	"github.com/semecky/sandbox/ducktyping"
)

type Human ducktyping.Human

func main() {
	var h = Human{}

	h.EatVegetable()
	h.EatMeat()

}
