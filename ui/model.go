package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type menuOption int


const (
	menuViewDueCards menuOption = iota
	menuAddConceptCard
	menuAddProblemCard
	menuExit
)

type model struct {
	cursor  int
	options []string
}

func (m model) Init() tea.Cmd {
	// _, err := storage.LoadCards("data/cards.json")
	// if err != nil {
	// 	log.Fatal("Error loading cards:", err)
	// }
	return nil // No initial command needed
}

func initialModel() model {
	return model{
		cursor:   0,
		options: []string{
			"📋 Review Due Cards",
			"🧠 Add Concept Card",
			"🧩 Add Problem Card",
			"❌ Exit",
		},
	}
}
