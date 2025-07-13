package ui

import "github.com/charmbracelet/lipgloss"

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

	errorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Bold(true)
	successStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("63")).
			Background(lipgloss.Color("235")).
			Padding(0, 1).
			MarginBottom(1).
			Bold(true)

	promptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("36"))

	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Italic(true)

	formPadding = lipgloss.NewStyle().MarginLeft(2)

	placeholderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Italic(true)
)