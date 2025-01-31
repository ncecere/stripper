# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v0.1.6] - 2025-01-31

### Added
- AI-powered content summarization
- Support for multiple AI models:
  - deepseek-r1 (faster processing)
  - grog-llama-3.1-8b (balanced)
  - grog-llama-3.2-3b (higher quality)
- Intelligent rate limiting with exponential backoff
- Configurable batch processing for AI requests
- Separate storage for AI summaries with flat file organization
- New configuration options for AI features
- Additional command line flags for AI control

## [v0.1.3] - 2025-01-31

### Fixed
- Improved ARM64 build configuration
- Fixed cross-compilation dependency conflicts
- Added proper SQLite support for all architectures

## [v0.1.2] - 2025-01-31

### Fixed
- Fixed cross-compilation setup in GitHub Actions workflow
- Updated dependency installation for multi-architecture builds
- Improved ARM64 support with proper SQLite compilation
- Enhanced build configuration for all platforms

## [v0.1.1] - 2025-01-31

### Fixed
- Enable CGO for SQLite database support
- Add cross-compilation dependencies for all platforms
- Include SQLite development libraries in build process

### Added
- SQLite dependency requirements in documentation
- Platform-specific installation instructions for SQLite

### Changed
- Updated build configuration in GoReleaser
- Improved GitHub Actions workflow for native SQLite support

## [v0.1.0] - 2025-01-31

### Added
- Initial release
- Web content crawling with configurable depth
- Clean content extraction via Reader API
- Multiple output formats (markdown, text, html)
- Progress tracking with TUI
- Configurable rescan intervals
- Extension-based filtering
- SQLite-based URL tracking
- YAML configuration support
- Command-line interface
- Automated releases with GitHub Actions
- Cross-platform builds (Linux, macOS, Windows)
- Comprehensive documentation
- MIT License

[v0.1.6]: https://github.com/ncecere/stripper/releases/tag/v0.1.6
[v0.1.3]: https://github.com/ncecere/stripper/releases/tag/v0.1.3
[v0.1.2]: https://github.com/ncecere/stripper/releases/tag/v0.1.2
[v0.1.1]: https://github.com/ncecere/stripper/releases/tag/v0.1.1
[v0.1.0]: https://github.com/ncecere/stripper/releases/tag/v0.1.0
