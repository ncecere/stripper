# Stripper - Web Content Crawler

A CLI tool for systematically crawling and archiving web content using the Reader API.

## Features

- **Automated Crawling**
  - Recursive crawling with configurable depth
  - Smart link discovery and tracking
  - Batch processing with rate limiting
  - Progress tracking and statistics

- **Content Management**
  - Stores content in markdown, text, or HTML format
  - SQLite database for URL tracking
  - Configurable rescan intervals
  - Force flag for complete recrawls

- **Error Handling**
  - Robust retry logic
  - Detailed error tracking
  - Domain boundary enforcement
  - Extension-based filtering

## Installation

```bash
go install github.com/yourusername/stripper@latest
```

## Usage

Basic usage:
```bash
stripper crawl https://example.com
```

With options:
```bash
stripper crawl https://example.com \
  --depth 2 \
  --format markdown \
  --force \
  --rescan 24h \
  --reader-api-url https://read.tabnot.space
```

### Command Line Options

- `--config, -c`: Path to config file (default: .stripper.yaml in current directory)
- `--depth, -d`: Maximum crawl depth (default: 1)
- `--format, -f`: Output format - markdown, text, or html (default: markdown)
- `--force`: Force re-crawl of already crawled URLs
- `--ignore, -i`: File extensions to ignore
- `--output, -o`: Output directory for crawled content (default: output)
- `--rescan, -r`: Rescan interval for previously crawled pages (e.g., 24h, 1h30m, 15m)
- `--reader-api-url`: Reader API base URL (default: https://read.tabnot.space)

### Configuration File

The application looks for configuration in the following locations (in order):
1. Path specified by --config flag
2. .stripper.yaml in current directory
3. $HOME/.stripper.yaml
4. /etc/stripper/config.yaml

See `.stripper.yaml.example` for a complete example of the configuration format.

```yaml
crawler:
  depth: 2
  format: "markdown"
  output_dir: "output"
  rescan_interval: "24h"
  reader_api:
    url: "https://read.tabnot.space"
```

## How It Works

1. **Link Discovery**
   - Starts from the provided URL
   - Recursively discovers links up to specified depth
   - Filters external domains and ignored extensions

2. **Content Processing**
   - Uses Reader API to extract clean content
   - Supports multiple output formats
   - Maintains metadata about crawl status

3. **Smart Recrawling**
   - Tracks last crawl time for each URL
   - Respects configured rescan interval
   - Force flag available for complete recrawls

4. **Progress Tracking**
   - Real-time progress display
   - Crawl statistics and status
   - Detailed debug logging (with STRIPPER_DEBUG=1)

## Development

Build from source:
```bash
git clone https://github.com/yourusername/stripper.git
cd stripper
go build
```

Run with debug logging:
```bash
STRIPPER_DEBUG=1 ./stripper crawl https://example.com
```

## License

MIT License - see LICENSE file for details.
