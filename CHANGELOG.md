# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.4.0] - 2024-12-19

### Added
- **Simplified Handlers**: New lightweight handler generation for rapid prototyping
  - Handlers now generate with simple text responses using `c.SendString()`
  - No database dependencies required for initial testing
  - Perfect for API structure prototyping and endpoint testing
- **Improved Route Generation**: Fixed import path issues in generated route files
- **Enhanced CLI**: Better module name detection from `go.mod` files

### Changed
- Handler templates now generate simplified implementations by default
- Route templates no longer require database adapter parameters
- Improved error handling in route generation

### Fixed
- Fixed malformed import paths in generated route files
- Fixed `ProjectName` template variable resolution
- Improved module name detection for correct import paths

### Technical Details
- Updated `resourceHandlerTemplate` to generate simplified handlers
- Updated `resourceRouteTemplate` to remove database dependencies
- Added `getModuleName()` function for proper import path generation
- Enhanced template data passing for consistent code generation

## [1.3.0] - Previous Release

### Features
- Clean architecture pattern implementation
- GORM integration with scoping support
- Fiber framework integration
- Docker containerization support
- Comprehensive CLI tool for scaffolding