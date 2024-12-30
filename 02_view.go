package main

import (
	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	if m.quitting {
		return quitTextStyle.Render("Not hungry? That’s cool.")
	}

	lists := lipgloss.JoinHorizontal(lipgloss.Top, m.list.View(), m.choice)
	return lists
}
