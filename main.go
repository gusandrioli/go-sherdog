package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	fmt.Println("Go Sherdog Wrapper")
	f := FindFighterByID(FighterID("Michael-Whittaker-75667"))
	FindFighterByID(FighterID("Robert-Whittaker-45132"))
	FindFighterByID(FighterID("Robert-Whittaker-372576"))
	FindFighterByID(FighterID("Sean-Whittaker-168981"))

	spew.Dump(f)
}
