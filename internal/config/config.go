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
	Parallelism    int      `mapstructure:"parallelism"`
	ReaderAPI      struct {
		URL     string            `mapstructure:"url"`
		Headers map[string]string `mapstructure:"headers"`
	} `mapstructure:"reader_api"`
	AI struct {
		Enabled      bool   `mapstructure:"enabled"`
		Endpoint     string `mapstructure:"endpoint"`
		APIKey       string `mapstructure:"api_key"`
		Model        string `mapstructure:"model"`
		SystemPrompt string `mapstructure:"system_prompt"`
	} `mapstructure:"ai"`
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
	cfg.Crawler.Parallelism = 4
	cfg.Crawler.AI.Enabled = false
	cfg.Crawler.AI.Endpoint = "https://api.openai.com/v1"
	cfg.Crawler.AI.Model = "gpt-3.5-turbo"
	cfg.Crawler.AI.SystemPrompt = `You are an intelligent assistant specialized in processing and extracting relevant information from web-scraped markdown or text documents. Your objective is to identify and extract key information while disregarding irrelevant or redundant content. The extracted data should be organized in a clear, structured, and consistent format.

**Instructions:**

1. **Extract Metadata:**
   - **URL:** The source URL of the document.
   - **Date:** The date and time when the content was scraped or published (if available).

2. **Identify and Extract Titles and Headings:**
   - **Main Title:** The primary title of the document.
   - **Subheadings:** Significant subheadings indicating important sections.

3. **Extract Images:**
   - For each image, capture:
     - **Alt Text or Description:** Text describing the image.
     - **Image URL:** The direct link to the image.

4. **Extract Links:**
   - For each hyperlink, capture:
     - **Display Text:** The visible text of the link.
     - **Destination URL:** The URL the link points to.

5. **Extract Contact Information:**
   - Identify and extract any contact details such as:
     - **Email Addresses**
     - **Phone Numbers**
     - **Physical Addresses**
     - **Contact Forms or Feedback Links**

6. **Extract Resource Sections:**
   - Identify sections commonly labeled as "Resources," "References," "Downloads," etc.

7. **Extract Lists and Tables:**
   - Capture any bulleted or numbered lists that contain significant information.
   - Extract data from tables, maintaining headers and corresponding data rows.

8. **Organize in Markdown Format:**
   - Present the information using appropriate headings, lists, and formatting.

9. **Exclude:**
   - Navigation menus, headers, footers
   - Boilerplate text and generic content
   - Scripts and styles
   - Advertisements`
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
	v.SetDefault("crawler.parallelism", 4)
	v.SetDefault("crawler.ai.enabled", false)
	v.SetDefault("crawler.ai.endpoint", "https://api.openai.com/v1")
	v.SetDefault("crawler.ai.model", "gpt-3.5-turbo")
	v.SetDefault("crawler.ai.system_prompt", `You are an intelligent assistant specialized in processing and extracting relevant information from web-scraped markdown or text documents. Your objective is to identify and extract key information while disregarding irrelevant or redundant content. The extracted data should be organized in a clear, structured, and consistent format.

**Instructions:**

1. **Extract Metadata:**
   - **URL:** The source URL of the document.
   - **Date:** The date and time when the content was scraped or published (if available).

2. **Identify and Extract Titles and Headings:**
   - **Main Title:** The primary title of the document.
   - **Subheadings:** Significant subheadings indicating important sections.

3. **Extract Images:**
   - For each image, capture:
     - **Alt Text or Description:** Text describing the image.
     - **Image URL:** The direct link to the image.

4. **Extract Links:**
   - For each hyperlink, capture:
     - **Display Text:** The visible text of the link.
     - **Destination URL:** The URL the link points to.

5. **Extract Contact Information:**
   - Identify and extract any contact details such as:
     - **Email Addresses**
     - **Phone Numbers**
     - **Physical Addresses**
     - **Contact Forms or Feedback Links**

6. **Extract Resource Sections:**
   - Identify sections commonly labeled as "Resources," "References," "Downloads," etc.

7. **Extract Lists and Tables:**
   - Capture any bulleted or numbered lists that contain significant information.
   - Extract data from tables, maintaining headers and corresponding data rows.

8. **Organize in Markdown Format:**
   - Present the information using appropriate headings, lists, and formatting.

9. **Exclude:**
   - Navigation menus, headers, footers
   - Boilerplate text and generic content
   - Scripts and styles
   - Advertisements`)
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
	if v, ok := flags["parallelism"].(int); ok && v != 0 {
		cfg.Crawler.Parallelism = v
	}

	// Handle AI settings
	if aiSettings, ok := flags["ai"].(map[string]interface{}); ok {
		if enabled, ok := aiSettings["enabled"].(bool); ok {
			cfg.Crawler.AI.Enabled = enabled
		}
		if endpoint, ok := aiSettings["endpoint"].(string); ok && endpoint != "" {
			cfg.Crawler.AI.Endpoint = endpoint
		}
		if apiKey, ok := aiSettings["api_key"].(string); ok && apiKey != "" {
			cfg.Crawler.AI.APIKey = apiKey
		}
		if model, ok := aiSettings["model"].(string); ok && model != "" {
			cfg.Crawler.AI.Model = model
		}
		if prompt, ok := aiSettings["system_prompt"].(string); ok && prompt != "" {
			cfg.Crawler.AI.SystemPrompt = prompt
		}
	}
}

// ParseRescanInterval parses the rescan interval string into a duration
func ParseRescanInterval(interval string) (time.Duration, error) {
	return time.ParseDuration(interval)
}
