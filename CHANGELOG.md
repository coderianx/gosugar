# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/)
and this project adheres to [Semantic Versioning](https://semver.org/).

## [v0.5.0] - 2026-02-09
### Added
- **String Manipulation Module** (`strings.go`)
  - Basic operations: `Trim()`, `ToUpper()`, `ToLower()`, `Reverse()`
  - Searching & checking: `Contains()`, `HasPrefix()`, `HasSuffix()`
  - Replacement: `Replace()`, `ReplaceAll()`
  - Splitting & joining: `Split()`, `Join()`
  - Repetition & padding: `Repeat()`, `PadLeft()`, `PadRight()`
  - Truncation & formatting: `Truncate()`, `Capitalize()`, `Decapitalize()`
  - Case conversions: `CamelCase()`, `SnakeCase()`, `KebabCase()`, `PascalCase()`
  - URL slug generation: `Slugify()`
- **String Documentation**
  - English API reference (`docs/api/strings.md`)
  - Turkish API reference (`docs-tr/api/strings.md`)
  - 22 string helper functions with examples

## [v0.4.0] - 2026-02-8
### Added
- add HTTP PUT and DELETE methods
- PUT helper functions
- DELETE helper functions


## [v0.3.0] - 2026-02-06
### Added
- PostJSON helper function
- MustPostBody helper function
- MustPostHeader helper function

## [v0.2.0] - 2026-02-04
### Added
- GetJSON generic helper for decoding JSON responses
- MustGetBody helper function
- MustGetHeader helper function


## [v0.1.0] - 2026-02-03
### Added
- Initial release
- Basic project structure
- Core functionality
