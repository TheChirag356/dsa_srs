package ui

import (
	"fmt"
	"strings"
)

func (m model) View() string {
	var b strings.Builder

	b.WriteString("\nWelcome!\n")
	b.WriteString("Use ↑/↓ or j/k to move, enter to select, q to quit:\n\n")

	for i, choice := range m.options {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = "👉"
		}
		fmt.Fprintf(&b, "%s %s\n", cursor, choice)
	}
	return b.String()
}
