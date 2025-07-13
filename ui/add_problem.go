package ui

import (
	"fmt"
	"strings"

	"github.com/TheChirag356/dsa_anki/constants"
	"github.com/TheChirag356/dsa_anki/storage"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type addProblemModel struct {
	inputs     []textinput.Model
	focusIndex int
	submitted  bool
	saved      bool
	errMsg     string
}

func newAddProblemModel() addProblemModel {
	tiTitle := textinput.New()
	tiTitle.Placeholder = "Enter problem title"
	tiTitle.CharLimit = 64
	tiTitle.Prompt = "Title: "
	tiTitle.Focus()
	tiTitle.PromptStyle = focusedStyle
	tiTitle.TextStyle = focusedStyle

	tiLink := textinput.New()
	tiLink.Placeholder = "Leetcode problem URL"
	tiLink.CharLimit = 128
	tiLink.Prompt = "Link: "
	tiLink.PromptStyle = blurredStyle
	tiLink.TextStyle = blurredStyle

	tiTopic := textinput.New()
	tiTopic.Placeholder = "e.g. binary search"
	tiTopic.CharLimit = 64
	tiTopic.Prompt = "Topic: "
	tiTopic.PromptStyle = blurredStyle
	tiTopic.TextStyle = blurredStyle

	return addProblemModel{
		inputs:     []textinput.Model{tiTitle, tiLink, tiTopic},
		focusIndex: 0,
		submitted:  false,
	}
}

func (m addProblemModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m addProblemModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case returnToMenuMsg:
		return initialModel(), nil

	case tea.KeyMsg:
		if m.submitted {
			return m, waitAndReturn()
		}

		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			if s == "enter" && m.focusIndex == len(m.inputs)-1 && !m.submitted {
				m.submitted = true

				title := m.inputs[0].Value()
				link := m.inputs[1].Value()
				topic := m.inputs[2].Value()

				if title == "" || link == "" || topic == "" {
					m.errMsg = "All fields are required."
					return m, waitAndReturn()
				}

				store, err := storage.LoadCards(constants.CardFilePath)
				if err != nil {
					m.errMsg = fmt.Sprintf(" loading cards: %v", err)
					return m, waitAndReturn()
				}

				card, err := storage.NewProblemCard(title, link, topic)
				if err != nil {
					m.errMsg = fmt.Sprintf("Error creating problem card: %v", err)
					return m, waitAndReturn()
				}

				store.ProblemCards = append(store.ProblemCards, card)

				if err := storage.SaveCards(constants.CardFilePath, store); err != nil {
					m.errMsg = fmt.Sprintf("Error saving card: %v", err)
				} else {
					m.saved = true
				}

				return m, waitAndReturn()
			} else if s == "enter" && m.focusIndex < len(m.inputs)-1 && !m.submitted {
				m.focusIndex++
			}

			// Focus management
			if s == "up" || s == "shift+tab" {
				if m.focusIndex > 0 {
					m.focusIndex--
				}
			} else if s == "down" || s == "tab" {
				if m.focusIndex < len(m.inputs)-1 {
					m.focusIndex++
				}
			}

			for i := range m.inputs {
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

func (m addProblemModel) View() string {
	if m.submitted {
		if m.errMsg != "" {
			return errorStyle.Render("Error: "+m.errMsg) + "\n\nReturning to main menu in 3 seconds."
		}
		return successStyle.Render("Problem card saved!") + "\n\nReturning to main menu in 3 seconds."
	}

	var b strings.Builder
	b.WriteString(titleStyle.Render("Add New Problem Card") + "\n\n")

	for _, input := range m.inputs {
		b.WriteString(formPadding.Render(input.View()) + "\n")
	}

	b.WriteString("\n" + helpStyle.Render("[Tab] Next • [Enter] Submit • [Ctrl+C] Quit"))
	return b.String()
}
