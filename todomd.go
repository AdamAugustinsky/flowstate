package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	todoMDHeader = `# TUITodo
A terminal-based todo list application

`
	pendingSection = "### Tasks"
	completedSection = "### Completed ✓"
)

var (
	taskRegex = regexp.MustCompile(`^-\s+\[([x ])\]\s+(.+?)(?:\s{2,})?$`)
	simpleTaskRegex = regexp.MustCompile(`^-\s+(.+?)(?:\s{2,})?$`)
	sectionRegex = regexp.MustCompile(`^###\s+(.+)$`)
	subTaskRegex = regexp.MustCompile(`^  -\s+(.+?)(?:\s{2,})?$`)
)

type TodoMDStore struct {
	todos    []Todo
	filename string
}

func NewTodoMDStore(filename string) *TodoStore {
	store := &TodoStore{
		todos:    []Todo{},
		filename: filename,
	}
	store.Load()
	return store
}

func (s *TodoStore) parseTodoMD() error {
	file, err := os.Open(s.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	s.todos = []Todo{}
	scanner := bufio.NewScanner(file)
	
	var currentSection string
	var currentTodo *Todo
	var todoID int

	for scanner.Scan() {
		line := scanner.Text()
		
		// Check for section headers
		if matches := sectionRegex.FindStringSubmatch(line); matches != nil {
			currentSection = matches[1]
			currentTodo = nil
			continue
		}
		
		// Check for tasks with checkboxes
		if matches := taskRegex.FindStringSubmatch(line); matches != nil {
			completed := matches[1] == "x"
			title := strings.TrimSpace(matches[2])
			
			// Check if this section indicates completed tasks
			if strings.Contains(currentSection, "✓") || strings.Contains(currentSection, "[x]") {
				completed = true
			}
			
			todo := Todo{
				ID:          fmt.Sprintf("%d", todoID),
				Title:       title,
				Description: "",
				Completed:   completed,
				CreatedAt:   time.Now(), // We don't store creation time in TODO.md
			}
			
			if completed {
				now := time.Now()
				todo.CompletedAt = &now
			}
			
			s.todos = append(s.todos, todo)
			currentTodo = &s.todos[len(s.todos)-1]
			todoID++
			continue
		}
		
		// Check for simple tasks without checkboxes
		if matches := simpleTaskRegex.FindStringSubmatch(line); matches != nil {
			title := strings.TrimSpace(matches[1])
			
			// Skip empty lines that match the regex
			if title == "" {
				continue
			}
			
			completed := strings.Contains(currentSection, "✓") || strings.Contains(currentSection, "[x]")
			
			todo := Todo{
				ID:          fmt.Sprintf("%d", todoID),
				Title:       title,
				Description: "",
				Completed:   completed,
				CreatedAt:   time.Now(),
			}
			
			if completed {
				now := time.Now()
				todo.CompletedAt = &now
			}
			
			s.todos = append(s.todos, todo)
			currentTodo = &s.todos[len(s.todos)-1]
			todoID++
			continue
		}
		
		// Check for sub-tasks (descriptions)
		if currentTodo != nil && strings.HasPrefix(line, "  ") {
			trimmed := strings.TrimSpace(line)
			if strings.HasPrefix(trimmed, "- ") {
				trimmed = strings.TrimPrefix(trimmed, "- ")
			}
			if currentTodo.Description != "" {
				currentTodo.Description += "\n"
			}
			currentTodo.Description += trimmed
		}
	}
	
	return scanner.Err()
}

func (s *TodoStore) writeTodoMD() error {
	file, err := os.Create(s.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	
	// Write header
	writer.WriteString(todoMDHeader)
	
	// Separate todos into pending and completed
	var pending, completed []Todo
	for _, todo := range s.todos {
		if todo.Completed {
			completed = append(completed, todo)
		} else {
			pending = append(pending, todo)
		}
	}
	
	// Write pending tasks
	if len(pending) > 0 {
		writer.WriteString(pendingSection + "\n")
		for _, todo := range pending {
			writer.WriteString(fmt.Sprintf("- [ ] %s  \n", todo.Title))
			if todo.Description != "" {
				lines := strings.Split(todo.Description, "\n")
				for _, line := range lines {
					writer.WriteString(fmt.Sprintf("  - %s  \n", line))
				}
			}
		}
		writer.WriteString("\n")
	}
	
	// Write completed tasks
	if len(completed) > 0 {
		writer.WriteString(completedSection + "\n")
		for _, todo := range completed {
			writer.WriteString(fmt.Sprintf("- [x] %s  \n", todo.Title))
			if todo.Description != "" {
				lines := strings.Split(todo.Description, "\n")
				for _, line := range lines {
					writer.WriteString(fmt.Sprintf("  - %s  \n", line))
				}
			}
		}
		writer.WriteString("\n")
	}
	
	return writer.Flush()
}

// Update the Save and Load methods in TodoStore
func (s *TodoStore) Save() error {
	return s.writeTodoMD()
}

func (s *TodoStore) Load() error {
	return s.parseTodoMD()
}