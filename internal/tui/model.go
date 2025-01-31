package tui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ncecere/stripper/internal/crawler"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 1)

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF"))

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FF00"))

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000"))
)

type Model struct {
	url       string
	progress  progress.Model
	spinner   spinner.Model
	stopwatch stopwatch.Model
	stats     crawler.Stats
	ready     bool
	quitting  bool
	err       error
	statsChan chan crawler.Stats
	doneChan  chan bool
	startTime time.Time
}

func NewModel(url string, statsChan chan crawler.Stats, doneChan chan bool) Model {
	p := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(40),
	)
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	sw := stopwatch.NewWithInterval(time.Second)

	return Model{
		url:       url,
		progress:  p,
		spinner:   s,
		stopwatch: sw,
		statsChan: statsChan,
		doneChan:  doneChan,
		startTime: time.Now(),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		m.stopwatch.Init(),
		waitForStats(m.statsChan),
		waitForDone(m.doneChan),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.ready {
			m.ready = true
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		}

	case crawler.Stats:
		m.stats = msg
		if m.stats.TotalURLs > 0 {
			progressCmd := m.progress.SetPercent(float64(m.stats.URLsProcessed) / float64(m.stats.TotalURLs))
			cmds = append(cmds, progressCmd)
		}
		cmds = append(cmds, waitForStats(m.statsChan))

	case bool:
		// Done message
		m.quitting = true
		// Print completion message
		fmt.Printf("\nCrawling complete!\n\n")
		return m, tea.Quit

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		cmds = append(cmds, cmd)

	case spinner.TickMsg:
		spinnerModel, cmd := m.spinner.Update(msg)
		m.spinner = spinnerModel
		cmds = append(cmds, cmd)

	case stopwatch.TickMsg:
		stopwatchModel, cmd := m.stopwatch.Update(msg)
		m.stopwatch = stopwatchModel
		cmds = append(cmds, cmd)
	}

	if !m.quitting {
		cmds = append(cmds, m.spinner.Tick)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if !m.ready {
		return "Initializing...\n"
	}

	var percentage float64
	if m.stats.TotalURLs > 0 {
		percentage = float64(m.stats.URLsProcessed) / float64(m.stats.TotalURLs) * 100
	}

	return fmt.Sprintf("\rProcessed: %d/%d (%.1f%%) %s %s Successful: %d, Failed: %d [%s]",
		m.stats.URLsProcessed,
		m.stats.TotalURLs,
		percentage,
		m.spinner.View(),
		m.progress.View(),
		m.stats.URLsSuccessful,
		m.stats.URLsFailed,
		m.stopwatch.View(),
	)
}

func getProgress(stats crawler.Stats) float64 {
	if stats.TotalURLs > 0 {
		return float64(stats.URLsProcessed) / float64(stats.TotalURLs) * 100
	}
	return 0
}

func waitForStats(statsChan chan crawler.Stats) tea.Cmd {
	return func() tea.Msg {
		return <-statsChan
	}
}

func waitForDone(doneChan chan bool) tea.Cmd {
	return func() tea.Msg {
		return <-doneChan
	}
}
