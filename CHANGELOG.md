# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v0.1.2] - 2025-01-31

### Fixed
- Fixed cross-compilation setup in GitHub Actions workflow
- Updated dependency installation for multi-architecture builds
- Added separate builds for CGO-enabled and CGO-disabled binaries
- Improved ARM64 support with SQLite-free builds

### Added
- SQLite-free builds for better platform compatibility
- Build tags for conditional SQLite support

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

[v0.1.2]: https://github.com/ncecere/stripper/releases/tag/v0.1.2
[v0.1.1]: https://github.com/ncecere/stripper/releases/tag/v0.1.1
[v0.1.0]: https://github.com/ncecere/stripper/releases/tag/v0.1.0
