package component

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func StaticListComponent(title string, data []string) staticListModel {
	m := staticListModel{
		title: title,
		data:  data,
	}

	p := tea.NewProgram(m)

	model, err := p.Run()

	if err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}

	return model.(staticListModel)
}

type staticListModel struct {
	title string
	data  []string
}

func (m staticListModel) Init() tea.Cmd {
	return nil
}

func (m staticListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Quit
}

func (m staticListModel) View() string {
	s := "\n" + titleStyle(m.title) + "\n\n"
	for _, d := range m.data {
		s += "  " + textStyle(d) + "\n\n"
	}
	return s
}
