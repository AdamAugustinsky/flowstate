package main

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 2 }
func (d itemDelegate) Spacing() int                            { return 1 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Todo)
	if !ok {
		return
	}

	title := i.DisplayTitle()
	desc := i.Description
	if desc == "" {
		desc = "No description"
	} else if len(desc) > 60 {
		desc = desc[:57] + "..."
	}

	if index == m.Index() {
		title = selectedItemStyle.Render("> " + title)
		desc = selectedItemStyle.Render("  " + desc)
	} else {
		title = itemStyle.Render(title)
		desc = dimStyle.Render(itemStyle.Render(desc))
	}

	fmt.Fprintf(w, "%s\n%s", title, desc)
}