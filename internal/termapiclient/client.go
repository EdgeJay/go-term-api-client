package termapiclient

import (
	"fmt"

	"github.com/EdgeJay/go-term-api-client/internal/terminal"

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

func (client *TermApiClient) BuildGrid(builder *grid.Builder) error {

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

	return nil
}

func (client *TermApiClient) Start() error {
	return client.Terminal.Start(client)
}
