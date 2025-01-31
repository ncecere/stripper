# Stripper - Web Content Crawler

A command-line tool for systematically crawling and archiving web content from specified domains. It uses the Reader API to retrieve clean, formatted content and stores it locally.

## Features

- Domain-specific crawling with configurable depth
- Clean content extraction via Reader API
- Multiple output formats (text, markdown, html)
- SQLite-based link tracking and crawl history
- Rate limiting and polite crawling
- Real-time progress monitoring via TUI

## Installation

```bash
go install github.com/yourusername/stripper@latest
```

Or build from source:

```bash
git clone https://github.com/yourusername/stripper.git
cd stripper
go build -o stripper
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
  --ignore pdf,jpg,css \
  --output ./archive
```

### Command Line Flags

| Flag | Description | Type | Default | Example |
|------|-------------|------|---------|---------|
| `--depth`, `-d` | Maximum crawl depth from starting URL | int | 1 | `--depth 2` |
| `--format`, `-f` | Output format (markdown, text, html) | string | "markdown" | `--format text` |
| `--force` | Force re-crawl of already crawled URLs | bool | false | `--force` |
| `--ignore`, `-i` | File extensions to ignore | []string | [pdf,jpg,jpeg,png,gif,css,js,ico,woff,woff2,ttf,eot,mp4,webm,mp3,wav,zip,tar,gz,rar] | `--ignore pdf,jpg,css` |
| `--output`, `-o` | Output directory for crawled content | string | "output" | `--output ./archive` |

### Configuration File

The application supports configuration via `.stripper.yaml` in the current directory or home directory. Example configuration:

```yaml
# Crawler settings
crawler:
  # Maximum depth to crawl (default: 1)
  depth: 2
  
  # Output format: markdown, text, or html (default: markdown)
  format: "markdown"
  
  # Output directory for crawled content (default: output)
  output_dir: "output"
  
  # File extensions to ignore during crawling
  ignore_extensions:
    - pdf
    - jpg
    - png
    - gif
    - zip
    - exe
    - doc
    - docx
    - xls
    - xlsx

# HTTP client settings
http:
  # Request timeout in seconds
  timeout: 30
  
  # Number of retry attempts for failed requests
  retry_attempts: 3
  
  # Delay between retries in seconds
  retry_delay: 5
  
  # User agent string for requests
  user_agent: "Stripper/1.0 Web Content Crawler"
  
  # Delay between requests in milliseconds
  request_delay: 1000
```

### Examples

1. Basic crawl of a website (stays within the domain):
```bash
stripper crawl https://it.ufl.edu
# Only crawls it.ufl.edu domain, ignores links to other domains
```

2. Deep crawl with text output:
```bash
stripper crawl https://it.ufl.edu \
  --depth 3 \
  --format text \
  --output ./docs
# Crawls up to 3 levels deep within it.ufl.edu
```

3. Force re-crawl with debug logging:
```bash
STRIPPER_DEBUG=1 stripper crawl https://it.ufl.edu \
  --force \
  --format text
# Shows detailed debug output during crawling
```

4. Crawl with custom ignore patterns and rate limiting:
```bash
stripper crawl https://it.ufl.edu \
  --ignore pdf,jpg,doc,docx \
  --depth 2
# Respects rate limiting from .stripper.yaml
```

5. Using configuration file:
```bash
# Create .stripper.yaml in your directory
cp .stripper.yaml.example .stripper.yaml
# Edit settings as needed
stripper crawl https://it.ufl.edu
```

### Debug Mode

Enable debug logging by setting the `STRIPPER_DEBUG` environment variable:

```bash
STRIPPER_DEBUG=1 stripper crawl https://example.com
```

## Output Format

Content is saved with metadata headers:
```text
URL: https://example.com/page
Date: 2025-01-31T07:36:47-05:00

[Content follows here...]
```

Files are saved using URL-based paths in the output directory:
```
output/
  example.com/
    index.txt
    about.txt
    docs/
      getting-started.txt
```

## Database

The application uses SQLite to track:
- Crawled URLs and their status
- Last crawl dates
- Error history
- Crawl statistics

The database is stored in `{output_dir}/crawler.db`.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
