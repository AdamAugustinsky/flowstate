package main

import (
	"fmt"
	"time"
)

type Todo struct {
	ID          string
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

func (t *Todo) Toggle() {
	t.Completed = !t.Completed
	if t.Completed {
		now := time.Now()
		t.CompletedAt = &now
	} else {
		t.CompletedAt = nil
	}
}

func (t Todo) FilterValue() string {
	return t.Title
}

func (t Todo) DisplayTitle() string {
	prefix := "○ "
	if t.Completed {
		prefix = "✓ "
	}
	
	title := t.Title
	const maxLen = 50
	if len(title) > maxLen {
		title = title[:maxLen-3] + "..."
	}
	
	return prefix + title
}

type TodoStore struct {
	todos    []Todo
	filename string
}

func NewTodoStore(filename string) *TodoStore {
	store := &TodoStore{
		todos:    []Todo{},
		filename: filename,
	}
	store.Load()
	return store
}

func (s *TodoStore) Add(title, description string) {
	todo := Todo{
		ID:          fmt.Sprintf("%d", time.Now().UnixNano()),
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	s.todos = append(s.todos, todo)
	s.Save()
}

func (s *TodoStore) Delete(id string) {
	for i, todo := range s.todos {
		if todo.ID == id {
			s.todos = append(s.todos[:i], s.todos[i+1:]...)
			s.Save()
			return
		}
	}
}

func (s *TodoStore) Toggle(id string) {
	for i := range s.todos {
		if s.todos[i].ID == id {
			s.todos[i].Toggle()
			s.Save()
			return
		}
	}
}

func (s *TodoStore) GetAll() []Todo {
	return s.todos
}

func (s *TodoStore) GetByID(id string) *Todo {
	for i := range s.todos {
		if s.todos[i].ID == id {
			return &s.todos[i]
		}
	}
	return nil
}

func (s *TodoStore) Update(id, title, description string) {
	for i := range s.todos {
		if s.todos[i].ID == id {
			s.todos[i].Title = title
			s.todos[i].Description = description
			s.Save()
			return
		}
	}
}

