package tui

import (
	"fmt"
	"strings"
	"time"

	"stripper/internal/database"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// TickMsg is sent when the stats should update
type TickMsg time.Time

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF75B7"))

	urlStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575"))

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000"))
)

type model struct {
	db     *database.DB
	width  int
	height int
}

func New(db *database.DB) *tea.Program {
	m := &model{
		db: db,
	}
	return tea.NewProgram(m)
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}

	case TickMsg:
		return m, nil
	}

	return m, nil
}

func (m model) View() string {
	if m.width == 0 {
		return "Initializing..."
	}

	var b strings.Builder

	// Title
	b.WriteString(titleStyle.Render("Stripper - Web Content Crawler") + "\n\n")

	// Stats
	total, pending, completed, failed, err := m.db.GetStats()
	if err != nil {
		b.WriteString(errorStyle.Render(fmt.Sprintf("Error getting stats: %v\n", err)))
	} else {
		progress := 0.0
		if total > 0 {
			progress = float64(completed+failed) / float64(total) * 100
		}

		b.WriteString(fmt.Sprintf("Progress: %.1f%% (%d/%d URLs)\n", progress, completed+failed, total))
		b.WriteString(fmt.Sprintf("Status:\n"))
		b.WriteString(fmt.Sprintf("  • Completed: %d\n", completed))
		b.WriteString(fmt.Sprintf("  • Pending: %d\n", pending))
		if failed > 0 {
			b.WriteString(fmt.Sprintf("  • Failed: %d\n", failed))
		}
	}

	// Help
	b.WriteString("\nPress q to quit")

	// Left-align content
	return lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Render(b.String())
}
