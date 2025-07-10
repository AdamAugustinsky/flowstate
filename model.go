package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type viewMode int

const (
	listView viewMode = iota
	inputView
	detailView
	editView
)

type model struct {
	list         list.Model
	textInput    textinput.Model
	textArea     textarea.Model
	viewport     viewport.Model
	store        *TodoStore
	mode         viewMode
	width        int
	height       int
	selectedTodo *Todo
	inputStep    int // 0=title, 1=description
}

func initialModel() model {
	store := NewTodoStore("TODO.md")
	
	items := []list.Item{}
	for _, todo := range store.GetAll() {
		items = append(items, todo)
	}

	const defaultWidth = 20
	const defaultHeight = 14

	l := list.New(items, itemDelegate{}, defaultWidth, defaultHeight)
	l.Title = "Flowstate"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	ti := textinput.New()
	ti.Placeholder = "Todo title"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 50

	ta := textarea.New()
	ta.Placeholder = "Todo description (optional)"
	ta.CharLimit = 500
	ta.SetWidth(50)
	ta.SetHeight(4)

	vp := viewport.New(50, 10)

	return model{
		list:      l,
		textInput: ti,
		textArea:  ta,
		viewport:  vp,
		store:     store,
		mode:      listView,
		inputStep: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.list.SetWidth(msg.Width)
		m.list.SetHeight(msg.Height - 4)
		m.viewport.Width = msg.Width - 4
		m.viewport.Height = msg.Height - 4
		return m, nil

	case tea.KeyMsg:
		switch m.mode {
		case listView:
			switch keypress := msg.String(); keypress {
			case "q", "ctrl+c":
				return m, tea.Quit
			case "a":
				m.mode = inputView
				m.inputStep = 0
				m.textInput.SetValue("")
				m.textArea.SetValue("")
				return m, m.textInput.Focus()
			case "enter":
				if i, ok := m.list.SelectedItem().(Todo); ok {
					m.selectedTodo = &i
					m.mode = detailView
					m.updateViewport()
				}
				return m, nil
			case " ":
				if i, ok := m.list.SelectedItem().(Todo); ok {
					m.store.Toggle(i.ID)
					m.refreshList()
				}
				return m, nil
			case "d":
				if i, ok := m.list.SelectedItem().(Todo); ok {
					m.store.Delete(i.ID)
					m.refreshList()
				}
				return m, nil
			case "e":
				if i, ok := m.list.SelectedItem().(Todo); ok {
					m.selectedTodo = &i
					m.mode = editView
					m.inputStep = 0
					m.textInput.SetValue(i.Title)
					m.textArea.SetValue(i.Description)
					return m, m.textInput.Focus()
				}
				return m, nil
			}
			
		case inputView:
			switch keypress := msg.String(); keypress {
			case "enter":
				if m.inputStep == 0 {
					m.inputStep = 1
					m.textArea.Focus()
					return m, textarea.Blink
				} else {
					if m.textInput.Value() != "" {
						m.store.Add(m.textInput.Value(), m.textArea.Value())
						m.refreshList()
					}
					m.mode = listView
					m.inputStep = 0
				}
				return m, nil
			case "esc":
				m.mode = listView
				m.inputStep = 0
				return m, nil
			}
		
		case detailView:
			switch keypress := msg.String(); keypress {
			case "q", "esc":
				m.mode = listView
				return m, nil
			case "e":
				m.mode = editView
				m.inputStep = 0
				m.textInput.SetValue(m.selectedTodo.Title)
				m.textArea.SetValue(m.selectedTodo.Description)
				return m, m.textInput.Focus()
			}
		
		case editView:
			switch keypress := msg.String(); keypress {
			case "enter":
				if m.inputStep == 0 {
					m.inputStep = 1
					m.textArea.Focus()
					return m, textarea.Blink
				} else {
					if m.selectedTodo != nil && m.textInput.Value() != "" {
						m.store.Update(m.selectedTodo.ID, m.textInput.Value(), m.textArea.Value())
						m.refreshList()
						// Update the selected todo with new values
						if todo := m.store.GetByID(m.selectedTodo.ID); todo != nil {
							m.selectedTodo = todo
						}
					}
					m.mode = listView
					m.inputStep = 0
				}
				return m, nil
			case "esc":
				m.mode = listView
				m.inputStep = 0
				return m, nil
			}
		}
	}

	var cmd tea.Cmd
	switch m.mode {
	case listView:
		m.list, cmd = m.list.Update(msg)
	case inputView:
		if m.inputStep == 0 {
			m.textInput, cmd = m.textInput.Update(msg)
		} else {
			m.textArea, cmd = m.textArea.Update(msg)
		}
	case detailView:
		m.viewport, cmd = m.viewport.Update(msg)
	case editView:
		if m.inputStep == 0 {
			m.textInput, cmd = m.textInput.Update(msg)
		} else {
			m.textArea, cmd = m.textArea.Update(msg)
		}
	}
	return m, cmd
}

func (m model) View() string {
	if m.width == 0 {
		return "loading..."
	}

	var content string
	switch m.mode {
	case listView:
		content = docStyle.Render(m.list.View())
	case inputView:
		var inputContent string
		if m.inputStep == 0 {
			inputContent = lipgloss.JoinVertical(lipgloss.Left,
				titleStyle.Render("Add New Todo - Title"),
				"",
				m.textInput.View(),
				"",
				helpStyle.Render("Press Enter to continue to description, Esc to cancel"),
			)
		} else {
			inputContent = lipgloss.JoinVertical(lipgloss.Left,
				titleStyle.Render("Add New Todo - Description"),
				"",
				m.textArea.View(),
				"",
				helpStyle.Render("Press Enter to save, Esc to cancel"),
			)
		}
		content = docStyle.Render(inputContent)
	case detailView:
		content = docStyle.Render(m.detailView())
	case editView:
		var editContent string
		if m.inputStep == 0 {
			editContent = lipgloss.JoinVertical(lipgloss.Left,
				titleStyle.Render("Edit Todo - Title"),
				"",
				m.textInput.View(),
				"",
				helpStyle.Render("Press Enter to continue to description, Esc to cancel"),
			)
		} else {
			editContent = lipgloss.JoinVertical(lipgloss.Left,
				titleStyle.Render("Edit Todo - Description"),
				"",
				m.textArea.View(),
				"",
				helpStyle.Render("Press Enter to save, Esc to cancel"),
			)
		}
		content = docStyle.Render(editContent)
	}

	help := m.helpView()
	height := m.height - lipgloss.Height(help) - 1
	
	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().Height(height).Render(content),
		help,
	)
}

func (m *model) refreshList() {
	items := []list.Item{}
	for _, todo := range m.store.GetAll() {
		items = append(items, todo)
	}
	m.list.SetItems(items)
}

func (m model) helpView() string {
	var help string
	switch m.mode {
	case listView:
		help = helpStyle.Render("a: add • e: edit • enter: details • space: toggle • d: delete • q: quit")
	case inputView:
		help = helpStyle.Render("enter: next/save • esc: cancel")
	case detailView:
		help = helpStyle.Render("e: edit • esc/q: back to list")
	case editView:
		help = helpStyle.Render("enter: next/save • esc: cancel")
	}
	return help
}

func (m *model) updateViewport() {
	if m.selectedTodo != nil {
		content := lipgloss.JoinVertical(lipgloss.Left,
			titleStyle.Render("Todo Details"),
			"",
			lipgloss.NewStyle().Bold(true).Render("Title:"),
			m.selectedTodo.Title,
			"",
			lipgloss.NewStyle().Bold(true).Render("Description:"),
			m.selectedTodo.Description,
			"",
			lipgloss.NewStyle().Bold(true).Render("Status:"),
			func() string {
				if m.selectedTodo.Completed {
					return "✓ Completed"
				}
				return "○ Pending"
			}(),
			"",
			lipgloss.NewStyle().Bold(true).Render("Created:"),
			m.selectedTodo.CreatedAt.Format("2006-01-02 15:04:05"),
		)
		m.viewport.SetContent(content)
	}
}

func (m model) detailView() string {
	return m.viewport.View()
}