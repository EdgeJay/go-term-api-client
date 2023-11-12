package termapiclient

import (
	"fmt"

	"github.com/EdgeJay/go-term-api-client/internal/terminal"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgets/button"
)

type TermApiClient struct {
	Terminal *terminal.Terminal
}

func NewTermApiClient(cfg *terminal.TerminalConfig) *TermApiClient {
	t := terminal.NewTerminal(cfg)
	return &TermApiClient{
		Terminal: t,
	}
}

func (client *TermApiClient) addMenuBar(builder *grid.Builder) []error {

	errs := make([]error, 0)

	buttons := make([]MenubarButton, 0)

	for idx := 1; idx <= 3; idx++ {
		buttons = append(buttons, MenubarButton{
			Label: fmt.Sprintf("Menu %d", idx),
			Color: cell.ColorBlue,
			Handler: func() error {
				return nil
			},
		})
	}

	colPerc := 100 / len(buttons)

	items := make([]grid.Element, len(buttons))

	for idx := 0; idx < len(buttons); idx++ {

		if b, err := button.New(
			buttons[idx].Label,
			buttons[idx].Handler,
			button.FillColor(buttons[idx].Color),
			button.Height(1),
		); err == nil {

			items[idx] = grid.ColWidthPerc(
				colPerc,
				grid.Widget(b),
			)
		} else {
			errs = append(errs, err)
		}
	}

	builder.Add(
		grid.RowHeightPerc(
			10,
			grid.RowHeightPerc(
				99,
				items...,
			),
		),
	)

	return errs
}

func (client *TermApiClient) BuildGrid(builder *grid.Builder) error {

	client.addMenuBar(builder)

	builder.Add(
		grid.RowHeightPerc(
			90,
			grid.Widget(nil,
				container.Border(linestyle.Light),
			),
		),
	)

	/*
		b, err := button.New("hello world", func() error {
			return nil
		})
		if err != nil {
			return fmt.Errorf("button.New => %v", err)
		}

		b2, err := button.New("hello world 2", func() error {
			return nil
		})
		if err != nil {
			return fmt.Errorf("button.New => %v", err)
		}

		builder.Add(
			grid.RowHeightPerc(50,
				grid.Widget(b,
					container.Border(linestyle.Light),
					container.BorderTitle("Button"),
					container.BorderTitleAlignCenter(),
				),
			),
			grid.RowHeightPerc(50,
				grid.Widget(b2,
					container.Border(linestyle.Double),
					container.BorderTitle("Button 2"),
					container.BorderTitleAlignCenter(),
				),
			),
		)
	*/

	return nil
}

func (client *TermApiClient) Start() error {
	return client.Terminal.Start(client)
}
