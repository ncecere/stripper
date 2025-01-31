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
- AI-powered content summarization with support for multiple models
- Configurable rate limiting and retry strategies for AI processing
- Parallel processing with adjustable batch sizes

## AI Features

- Generate AI summaries of crawled content
- Support for multiple AI models:
  - deepseek-r1 (faster, lower latency)
  - grog-llama-3.1-8b (balanced)
  - grog-llama-3.2-3b (higher quality)
  - mistral-7b-instruct (fast, high quality)
- Intelligent rate limiting and retry strategies
- Parallel processing with configurable batch sizes
- Separate storage for original content and AI summaries
- Flat file organization for easy navigation

## Installation

### Using Pre-built Binaries

#### Linux (x86_64/ARM64)
```bash
# Download latest release
curl -LO $(curl -s https://api.github.com/repos/ncecere/stripper/releases/latest | grep -o -E "https://.*stripper_Linux_x86_64.tar.gz")

# Extract binary
tar xzf stripper_Linux_x86_64.tar.gz

# Make binary executable and move to PATH
chmod +x stripper
sudo mv stripper /usr/local/bin/

# Verify installation
stripper --version
```

#### macOS (Intel/M1)
```bash
# For Intel Macs (x86_64)
curl -LO $(curl -s https://api.github.com/repos/ncecere/stripper/releases/latest | grep -o -E "https://.*stripper_Darwin_x86_64.tar.gz")
tar xzf stripper_Darwin_x86_64.tar.gz

# For M1/M2 Macs (arm64)
curl -LO $(curl -s https://api.github.com/repos/ncecere/stripper/releases/latest | grep -o -E "https://.*stripper_Darwin_arm64.tar.gz")
tar xzf stripper_Darwin_arm64.tar.gz

# Make binary executable and move to PATH
chmod +x stripper
sudo mv stripper /usr/local/bin/

# Verify installation
stripper --version
```

#### Windows (PowerShell)
```powershell
# Download latest release
$release = Invoke-RestMethod -Uri "https://api.github.com/repos/ncecere/stripper/releases/latest"
$url = $release.assets | Where-Object { $_.name -like "*Windows_x86_64.zip" } | Select-Object -ExpandProperty browser_download_url
Invoke-WebRequest -Uri $url -OutFile "stripper_Windows_x86_64.zip"

# Extract binary
Expand-Archive -Path "stripper_Windows_x86_64.zip" -DestinationPath "."

# Move to PATH (requires Administrator)
Move-Item -Path "stripper.exe" -Destination "C:\Windows\System32\"

# Verify installation
stripper --version
```

#### Manual Download
1. Visit the [Releases](https://github.com/ncecere/stripper/releases) page
2. Download the appropriate binary for your system:
   - Linux: `stripper_Linux_x86_64.tar.gz` or `stripper_Linux_arm64.tar.gz`
   - macOS: `stripper_Darwin_x86_64.tar.gz` or `stripper_Darwin_arm64.tar.gz`
   - Windows: `stripper_Windows_x86_64.zip`
3. Extract and install manually using the commands shown above

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
  
  ai:
    enabled: true
    endpoint: "https://ai.example.com/v1"
    api_key: "your-api-key"
    model: "deepseek-r1"  # or grog-llama-3.1-8b, grog-llama-3.2-3b, mistral-7b-instruct
    system_prompt: "Summarize the following content:"

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
- `--ai`: Enable AI summarization
- `--ai-endpoint`: AI API endpoint URL
- `--ai-key`: AI API key
- `--ai-model`: AI model to use
- `--ai-system-prompt`: System prompt for AI summarization

## Development

### Requirements

- Go 1.21 or later
- Make (optional, for using Makefile commands)
- SQLite development libraries:
  ```bash
  # Ubuntu/Debian
  sudo apt-get install libsqlite3-dev

  # CentOS/RHEL
  sudo yum install sqlite-devel

  # macOS
  brew install sqlite3

  # Windows
  # SQLite is included in the Windows binary
  ```

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

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for release history and version details.

## Author

Nick Cecere (@ncecere)
