package menu

import (
	"errors"
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

type Menu struct {
	options     []string
	chosenPoint int
}

var ErrNoOptions = errors.New("at least 1 option should be provided")

var cyan = color.New(color.FgCyan).SprintFunc()

func New(options ...string) (*Menu, error) {
	if len(options) == 0 {
		return nil, ErrNoOptions
	}

	return &Menu{
		options: options,
	}, nil
}

func (m *Menu) Choose() {
	selected := 0
	err := keyboard.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err = keyboard.Close()
		if err != nil {
			panic(err)
		}
	}()

	m.print(selected)

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println(err)
			return
		}

		switch key {
		case keyboard.KeyArrowUp:
			if selected > 0 {
				selected--
			}
		case keyboard.KeyArrowDown:
			if selected < len(m.options)-1 {
				selected++
			}
		case keyboard.KeyEnter:
			m.chosenPoint = selected
			return
		}

		fmt.Print("\033[H\033[2J")
		m.print(selected)
	}
}

func (m *Menu) GetChosenAnswer() (int, string) {
	return m.chosenPoint, m.options[m.chosenPoint]
}

func (m *Menu) print(chosenPoint int) {
	fmt.Println("Choose an option:")

	for i, option := range m.options {
		if i == chosenPoint {
			fmt.Printf(cyan(fmt.Sprintf("%d. %s <<\n", i+1, option)))
		} else {
			fmt.Printf("%d. %s\n", i+1, option)
		}
	}
}
