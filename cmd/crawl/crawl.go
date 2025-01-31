package crawl

import (
	"fmt"
	"os"
	"path"
	"time"

	"stripper/internal/crawler"

	"github.com/spf13/cobra"
)

type CrawlOptions struct {
	URL            string
	Depth          int
	Format         string
	Force          bool
	Ignore         []string
	OutputDir      string
	RescanInterval string
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
	cmd.Flags().IntVarP(&opts.Depth, "depth", "d", 1, "Maximum crawl depth")
	cmd.Flags().StringVarP(&opts.Format, "format", "f", "markdown", "Output format (markdown, text, html)")
	cmd.Flags().BoolVar(&opts.Force, "force", false, "Force re-crawl of already crawled URLs")
	cmd.Flags().StringSliceVarP(&opts.Ignore, "ignore", "i", []string{
		"pdf", "jpg", "jpeg", "png", "gif", "css", "js", "ico",
		"woff", "woff2", "ttf", "eot", "mp4", "webm", "mp3", "wav",
		"zip", "tar", "gz", "rar",
	}, "File extensions to ignore")
	cmd.Flags().StringVarP(&opts.OutputDir, "output", "o", "output", "Output directory for crawled content")
	cmd.Flags().StringVarP(&opts.RescanInterval, "rescan", "r", "24h", "Rescan interval for previously crawled pages (e.g., 24h, 1h30m, 15m)")

	return cmd
}

func runCrawl(opts *CrawlOptions) error {
	// Create output directory if it doesn't exist
	outputDir := path.Clean(opts.OutputDir)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Parse rescan interval
	rescanInterval, err := time.ParseDuration(opts.RescanInterval)
	if err != nil {
		return fmt.Errorf("invalid rescan interval format (use format like 24h, 1h30m, 15m): %w", err)
	}

	// Initialize crawler
	c, err := crawler.New(crawler.Options{
		URL:            opts.URL,
		Depth:          opts.Depth,
		Format:         opts.Format,
		Force:          opts.Force,
		Ignore:         opts.Ignore,
		OutputDir:      outputDir,
		RescanInterval: rescanInterval,
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
