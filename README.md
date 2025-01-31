# Stripper - Web Content Crawler and Archiver

Stripper is a powerful CLI tool that systematically crawls and archives web content from specified domains. It uses the Reader API to retrieve content in a structured format, making it ideal for archiving documentation, articles, and other textual resources.

## Features

- Recursive web crawling with configurable depth
- Domain-scoped crawling
- Parallel processing with configurable batch size
- Rate limiting to respect server resources
- Multiple output formats (text, markdown, HTML)
- Structured content storage with metadata
- Progress visualization with TUI
- Configurable via command line or config file

## Installation

```bash
# Clone the repository
git clone https://github.com/ncecere/stripper.git
cd stripper

# Build the binary
go build -o stripper cmd/stripper/main.go

# Optionally, move to your PATH
sudo mv stripper /usr/local/bin/
```

## Usage

Basic usage:
```bash
stripper crawl <url>
```

The tool can be controlled using:
- `Ctrl+C`: Gracefully stop the crawling process

### Command Line Arguments

```bash
stripper crawl [url] [flags]

Flags:
  --config string       Config file path (default "$HOME/.stripper.yaml")
  --output string      Output directory for archived content (default "$HOME/.stripper")
  --batch-size int     Number of URLs to process in parallel (default 5)
  --delay duration     Delay between requests (default 1s)
  --max-retries int    Maximum number of retry attempts (default 3)
  --format string      Output format: text, markdown, html (default "markdown")
  --depth int          Maximum crawl depth, 0 for unlimited (default 0)
  --stay-in-domain     Only crawl URLs in the same domain (default true)
```

### Configuration

Stripper can be configured in three ways:
1. Command line flags
2. Configuration file
3. Environment variables

#### Configuration File

Stripper supports configuration through a YAML file. By default, it looks for `.stripper.yaml` in your home directory.

#### Environment Variables

All configuration options can be set using environment variables. The variables should be prefixed with `STRIPPER_` and use uppercase. For example:

```bash
export STRIPPER_OUTPUT="/path/to/output"
export STRIPPER_BATCH_SIZE="10"
export STRIPPER_DELAY="2s"
export STRIPPER_FORMAT="html"
```

Example configuration file:
```yaml
# Output directory for archived content
output: "~/websites/archive"

# Number of URLs to process in parallel
batch-size: 10

# Delay between requests
delay: 2s

# Maximum retry attempts for failed requests
max-retries: 5

# Output format (text, markdown, html)
format: "markdown"

# Maximum crawl depth (0 for unlimited)
depth: 3

# Stay within the same domain
stay-in-domain: true
```

## Examples

1. Basic crawling of a website:
```bash
stripper crawl https://example.com
```

2. Crawl with depth limit and custom output directory:
```bash
stripper crawl https://docs.example.com --depth 2 --output ./docs-archive
```

3. Fast crawling with increased parallelism:
```bash
stripper crawl https://blog.example.com --batch-size 10 --delay 500ms
```

4. Archive in HTML format with unlimited depth:
```bash
stripper crawl https://wiki.example.com --format html --depth 0
```

5. Using a custom config file:
```bash
stripper crawl https://docs.example.com --config ./custom-config.yaml
```

### Output Formats

The tool supports three output formats:

1. `text`: Plain text format
   - Strips all HTML formatting
   - Preserves basic structure and readability
   - Smallest file size

2. `markdown`: Markdown format (default)
   - Preserves headings, lists, and basic formatting
   - Maintains links in markdown syntax
   - Good balance of readability and structure

3. `html`: HTML format
   - Preserves all HTML structure
   - Maintains original formatting
   - Largest file size but most complete representation

## Output Directory Structure

The tool creates the following structure in your output directory:

```
output/
├── content/                 # Archived content files
│   ├── example.com-page1.txt
│   ├── example.com-page2.txt
│   └── ...
├── metadata/               # Metadata files for each page
│   ├── example.com-page1.json
│   ├── example.com-page2.json
│   └── ...
└── crawler.db             # SQLite database for crawl state
```

### Content Files
Content files contain the actual text content of the crawled pages, formatted according to the specified output format (text, markdown, or HTML).

### Metadata Files
Each content file has a corresponding metadata JSON file containing:
- URL
- Title
- Timestamp
- Format
- Properties (latency, depth, etc.)

### Database
The SQLite database tracks:
- Crawled URLs
- Processing status
- Error information
- Relationships between pages

## Error Handling

The tool handles various types of errors:
- Network errors (connection issues, timeouts)
- HTTP errors (404, 500, etc.)
- Parser errors (malformed content)
- Storage errors (disk space, permissions)

Failed URLs are logged with their error types and can be reviewed in the crawler.db database.

## Progress Display

The tool provides a real-time TUI (Terminal User Interface) showing:
- Current URL being processed
- Progress bar
- Processing statistics
  - Total URLs processed
  - Successful crawls
  - Failed crawls
- Error counts by type

## Best Practices

1. Start with a small depth value to test the crawling behavior
2. Adjust batch-size and delay based on the target server's capacity
3. Use the stay-in-domain flag to prevent unwanted external crawling
4. Consider using markdown format for best content preservation
5. Monitor the target server's response times and adjust accordingly

## Limitations

- JavaScript-rendered content is not processed
- Binary files are not archived
- Maximum URL queue size is 1000
- Some websites may block automated crawling

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
