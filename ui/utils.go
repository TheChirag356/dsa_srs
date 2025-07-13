package ui

import (
	"time"
	tea "github.com/charmbracelet/bubbletea"
)

type returnToMenuMsg struct{}

func waitAndReturn() tea.Cmd {
	return tea.Tick(3*time.Second, func(t time.Time) tea.Msg {
		return returnToMenuMsg{}
	})
}