package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	fmt.Println("Go Sherdog Wrapper")
	f := FindFighterByID(FighterID("Robert-Whittaker-45132"))
	spew.Dump(f)
}
