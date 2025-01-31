package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ncecere/stripper/internal/crawler"
	"github.com/ncecere/stripper/internal/storage"
	"github.com/ncecere/stripper/internal/tui"
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
)

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

		// Print initial information
		fmt.Printf("Content will be saved to: %s, Crawling: %s, Depth: %d\n\n", outputDir, url, maxDepth)

		// Create channels for communication
		statsChan := make(chan crawler.Stats, 100)
		doneChan := make(chan bool)
		errChan := make(chan error, 1)

		// Set up progress callback
		c.SetProgressCallback(func(stats crawler.Stats) {
			select {
			case statsChan <- stats:
			default:
				// Skip update if channel is full
			}
		})

		// Create a context with cancellation
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Start crawling in a goroutine
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
			errChan <- err
			doneChan <- true
		}()

		// Start the TUI
		if err := tui.StartTUI(url, statsChan, doneChan); err != nil {
			return fmt.Errorf("TUI error: %w", err)
		}

		// Check for any crawler errors
		if err := <-errChan; err != nil {
			return fmt.Errorf("crawler error: %w", err)
		}

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
