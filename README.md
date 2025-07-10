# TUITodo

A beautiful terminal-based todo list application built with Go and Charm libraries.

## Features

- ✨ Interactive terminal UI with keyboard navigation
- ✅ Mark todos as complete/incomplete
- ➕ Add new todos with title and description
- 📝 Multi-line description support
- 👁️ Detail view to see full todo information
- 🗑️ Delete todos
- 📄 Automatic text truncation for long titles
- 💾 Persistent storage (saves to `todos.json`)
- 🎨 Beautiful styling with Lip Gloss

## Installation

```bash
git clone <repository>
cd tuitodo
go build
```

## Usage

Run the application:

```bash
./tuitodo
```

### Keyboard Shortcuts

**List View:**
- `↑/↓` or `j/k` - Navigate through todos
- `a` - Add a new todo
- `Enter` - View todo details
- `space` - Toggle todo completion
- `d` - Delete selected todo
- `q` - Quit application

**Input View:**
- `Enter` - Continue to description (from title) or save todo (from description)
- `Esc` - Cancel and return to list

**Detail View:**
- `Esc` or `q` - Return to list view

## Data Storage

Todos are automatically saved to `todos.json` in the current directory. The file is created automatically on first use.

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) - TUI components
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Style definitions

## License

MIT