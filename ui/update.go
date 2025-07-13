package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor -= 1
			}

		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor += 1
			}

		case "enter":
			switch m.cursor {
			case int(menuViewDueCards):
				return newReviewModel(), nil
			case int(menuAddConceptCard):
				return newAddConceptModel(), nil
			case int(menuAddProblemCard):
				return newAddProblemModel(), nil
			case int(menuExit):
				return m, tea.Quit
			}
		}
	}
	return m, nil
}