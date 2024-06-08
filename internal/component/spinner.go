package component

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func SpinnerComponent[T interface{}](ctx context.Context, fn func(ctx context.Context) *T) spinnermodel[T] {
	m := spinnermodel[T]{
		ctx:         ctx,
		fn:          fn,
		spinnerType: spinner.Dot,
	}
	m.resetSpinner()

	p := tea.NewProgram(m)
	model, err := p.Run()

	if err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}

	return model.(spinnermodel[T])
}

type spinnermodel[T interface{}] struct {
	ctx         context.Context
	fn          func(context.Context) *T
	Data        *T
	spinnerType spinner.Spinner
	spinner     spinner.Model
	quitting    bool
}

func (m spinnermodel[T]) Init() tea.Cmd {
	return tea.Batch(func() tea.Msg {
		value := m.fn(m.ctx)
		return value
	}, m.spinner.Tick)
}

func (m spinnermodel[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case *T:
		m.Data = msg
		m.quitting = true
		return m, tea.Quit
	default:
		return m, nil
	}
}

func (m *spinnermodel[T]) resetSpinner() {
	m.spinner = spinner.New()
	m.spinner.Style = spinnerStyle
	m.spinner.Spinner = m.spinnerType
}

func (m spinnermodel[T]) View() string {
	if m.quitting {
		return ""
	}
	s := fmt.Sprintf("\n %s%s%s\n\n", m.spinner.View(), " ", textStyle("Loading..."))
	return s
}
