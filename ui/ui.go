package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	// "github.com/charmbracelet/lipgloss"
	"log"
)

func StartTUI() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		log.Fatal("Error running TUI:", err)
	}
}
