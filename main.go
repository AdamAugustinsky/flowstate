package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	var globalFlag bool
	flag.BoolVar(&globalFlag, "g", false, "Use global TODO.md from config directory")
	flag.Parse()

	var todoPath string
	if globalFlag {
		configDir, err := os.UserConfigDir()
		if err != nil {
			fmt.Printf("Error getting config directory: %v\n", err)
			os.Exit(1)
		}
		// Create flowstate config directory if it doesn't exist
		flowstateDir := filepath.Join(configDir, "flowstate")
		if err := os.MkdirAll(flowstateDir, 0755); err != nil {
			fmt.Printf("Error creating config directory: %v\n", err)
			os.Exit(1)
		}
		todoPath = filepath.Join(flowstateDir, "TODO.md")
	} else {
		todoPath = "TODO.md"
	}

	p := tea.NewProgram(initialModel(todoPath, globalFlag))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}