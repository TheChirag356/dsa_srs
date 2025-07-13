package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

type returnToMenuMsg struct{}

func waitAndReturn() tea.Cmd {
	return tea.Tick(3*time.Second, func(t time.Time) tea.Msg {
		return returnToMenuMsg{}
	})
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
