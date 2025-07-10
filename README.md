# TUITodo

A beautiful terminal-based todo list application built with Go and Charm libraries.

## Features

- ✨ Interactive terminal UI with keyboard navigation
- ✅ Mark todos as complete/incomplete
- ➕ Add new todos with title and description
- ✏️ Edit existing todos (title and description)
- 📝 Multi-line description support
- 👁️ Detail view to see full todo information
- 🗑️ Delete todos
- 📄 Automatic text truncation for long titles
- 💾 Persistent storage using TODO.md format
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
- `e` - Edit selected todo
- `Enter` - View todo details
- `space` - Toggle todo completion
- `d` - Delete selected todo
- `q` - Quit application

**Input/Edit View:**
- `Enter` - Continue to description (from title) or save todo (from description)
- `Esc` - Cancel and return to list

**Detail View:**
- `e` - Edit this todo
- `Esc` or `q` - Return to list view

## Data Storage

Todos are automatically saved to `TODO.md` in the current directory using the standard TODO.md format. The file is created automatically on first use.

### TODO.md Format

The app uses the TODO.md format which is based on GitHub Flavored Markdown task lists:

```markdown
# TUITodo
A terminal-based todo list application

### Tasks
- [ ] Task title  
  - Task description  

### Completed ✓
- [x] Completed task title  
```

This format is portable, version-control friendly, and can be viewed/edited in any text editor or on GitHub.

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) - TUI components
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Style definitions

## License

MIT