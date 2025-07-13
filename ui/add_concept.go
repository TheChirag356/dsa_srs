package ui

import (
	"fmt"
	"strings"

	"github.com/TheChirag356/dsa_anki/constants"
	"github.com/TheChirag356/dsa_anki/storage"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type addConceptModel struct {
	inputs     []textinput.Model
	focusIndex int
	submitted  bool
	saved      bool
	errMsg     string
}

func newAddConceptModel() addConceptModel {
	tiTitle := textinput.New()
	tiTitle.Placeholder = "Enter concept title"
	tiTitle.Focus()
	tiTitle.CharLimit = 64
	tiTitle.Prompt = "Title: "

	tiNotes := textinput.New()
	tiNotes.Placeholder = "Enter your notes here"
	tiNotes.CharLimit = 256
	tiNotes.Prompt = "Notes: "
	tiNotes.PromptStyle.Faint(true)

	return addConceptModel{
		inputs:     []textinput.Model{tiTitle, tiNotes},
		focusIndex: 0,
		submitted:  false,
	}
}

func (m addConceptModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m addConceptModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case returnToMenuMsg:
		return initialModel(), nil
	case tea.KeyMsg:
		if m.submitted {
			return m, tea.Quit
		}
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			if s == "enter" && m.focusIndex == len(m.inputs)-1 && !m.submitted {
				m.submitted = true

				// Extract input values
				title := m.inputs[0].Value()
				notes := m.inputs[1].Value()

				// Validation
				if title == "" || notes == "" {
					m.errMsg = "Title and Notes cannot be empty."
					return m, waitAndReturn()
				}

				// Save to storage
				store, err := storage.LoadCards(constants.CardFilePath)
				if err != nil {
					m.errMsg = fmt.Sprintf("Error loading cards: %v", err)
					return m, waitAndReturn()
				}

				card, err := storage.NewConceptCard(title, notes)
				if err != nil {
					m.errMsg = fmt.Sprintf("Error creating concept card: %v", err)
					return m, waitAndReturn()
				}

				store.AppendConceptCard(card)

				if err := storage.SaveCards(constants.CardFilePath, store); err != nil {
					m.errMsg = fmt.Sprintf("Error saving card: %v", err)
				} else {
					m.saved = true
				}

				return m, waitAndReturn()
			} else if s == "enter" && m.focusIndex < len(m.inputs)-1 && !m.submitted {
				m.focusIndex++
			}

			// Navigate inputs
			if s == "up" || s == "shift+tab" {
				if m.focusIndex > 0 {
					m.focusIndex--
				}
			} else if s == "down" || s == "tab" {
				if m.focusIndex < len(m.inputs)-1 {
					m.focusIndex++
				}
			}

			for i := 0; i < len(m.inputs); i++ {
				if i == m.focusIndex {
					m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
				} else {
					m.inputs[i].Blur()
					m.inputs[i].PromptStyle = blurredStyle
					m.inputs[i].TextStyle = blurredStyle
				}
			}

		default:
			// Update inputs
			cmds := make([]tea.Cmd, len(m.inputs))
			for i := range m.inputs {
				m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
			}
			return m, tea.Batch(cmds...)
		}
	}

	// Update inputs
	cmds := make([]tea.Cmd, len(m.inputs))
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m addConceptModel) View() string {
	if m.submitted {
		if m.errMsg != "" {
			return errorStyle.Render("Error: "+m.errMsg) + "\n\nReturning to main menu in 3 seconds."
		}
		return successStyle.Render("Concept card saved!") + "\n\nReturning to main menu in 3 seconds."
	}

	var b strings.Builder
	b.WriteString(titleStyle.Render("Add New Concept Card") + "\n\n")

	for _, input := range m.inputs {
		b.WriteString(formPadding.Render(input.View()) + "\n")
	}

	b.WriteString("\n" + helpStyle.Render("[Tab] Next • [Enter] Submit • [Ctrl+C] Quit"))
	return b.String()
}
