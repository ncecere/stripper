package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ncecere/stripper/internal/crawler"
)

// StartTUI initializes and runs the terminal user interface
func StartTUI(url string, statsChan chan crawler.Stats, doneChan chan bool) error {
	m := NewModel(url, statsChan, doneChan)
	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}

	return nil
}
