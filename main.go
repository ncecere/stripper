package main

import (
	"fmt"
	"os"

	"stripper/cmd/crawl"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "stripper",
		Short: "A web content scraper and archiver",
		Long: `Stripper is a CLI tool that systematically crawls and archives web content 
from specified domains. It leverages the Reader API to retrieve content in a 
structured format, making it ideal for archiving documentation, articles, 
and other textual resources.`,
	}

	// Add commands
	rootCmd.AddCommand(crawl.NewCrawlCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
