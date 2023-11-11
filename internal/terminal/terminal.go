package terminal

import (
	"context"
	"time"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
)

type Terminal struct {
	term      *tcell.Terminal
	container *container.Container
	Config    *TerminalConfig
}

type TerminalConfig struct {
	RedrawInterval time.Duration
	RootID         string
}

type InterfaceBuilder interface {
	BuildGrid(builder *grid.Builder) error
}

// redrawInterval is how often termdash redraws the screen.
const defaultRedrawInterval = 250 * time.Millisecond

// rootID is the ID assigned to the root container.
const defaultRootID = "root"

func NewTerminal(cfg *TerminalConfig) *Terminal {
	t := &Terminal{
		Config: &TerminalConfig{
			RedrawInterval: defaultRedrawInterval,
			RootID:         defaultRootID,
		},
	}

	if cfg != nil {
		t.Config = cfg
	}

	return t
}

func (t *Terminal) buildInterface(uiBuilder InterfaceBuilder) error {

	builder := grid.New()
	if err := uiBuilder.BuildGrid(builder); err != nil {
		return err
	}

	gridOpts, err := builder.Build()
	if err != nil {
		return err
	}

	if len(gridOpts) > 0 {
		t.container.Update(t.Config.RootID, gridOpts...)
	}

	return nil
}

func (t *Terminal) Start(uiBuilder InterfaceBuilder) error {

	term, err := tcell.New(tcell.ColorMode(terminalapi.ColorMode256))
	if err != nil {
		return err
	}
	defer term.Close()

	t.term = term

	container, err := container.New(t.term, container.ID(t.Config.RootID))
	if err != nil {
		return err
	}
	t.container = container

	// build UI
	if err := t.buildInterface(uiBuilder); err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())

	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == keyboard.KeyEsc || k.Key == keyboard.KeyCtrlC {
			cancel()
		}
	}

	err = termdash.Run(
		ctx,
		t.term,
		t.container,
		termdash.KeyboardSubscriber(quitter),
		termdash.RedrawInterval(t.Config.RedrawInterval),
	)
	if err != nil {
		return err
	}

	return nil
}
