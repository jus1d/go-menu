# GO: Menu
**go-menu** - tool to simplify terminal menu management in Go

## Installation:
```bash
go get github.com/jus1d/go-menu
```

## Usage:

```go
package main

import (
	"fmt"
	"github.com/jus1d/go-menu/menu"
)

func main() {
	mn, _ := menu.New("First option here", "Second option performed here", "It's third option")

	mn.Choose()

	chosen, option := mn.GetChosenAnswer()

  fmt.Printf("Choosed option number: %d, it's text: %s", chosen, option)
}
```