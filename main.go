package main

import (
	"fmt"
	"go-menu/menu"
)

func main() {
	mn, _ := menu.New("First option here", "Second option performed here", "It's third option")

	mn.Choose()

	chosen, option := mn.GetChosenAnswer()

	fmt.Printf("Choosed option number: %d, it's text: %d", chosen, option)
}
