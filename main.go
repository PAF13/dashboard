package main

import (
	"fmt"
	"log"
	"os"

	neondb "github.com/PAF13/com_neondb"
	"github.com/PAF13/dashboard/ui/httpserver"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

func main() {
	neondb.GetTrans()

	if false {
		tempFunc()
	}
}

func tempFunc() {
	// Start logging server
	logBuf := &httpserver.LogBuffer{}
	log.SetOutput(logBuf) // Redirect log output to the buffer.
	go httpserver.Server(logBuf)

	items := newItems()

	const defaultWidth = 20

	l := list.New(*items, itemDelegate{}, 100, 21)
	l.Title = ""
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := model{list: l}

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
