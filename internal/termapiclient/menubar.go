package termapiclient

import "github.com/mum4k/termdash/cell"

type MenubarButton struct {
	Label   string
	Color   cell.Color
	Handler func() error
}
