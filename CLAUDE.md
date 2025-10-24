# Claude Code Instructions for dlo

## Project Overview

**dlo** is a lightweight Go CLI client for the German-English dictionary at dict.leo.org. It provides terminal-based dictionary lookups with formatted table output.

- **Language**: Go 1.19+
- **Primary Purpose**: Command-line dictionary lookups (German-English)
- **Total Code**: ~183 lines of Go
- **License**: MIT (Lucas Bremgartner, 2016)

## Architecture

### Key Files

| File | Purpose | Lines |
|------|---------|-------|
| `main.go` | Entry point; CLI argument handling and HTTP requests | 33 |
| `dict.go` | Core logic; XML parsing and output formatting | 75 |
| `dict_test.go` | Unit tests for XML parsing | 75 |
| `go.mod` / `go.sum` | Dependency management | - |
| `README.md` | User documentation with future ideas | - |
| `dict.leo.org.md` | API documentation and query parameters | - |

### Dependencies

- `github.com/TreyBastian/colourize` - Terminal color formatting
- `github.com/olekukonko/tablewriter` - ASCII table rendering
- `github.com/mattn/go-runewidth` - Unicode width calculations (indirect)

### Data Flow

1. User provides search term via CLI argument
2. `main.go` constructs HTTP GET request to dict.leo.org API
3. API returns XML response with translation data
4. `dict.go` parses XML into structured data (Doc → Sections → Entries → Sides)
5. Output formatter creates colored table with search term highlighting
6. Results displayed to stdout organized by section (adjectives, verbs, nouns, etc.)

## Core Data Structures

Located in `dict.go`:

- **Doc**: Root structure containing search term and sections
- **Section**: Groups entries by category (sectionType: "adjectives/adverbs", "verbs", etc.)
- **Entry**: Individual translation entry
- **Side**: Language-specific word data (lang: "en" or "de")

## Development Guidelines

### Building

```bash
go build                 # Produces 'dlo' binary
go run main.go <term>    # Run without building
```

### Testing

```bash
go test                  # Run unit tests
go test -v               # Verbose output
```

Tests are in `dict_test.go` and validate XML parsing with sample responses.

### Code Style

- Standard Go formatting (`gofmt`)
- Simple, functional approach
- Minimal error handling (currently panics on errors)
- No external configuration files

## dict.leo.org API

### Endpoint

```
http://dict.leo.org/dictQuery/m-vocab/ende/query.xml?<params>
```

### Query Parameters (see dict.leo.org.md)

- `tolerMode=nof` - Tolerance mode
- `lp=ende` - Language pair (English-German)
- `lang=de` - Default language
- `rmWords=off` - Keep words in response
- `rmSearch=on` - Remove search markers
- `sectLenMax=10` - Maximum results per section
- `resultOrder=basic` - Result ordering

### Response Format

XML structure with entries containing:
- Search term
- Sections with entries
- Each entry has two sides (English/German)
- Words with optional bold/small text markers

## Common Tasks

### Adding a New Feature

1. Check `README.md` for planned features and ideas
2. For new CLI flags: modify `main.go`
3. For parsing logic: modify `dict.go`
4. For output formatting: modify table/color code in `dict.go`
5. Add tests in `dict_test.go`
6. Update README.md if user-facing

### Improving Error Handling

Current code uses `panic()` for errors. Consider:
- Replacing panics with proper error returns
- Adding user-friendly error messages
- Handling network failures gracefully

### Debugging

To see raw XML response:
1. Print `resp` body before parsing in the `main()` function
2. Check `dict.leo.org.md` for API documentation
3. Test XML parsing with `dict_test.go` samples

## Important Notes

### API Limitations

- No official API documentation from leo.org
- API structure documented in `dict.leo.org.md` based on observation
- API may change without notice
- Consider adding user-agent and rate limiting for production use

### Terminal Compatibility

- Uses ANSI color codes via `colourize` package
- Table rendering requires monospace font
- May not render correctly in all terminals

## Testing Strategy

### Current Coverage

- XML parsing (`ProcessQueryXml()` function)
- Basic structure validation
- Test data in `dict_test.go` with sample XML

### Gaps

- No tests for HTTP request logic
- No tests for output formatting
- No integration tests with live API
- No error case testing

When adding tests, consider:
- Mocking HTTP responses
- Testing various XML response structures
- Validating color/formatting output
- Testing edge cases (empty results, special characters, etc.)

## Git Workflow

- Main development branch: `master`
- Feature branches use `claude/` prefix
- Clean commit history preferred
- MIT license - ensure all contributions are compatible

## Useful Commands

```bash
# Build and install to $GOPATH/bin
go install

# Run with search term
./dlo <search-term>

# Example
./dlo house

# Run tests with coverage
go test -cover

# Update dependencies
go mod tidy

# Format code
gofmt -w .
```

## Contact & Maintenance

- Author: Lucas Bremgartner
- Repository: github.com/breml/dlo
- GitHub issue tracker available for bug reports and feature requests
- Consider checking README for contribution guidelines

---

**Last Updated**: 2025-10-24
**For**: Claude Code assistance with dlo codebase
