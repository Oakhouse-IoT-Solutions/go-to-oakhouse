# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.18.0] - 2024-12-19

### Added
- **Wire Dependency Injection**: Enhanced Wire implementation with comprehensive documentation and detailed comments
- **Code Quality Improvements**: Added detailed comments throughout Wire dependency injection points
- **Documentation Enhancement**: Added detailed Wire implementation guide covering Provider Pattern, Constructor Injection, and Advanced Wire Patterns
- **Template Fixes**: Fixed missing adapter import in CLI-generated projects

### Changed
- **Version Update**: Updated to version 1.18.0 across all files and documentation
- **Wire Templates**: Enhanced Wire templates with detailed explanatory comments
- **CLI Templates**: Improved generated project templates with proper imports

### Fixed
- **Import Issues**: Resolved missing adapter import in newly generated projects
- **Template Generation**: Fixed CLI template generation to include all necessary imports
- **Wire Configuration**: Improved Wire setup and configuration in generated projects

### Improved
- **Developer Experience**: Better code clarity and onboarding through comprehensive Wire documentation
- **Code Maintainability**: Enhanced code structure with detailed dependency injection patterns
- **Testing Guidance**: Added Wire testing patterns and best practices

## [1.16.0] - 2024-12-19

### Added
- **Version Release**: Updated to version 1.16.0 with comprehensive version consistency across all files
- **Documentation Updates**: Synchronized all version references across README, DOCUMENTATION, PHILOSOPHY, and template files
- **Template Improvements**: Updated CLI templates to reflect current version information

### Changed
- **Version Consistency**: Updated all version references from various previous versions to 1.16.0
- **Installation Commands**: Updated installation commands across all documentation to use v1.16.0
- **Version Badges**: Updated version badges in HTML templates and documentation

### Fixed
- **Version Synchronization**: Resolved inconsistencies in version numbers across different files
- **Documentation Accuracy**: Ensured all documentation reflects the current version

### Improved
- **Release Process**: Streamlined version update process across all project files
- **User Experience**: Consistent version information provides clearer guidance for users

## [1.14.0] - 2024-12-19

### Added
- **Code Quality Fixes**: Fixed typos and improved code consistency throughout the framework
- **Documentation Improvements**: Enhanced documentation accuracy and completeness for better user experience
- **Build Optimizations**: Streamlined build process for better development experience and faster compilation
- **Developer Experience**: Continued improvements to CLI usability, feedback, and error handling

### Changed
- **Project Structure**: Further refinements to project organization and template maintainability
- **Documentation Structure**: Updated all documentation files to reflect version 1.14.0
- **Installation Instructions**: Updated installation commands across all documentation

### Fixed
- **Typo Corrections**: Fixed various typos in code comments and descriptions
- **Code Consistency**: Improved consistency in code formatting and structure

### Improved
- **Build Process**: Optimized build workflow for faster compilation and deployment
- **Documentation Quality**: Enhanced accuracy and completeness of all documentation
- **Developer Tools**: Better CLI experience with improved feedback and error handling

## [1.13.0] - 2024-12-19

### Added
- **Enhanced Code Quality**: Further improvements to error handling and code organization throughout the framework
- **Updated Documentation**: Comprehensive documentation updates reflecting latest features and best practices
- **Performance Enhancements**: Additional optimizations for faster development workflow and project scaffolding
- **Improved Developer Tools**: Enhanced CLI experience with better feedback, guidance, and error messages

### Changed
- **Architecture Refinements**: Continued improvements to project structure and template organization for maintainability
- **Documentation Structure**: Updated all documentation files to reflect version 1.13.0
- **Installation Instructions**: Updated installation commands across all documentation

### Improved
- **Developer Experience**: Streamlined workflow with enhanced CLI interactions and better error handling
- **Code Organization**: Further refinements to project structure for better maintainability
- **Performance**: Additional optimizations for faster project generation and development workflow
- **Documentation Quality**: Comprehensive updates to ensure accuracy and completeness

## [1.12.0] - 2024-12-19

### Added
- **Enhanced Error Handling**: Improved error messages and validation throughout the CLI tool
- **Code Documentation**: Added comprehensive inline documentation and code comments
- **Performance Monitoring**: Streamlined code generation processes for better performance
- **Developer Guidance**: Enhanced CLI feedback with better user guidance and help text

### Changed
- **Code Organization**: Refined project structure and template organization for better maintainability
- **CLI Feedback**: Improved command output formatting and user experience
- **Template Structure**: Enhanced generated code templates with better organization
- **Documentation**: Updated all documentation to reflect latest best practices

### Improved
- **Code Quality**: Enhanced overall code quality with better error handling patterns
- **Developer Experience**: Streamlined development workflow with improved CLI interactions
- **Performance**: Optimized code generation algorithms for faster project scaffolding
- **Maintainability**: Better code organization and structure for long-term maintenance

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