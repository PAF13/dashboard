package tui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	list     list.Model
	choice   string
	quitting bool
}

func newItems() *[]list.Item {
	return &[]list.Item{
		item("Overview"),
		item("Bank"),
		item("Calendar"),
		item("TO-DO"),
		item("Notes"),
	}
}

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("| " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

/*
 */
type appUI struct {
	mainMenu list.Model
	choice   int
	quitting bool

	vim      vim
	bank     bank
	calendar calendar
	todo     todo
	notes    notes
}

type vim struct {
	mode string
}

type bank struct{}
type calendar struct{}
type todo struct{}
type notes struct{}
