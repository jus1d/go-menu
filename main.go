package main

import "go-menu/menu"

func main() {
	mn, _ := menu.New("This is a 1 option", "This is a 2nd option", "Tis is a third option")

	mn.Choose()

	mn.GetChosenAnswer()
}
