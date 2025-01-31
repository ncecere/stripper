package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config holds all configuration settings
type Config struct {
	Crawler CrawlerConfig `mapstructure:"crawler"`
	HTTP    HTTPConfig    `mapstructure:"http"`
}

// CrawlerConfig holds crawler-specific settings
type CrawlerConfig struct {
	Depth          int      `mapstructure:"depth"`
	Format         string   `mapstructure:"format"`
	OutputDir      string   `mapstructure:"output_dir"`
	IgnoreExts     []string `mapstructure:"ignore_extensions"`
	RescanInterval string   `mapstructure:"rescan_interval"`
	ReaderAPI      struct {
		URL     string            `mapstructure:"url"`
		Headers map[string]string `mapstructure:"headers"`
	} `mapstructure:"reader_api"`
}

// HTTPConfig holds HTTP client settings
type HTTPConfig struct {
	Timeout       int    `mapstructure:"timeout"`
	RetryAttempts int    `mapstructure:"retry_attempts"`
	RetryDelay    int    `mapstructure:"retry_delay"`
	UserAgent     string `mapstructure:"user_agent"`
	RequestDelay  int    `mapstructure:"request_delay"`
}

// LoadConfig loads configuration from file and environment variables
func LoadConfig(configPath string) (*Config, error) {
	v := viper.New()

	// Set default values
	setViperDefaults(v)

	// If config file is provided, try to load it
	if configPath != "" {
		v.SetConfigType("yaml")
		v.SetConfigFile(configPath)
		if err := v.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &config, nil
}

// SetDefaults initializes a config struct with default values
func SetDefaults(cfg *Config) {
	// Crawler defaults
	cfg.Crawler.Depth = 1
	cfg.Crawler.Format = "markdown"
	cfg.Crawler.OutputDir = "output"
	cfg.Crawler.IgnoreExts = []string{
		"pdf", "jpg", "jpeg", "png", "gif",
		"css", "js", "ico", "woff", "woff2",
		"ttf", "eot", "mp4", "webm", "mp3",
		"wav", "zip", "tar", "gz", "rar",
	}
	cfg.Crawler.RescanInterval = "24h"
	cfg.Crawler.ReaderAPI.URL = "https://read.tabnot.space"
	cfg.Crawler.ReaderAPI.Headers = map[string]string{
		"X-Respond-With": "text",
	}

	// HTTP client defaults
	cfg.HTTP.Timeout = 30
	cfg.HTTP.RetryAttempts = 3
	cfg.HTTP.RetryDelay = 5
	cfg.HTTP.UserAgent = "Stripper/1.0 Web Content Crawler"
	cfg.HTTP.RequestDelay = 1000
}

// setViperDefaults sets default values for viper configuration
func setViperDefaults(v *viper.Viper) {
	// Crawler defaults
	v.SetDefault("crawler.depth", 1)
	v.SetDefault("crawler.format", "markdown")
	v.SetDefault("crawler.output_dir", "output")
	v.SetDefault("crawler.ignore_extensions", []string{
		"pdf", "jpg", "jpeg", "png", "gif",
		"css", "js", "ico", "woff", "woff2",
		"ttf", "eot", "mp4", "webm", "mp3",
		"wav", "zip", "tar", "gz", "rar",
	})
	v.SetDefault("crawler.rescan_interval", "24h")
	v.SetDefault("crawler.reader_api.url", "https://read.tabnot.space")
	v.SetDefault("crawler.reader_api.headers", map[string]string{
		"X-Respond-With": "text",
	})

	// HTTP client defaults
	v.SetDefault("http.timeout", 30)
	v.SetDefault("http.retry_attempts", 3)
	v.SetDefault("http.retry_delay", 5)
	v.SetDefault("http.user_agent", "Stripper/1.0 Web Content Crawler")
	v.SetDefault("http.request_delay", 1000)
}

// MergeWithFlags merges config file settings with command line flags
func MergeWithFlags(cfg *Config, flags map[string]interface{}) {
	if v, ok := flags["depth"].(int); ok && v != 0 {
		cfg.Crawler.Depth = v
	}
	if v, ok := flags["format"].(string); ok && v != "" {
		cfg.Crawler.Format = v
	}
	if v, ok := flags["output"].(string); ok && v != "" {
		cfg.Crawler.OutputDir = v
	}
	if v, ok := flags["ignore"].([]string); ok && len(v) > 0 {
		cfg.Crawler.IgnoreExts = v
	}
	if v, ok := flags["rescan"].(string); ok && v != "" {
		cfg.Crawler.RescanInterval = v
	}
	if v, ok := flags["reader-api-url"].(string); ok && v != "" {
		cfg.Crawler.ReaderAPI.URL = v
	}
}

// ParseRescanInterval parses the rescan interval string into a duration
func ParseRescanInterval(interval string) (time.Duration, error) {
	return time.ParseDuration(interval)
}
