# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.11.0] - 2024-12-19

### Added
- **Enhanced Branding**: Added signature "🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡" throughout the CLI tool
- **Professional Attribution**: Integrated creator signature in all generated code templates
- **Consistent Branding**: Added signature to all CLI command success messages
- **Visual Enhancements**: Updated all command outputs with modern emojis and styling

### Changed
- **CLI Interface**: Enhanced root command description with signature and feature list
- **Success Messages**: Updated all generate command outputs to include branding signature
- **User Experience**: Improved visual consistency across all CLI interactions
- **Template Branding**: All generated code now includes professional attribution

### Improved
- **Brand Recognition**: Consistent "Oakhouse" branding across all tool outputs
- **Professional Presentation**: Enhanced CLI tool appearance with signature styling
- **User Engagement**: More engaging and branded command-line experience

## [1.9.0] - 2024-12-19

### Added
- **Simplified database setup**: Streamlined database connection setup replacing complex migration system
- Enhanced database management with single-command setup
- Improved error handling and user feedback for database operations

### Removed
- **Migration system**: Removed `oakhouse migrate` commands and related functionality
- Deleted migration file management and auto-migration features
- Removed complex migration CLI commands (`migrate up`, `migrate create`, etc.)

### Changed
- **Simplified database workflow**: Projects now use `oakhouse add database` for database setup
- Updated all documentation to reflect new database management approach
- Modified project templates to use simplified database setup

### Fixed
- Resolved CLI compilation issues with missing function definitions
- Fixed import dependencies and module structure
- Corrected command registration for all CLI tools

### Improved
- Cleaner architecture without migration complexity
- Better user experience with single database setup command
- Enhanced documentation consistency across all files

## [1.8.0] - 2024-12-19

### Added
- **Enhanced .gitignore**: Improved Git ignore patterns for better project management
  - Added exclusion for generated code directories (handler/, model/, repository/, route/, service/)
  - Added explicit exclusion for testproject/ directory
  - Better organization of ignored files and directories

### Changed
- Updated CLI version to 1.8.0
- Enhanced documentation with new version references

### Improved
- Better Git workflow with comprehensive ignore patterns
- Cleaner repository structure by excluding generated code

## [1.7.0] - 2024-12-19

### Changed
- Updated CLI version to 1.7.0
- Enhanced documentation and version consistency

### Improved
- Better version management across all documentation files
- Streamlined release process

## [1.6.0] - 2024-12-19

### Added
- **Static File Serving**: Enhanced project templates with static file serving capabilities
  - Automatic static file server setup in generated projects
  - Built-in `index.html` template with modern design
  - Seamless integration with Fiber's static middleware
- **Enhanced Project Templates**: Improved project scaffolding
  - Projects now include a `static` directory with `index.html`
  - Updated `app_server.go` template to serve static files
  - Better default project structure for web applications

### Changed
- Updated CLI version to 1.6.0
- Enhanced project generation with static file support
- Improved route templates for better static content handling

### Fixed
- Resolved route generation issues in project templates
- Fixed static file serving configuration in generated projects
- Improved template consistency across generated files

## [1.5.0] - 2024-12-19

### Added
- **Home Page**: Beautiful landing page showcasing framework features
  - Modern responsive design with gradient background
  - Feature highlights and getting started guide
  - Author information and project details
- **Enhanced Documentation**: Improved project presentation

### Changed
- Updated version to 1.5.0
- Enhanced project branding and visual identity

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