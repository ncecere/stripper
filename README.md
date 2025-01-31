# Stripper - Web Content Crawler

A powerful web content crawler that uses the Reader API to extract clean, readable content from web pages. Built with Go, it features recursive crawling, content archiving, and a terminal user interface.

## Features

- Recursive web crawling with configurable depth
- Clean content extraction via Reader API
- Multiple output formats (markdown, text, html)
- Progress tracking with TUI
- Configurable rescan intervals
- Extension-based filtering
- SQLite-based URL tracking

## Installation

### Using Pre-built Binaries

1. Visit the [Releases](https://github.com/ncecere/stripper/releases) page
2. Download the appropriate binary for your system:
   - Linux: `stripper_Linux_x86_64.tar.gz` or `stripper_Linux_arm64.tar.gz`
   - macOS: `stripper_Darwin_x86_64.tar.gz` or `stripper_Darwin_arm64.tar.gz`
   - Windows: `stripper_Windows_x86_64.zip`
3. Extract the archive:
   ```bash
   # Linux/macOS
   tar xzf stripper_*_*.tar.gz
   
   # Windows (using PowerShell)
   Expand-Archive stripper_Windows_x86_64.zip
   ```
4. Move the binary to your PATH:
   ```bash
   # Linux/macOS
   sudo mv stripper /usr/local/bin/

   # Windows (using PowerShell as Administrator)
   Move-Item stripper.exe C:\Windows\System32\
   ```

### Using Go Install

```bash
go install github.com/ncecere/stripper@latest
```

### Building from Source

1. Clone the repository:
```bash
git clone https://github.com/ncecere/stripper.git
cd stripper
```

2. Build the binary:
```bash
make build
```

Or manually:
```bash
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
  --output ./content \
  --ignore "pdf,jpg,png" \
  --rescan 24h
```

### Configuration

You can configure Stripper using a YAML configuration file. Create `.stripper.yaml` in your home directory or the current directory:

```yaml
crawler:
  depth: 2
  format: markdown
  output_dir: output
  ignore_extensions:
    - pdf
    - jpg
    - png
  rescan_interval: 24h
  reader_api:
    url: https://read.tabnot.space
    headers:
      X-Respond-With: text

http:
  timeout: 30
  retry_attempts: 3
  retry_delay: 5
  user_agent: "Stripper/1.0 Web Content Crawler"
  request_delay: 1000
```

### Command Line Options

- `--depth, -d`: Maximum crawl depth (default: 1)
- `--format, -f`: Output format (markdown, text, html) (default: markdown)
- `--output, -o`: Output directory (default: output)
- `--ignore, -i`: File extensions to ignore
- `--rescan, -r`: Rescan interval (e.g., 24h, 1h30m)
- `--force`: Force re-crawl of already crawled URLs
- `--config, -c`: Path to config file
- `--reader-api-url`: Reader API base URL

## Development

### Requirements

- Go 1.21 or later
- Make (optional, for using Makefile commands)

### Setup

1. Install dependencies:
```bash
go mod download
```

2. Run tests:
```bash
make test
```

### Available Make Commands

- `make`: Run lint, test, and build
- `make build`: Build the binary
- `make test`: Run tests
- `make coverage`: Generate test coverage report
- `make lint`: Run linter
- `make clean`: Remove binary and artifacts
- `make install`: Install binary to GOPATH

## License

MIT License

## Author

Nick Cecere (@ncecere)
