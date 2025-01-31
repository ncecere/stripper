package crawl

import (
	"fmt"
	"os"
	"path"

	"stripper/internal/config"
	"stripper/internal/crawler"

	"github.com/spf13/cobra"
)

type CrawlOptions struct {
	URL            string
	ConfigFile     string
	Depth          int
	Format         string
	Force          bool
	Ignore         []string
	OutputDir      string
	RescanInterval string
	ReaderAPIURL   string
	Parallelism    int
}

// findConfigFile looks for config in standard locations
func findConfigFile(configPath string) string {
	// Check explicit path first
	if configPath != "" {
		if _, err := os.Stat(configPath); err == nil {
			return configPath
		}
	}

	// Check standard locations
	locations := []string{
		".stripper.yaml",
		path.Join(os.Getenv("HOME"), ".stripper.yaml"),
		"/etc/stripper/config.yaml",
	}

	for _, loc := range locations {
		if _, err := os.Stat(loc); err == nil {
			return loc
		}
	}

	return ""
}

func NewCrawlCmd() *cobra.Command {
	opts := &CrawlOptions{}

	cmd := &cobra.Command{
		Use:   "crawl [url]",
		Short: "Crawl and archive web content",
		Long: `Crawl and archive web content from a specified URL.
The content will be retrieved using the Reader API and stored locally.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.URL = args[0]
			return runCrawl(opts)
		},
	}

	// Add flags
	cmd.Flags().StringVarP(&opts.ConfigFile, "config", "c", "", "Path to config file")
	cmd.Flags().IntVarP(&opts.Depth, "depth", "d", 1, "Maximum crawl depth")
	cmd.Flags().IntVarP(&opts.Parallelism, "parallel", "p", 4, "Number of parallel workers")
	cmd.Flags().StringVarP(&opts.Format, "format", "f", "markdown", "Output format (markdown, text, html)")
	cmd.Flags().BoolVar(&opts.Force, "force", false, "Force re-crawl of already crawled URLs")
	cmd.Flags().StringSliceVarP(&opts.Ignore, "ignore", "i", []string{
		"pdf", "jpg", "jpeg", "png", "gif", "css", "js", "ico",
		"woff", "woff2", "ttf", "eot", "mp4", "webm", "mp3", "wav",
		"zip", "tar", "gz", "rar",
	}, "File extensions to ignore")
	cmd.Flags().StringVarP(&opts.OutputDir, "output", "o", "output", "Output directory for crawled content")
	cmd.Flags().StringVarP(&opts.RescanInterval, "rescan", "r", "24h", "Rescan interval for previously crawled pages (e.g., 24h, 1h30m, 15m)")
	cmd.Flags().StringVar(&opts.ReaderAPIURL, "reader-api-url", "https://read.tabnot.space", "Reader API base URL")

	return cmd
}

func runCrawl(opts *CrawlOptions) error {
	// Load configuration
	configPath := findConfigFile(opts.ConfigFile)
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		if opts.ConfigFile != "" {
			// Only return error if user explicitly specified a config file
			return fmt.Errorf("error loading config file: %w", err)
		}
		// Otherwise, use defaults
		cfg = &config.Config{}
		config.SetDefaults(cfg)
	}

	// Merge command line flags with config
	flags := map[string]interface{}{
		"depth":          opts.Depth,
		"format":         opts.Format,
		"output":         opts.OutputDir,
		"ignore":         opts.Ignore,
		"rescan":         opts.RescanInterval,
		"reader-api-url": opts.ReaderAPIURL,
		"parallelism":    opts.Parallelism,
	}
	config.MergeWithFlags(cfg, flags)

	// Create output directory if it doesn't exist
	outputDir := path.Clean(cfg.Crawler.OutputDir)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Parse rescan interval
	rescanInterval, err := config.ParseRescanInterval(cfg.Crawler.RescanInterval)
	if err != nil {
		return fmt.Errorf("invalid rescan interval format (use format like 24h, 1h30m, 15m): %w", err)
	}

	// Initialize crawler
	c, err := crawler.New(crawler.Options{
		URL:            opts.URL,
		Depth:          cfg.Crawler.Depth,
		Format:         cfg.Crawler.Format,
		Force:          opts.Force,
		Ignore:         cfg.Crawler.IgnoreExts,
		OutputDir:      outputDir,
		RescanInterval: rescanInterval,
		ReaderAPIURL:   cfg.Crawler.ReaderAPI.URL,
		Parallelism:    cfg.Crawler.Parallelism,
	})
	if err != nil {
		return fmt.Errorf("failed to initialize crawler: %w", err)
	}

	// Start crawling
	if err := c.Start(); err != nil {
		return fmt.Errorf("crawling failed: %w", err)
	}

	return nil
}
