# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Flowstate is a terminal-based todo list application built with Go and the Charm library ecosystem. It provides an elegant TUI for managing tasks using the TODO.md format.

## Commands

### Build and Run
```bash
# Build the application
go build -o flowstate

# Run the application
./flowstate       # Local mode (uses ./TODO.md)
./flowstate -g    # Global mode (uses ~/config/flowstate/TODO.md)

# Install globally
go install

# Run tests (when implemented)
go test ./...

# Update dependencies
go mod tidy
```

## Architecture

### Core Components

1. **main.go** - Entry point that initializes the Bubble Tea program
2. **model.go** - Contains the main application state and Bubble Tea MVC logic:
   - Manages view modes: listView, inputView, detailView, editView
   - Handles keyboard navigation and user input
   - Coordinates between UI and data storage
3. **todo.go** - Defines the Todo struct and TodoStore with CRUD operations
4. **todomd.go** - Handles parsing and writing TODO.md format:
   - Converts between Todo structs and markdown format
   - Only creates TODO.md when first todo is added
   - Removes TODO.md when all todos are deleted
5. **delegate.go** - List item rendering logic for the Bubbles list component
6. **styles.go** - Lip Gloss style definitions for UI consistency

### Key Design Decisions

- **TODO.md Format**: Uses standard markdown task lists instead of JSON for portability
- **No Default Todos**: Starts with empty state, respects user's directory
- **Multi-step Input**: Title first, then optional description for better UX
- **View Modes**: Separate modes for list, input, detail, and edit operations

### Data Flow

1. User input → model.Update() → TodoStore operations → todomd.go writes to TODO.md
2. On startup: todomd.go reads TODO.md → TodoStore → model displays in list

### Important Patterns

- All file operations go through TodoStore methods (Add, Update, Delete, Toggle)
- The TODO.md file uses specific regex patterns for parsing (see todomd.go)
- Empty directories are kept clean - no file creation until needed
- Text truncation happens in todo.go DisplayTitle() and delegate.go for long content

## Dependencies

- Bubble Tea: Main TUI framework
- Bubbles: Pre-built components (list, textinput, textarea, viewport)
- Lip Gloss: Styling and layout

## Module Path

The module is `github.com/adamaugustinsky/flowstate` - ensure go.mod matches when making changes.