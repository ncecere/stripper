package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ncecere/stripper/internal/crawler"
	"github.com/ncecere/stripper/internal/storage"
	"github.com/ncecere/stripper/pkg/reader"
)

var (
	cfgFile      string
	outputDir    string
	batchSize    int
	delay        time.Duration
	maxRetries   int
	format       string
	maxDepth     int
	stayInDomain bool
	styleTitle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
	styleError   = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	styleInfo    = lipgloss.NewStyle().Foreground(lipgloss.Color("86"))
	stylePrompt  = lipgloss.NewStyle().Foreground(lipgloss.Color("99"))
)

type frameMsg struct{}

type model struct {
	url      string
	stats    crawler.Stats
	err      error
	quitting bool
	done     bool
	lastLine string
}

func (m model) Init() tea.Cmd {
	return tickCmd()
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second/2, func(t time.Time) tea.Msg {
		return frameMsg{}
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case frameMsg:
		return m, tickCmd()
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			m.quitting = true
			return m, tea.Quit
		}
	case crawler.Stats:
		m.stats = msg
		if m.stats.TotalURLs > 0 && m.stats.URLsProcessed >= m.stats.TotalURLs {
			m.done = true
			return m, tea.Quit
		}
	case error:
		m.err = msg
		return m, tea.Quit
	case bool: // completion signal
		m.done = msg
		return m, tea.Quit
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("\r%s\n", styleError.Render(fmt.Sprintf("Error: %v", m.err)))
	}

	var percentage float64
	if m.stats.TotalURLs > 0 {
		percentage = float64(m.stats.URLsProcessed) / float64(m.stats.TotalURLs) * 100
	}

	m.lastLine = fmt.Sprintf("Processed: %d/%d URLs (%.1f%%)",
		m.stats.URLsProcessed,
		m.stats.TotalURLs,
		percentage,
	)

	if m.done {
		return fmt.Sprintf("\r%s\n", m.lastLine)
	}

	return fmt.Sprintf("\r%s", m.lastLine)
}

var rootCmd = &cobra.Command{
	Use:   "stripper",
	Short: "A web content scraper and archiver",
	Long: `Stripper is a CLI tool that systematically crawls and archives web content 
from specified domains. It uses the Reader API to retrieve content in a structured format, 
making it ideal for archiving documentation, articles, and other textual resources.`,
}

var crawlCmd = &cobra.Command{
	Use:   "crawl [url]",
	Short: "Crawl and archive a URL",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		url := args[0]

		// Initialize storage
		store, err := storage.NewSQLiteStorage(outputDir)
		if err != nil {
			return fmt.Errorf("failed to initialize storage: %w", err)
		}
		defer store.Close()

		// Initialize reader client
		readerClient := reader.NewClient(reader.Config{
			DefaultFormat: reader.Format(format),
		})

		// Initialize crawler
		c := crawler.NewCrawler(readerClient, store)

		// Create and configure the TUI model
		m := model{
			url: url,
		}

		fmt.Printf("Content will be saved to: %s\n\n", outputDir)
		fmt.Printf("Crawling: %s\n\n", url)

		// Create a channel for stats updates
		statsChan := make(chan crawler.Stats, 100)
		defer close(statsChan)

		// Set up progress callback
		c.SetProgressCallback(func(stats crawler.Stats) {
			select {
			case statsChan <- stats:
			default:
				// Skip update if channel is full
			}
		})

		// Create program without alternate screen
		p := tea.NewProgram(
			&m,
			tea.WithOutput(os.Stdout),
		)

		// Start a goroutine to forward stats to the program
		go func() {
			for stats := range statsChan {
				p.Send(stats)
			}
		}()

		// Create a context with cancellation
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Start crawling in a goroutine
		errCh := make(chan error, 1)
		doneCh := make(chan struct{})

		go func() {
			config := crawler.Config{
				BatchSize:    batchSize,
				RequestDelay: delay,
				MaxRetries:   maxRetries,
				BaseURL:      url,
				OutputDir:    outputDir,
				MaxDepth:     maxDepth,
				StayInDomain: stayInDomain,
				Format:       reader.Format(format),
			}
			err := c.Start(ctx, config)
			errCh <- err
			close(doneCh)
		}()

		// Start another goroutine to signal completion
		go func() {
			<-doneCh
			p.Send(true) // Send completion signal
		}()

		// Run the TUI
		finalModel, err := p.Run()
		if err != nil {
			cancel() // Cancel context on UI error
			return fmt.Errorf("failed to run UI: %w", err)
		}

		// Check for crawler errors
		if err := <-errCh; err != nil {
			return fmt.Errorf("crawler error: %w", err)
		}

		// Print final stats
		if m, ok := finalModel.(model); ok && !m.quitting {
			fmt.Printf("%s\n", styleInfo.Render("Crawling complete!"))
		}

		// Ensure proper cleanup
		c.Stop()

		return nil
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.stripper.yaml)")
	rootCmd.PersistentFlags().StringVar(&outputDir, "output", "", "output directory for archived content")
	rootCmd.PersistentFlags().IntVar(&batchSize, "batch-size", 5, "number of URLs to process in parallel")
	rootCmd.PersistentFlags().DurationVar(&delay, "delay", time.Second, "delay between requests")
	rootCmd.PersistentFlags().IntVar(&maxRetries, "max-retries", 3, "maximum number of retry attempts")
	rootCmd.PersistentFlags().StringVar(&format, "format", "markdown", "output format (text, markdown, html)")
	rootCmd.PersistentFlags().IntVar(&maxDepth, "depth", 0, "maximum crawl depth (0 for unlimited)")
	rootCmd.PersistentFlags().BoolVar(&stayInDomain, "stay-in-domain", true, "only crawl URLs in the same domain")

	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
	viper.BindPFlag("batch-size", rootCmd.PersistentFlags().Lookup("batch-size"))
	viper.BindPFlag("delay", rootCmd.PersistentFlags().Lookup("delay"))
	viper.BindPFlag("max-retries", rootCmd.PersistentFlags().Lookup("max-retries"))
	viper.BindPFlag("format", rootCmd.PersistentFlags().Lookup("format"))
	viper.BindPFlag("depth", rootCmd.PersistentFlags().Lookup("depth"))
	viper.BindPFlag("stay-in-domain", rootCmd.PersistentFlags().Lookup("stay-in-domain"))

	rootCmd.AddCommand(crawlCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".stripper")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	// Set default output directory if not specified
	if outputDir == "" {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		outputDir = filepath.Join(home, ".stripper")
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
