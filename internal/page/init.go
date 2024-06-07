package page

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kaziwaseef/stacker/internal/component"
)

func InitPage() {

	p := tea.NewProgram(component.InitialModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
