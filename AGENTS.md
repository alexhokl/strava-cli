# AGENTS.md - Strava CLI

This document provides guidelines for AI coding agents working on this Go CLI application
that interacts with the Strava API.

## Project Overview

- **Language**: Go 1.25+
- **Module**: `github.com/alexhokl/strava-cli`
- **CLI Framework**: [Cobra](https://github.com/spf13/cobra) + [Viper](https://github.com/spf13/viper)
- **Task Runner**: [Task](https://taskfile.dev/) (Taskfile.yml)
- **API Client**: Auto-generated from OpenAPI spec (swagger directory)

## Build Commands

```bash
# Build (verify compilation)
task build

# Install locally
task install

# Standard Go commands also work
go build
go install
```

## Test Commands

```bash
# Run all tests
task test
go test ./...

# Run tests with coverage
task coverage
go test --cover ./...

# Run a single test by name
go test -run TestName ./path/to/package

# Run a specific subtest
go test -run TestName/SubTestName ./path/to/package

# Run tests in a specific package
go test ./cmd/...

# Coverage with HTML report
task coverage-html
task open-coverage-html

# Benchmarks
task bench
go test -bench=. -benchmem ./...
```

**Note**: Most tests in `swagger/test/` are auto-generated and skipped by default.

## Lint & Security Commands

```bash
# Run linter (golangci-lint)
task lint
golangci-lint run

# Security scan
task sec
gosec ./...
```

## Code Generation

```bash
# Pull latest Strava API spec
task pull-swagger

# Regenerate API client from swagger.json
task generate-swagger
```

**Important**: Never manually edit files in the `swagger/` directory - they are auto-generated.

## Project Structure

```
strava-cli/
├── main.go              # Entry point, calls cmd.Execute()
├── cmd/                 # CLI commands (Cobra-based)
│   ├── root.go          # Root command, config, token validation
│   ├── login.go         # OAuth login command
│   ├── list.go          # Parent command for list subcommands
│   ├── list_activity.go # List activities subcommand
│   ├── list_segment.go  # List segments subcommand
│   ├── show.go          # Parent command for show subcommands
│   ├── edit.go          # Parent command for edit subcommands
│   └── update.go        # Parent command for update subcommands
├── ui/                  # Terminal UI (Bubble Tea models)
├── swagger/             # Auto-generated API client (DO NOT EDIT)
├── Taskfile.yml         # Task runner configuration
└── swagger.json         # Strava API specification
```

## Code Style Guidelines

### Import Organization

Group imports in this order with blank lines between groups:
1. Standard library
2. External packages
3. Internal packages

```go
import (
    "context"
    "fmt"
    "os"

    "github.com/alexhokl/helper/authhelper"
    "github.com/olekukonko/tablewriter"
    "github.com/spf13/cobra"

    "github.com/alexhokl/strava-cli/swagger"
    "github.com/alexhokl/strava-cli/ui"
)
```

### Naming Conventions

| Element | Convention | Example |
|---------|------------|---------|
| Files | snake_case | `list_activity.go`, `editor_model.go` |
| Packages | lowercase, single word | `cmd`, `ui`, `swagger` |
| Command variables | camelCase + `Cmd` suffix | `listCmd`, `loginCmd` |
| Options structs | `<command>Options` | `editActivityOptions` |
| Options variables | `<command>Opts` | `editActivityOpts` |
| Run functions | `run<Command>` | `runLogin`, `runListActivities` |
| Exported types | PascalCase | `EditorModel`, `EditorKeys` |
| Unexported fields | camelCase | `nameInput`, `hasUpdate` |

### Error Handling

- Use `RunE` (returns error) instead of `Run` for Cobra commands
- Return errors rather than panicking
- Wrap errors with context using `fmt.Errorf`
- Use `_ = someFunc()` only for non-critical operations (e.g., saving config)

```go
func runSomeCommand(_ *cobra.Command, _ []string) error {
    result, err := someOperation()
    if err != nil {
        return fmt.Errorf("failed to do something: %v", err)
    }
    return nil
}
```

### Command Structure Pattern

Each command file follows this structure:

```go
package cmd

// Options struct (if flags needed)
type editActivityOptions struct {
    id int64
}

var editActivityOpts editActivityOptions

// Command definition
var editActivityCmd = &cobra.Command{
    Use:   "activity",
    Short: "Edit an activity",
    RunE:  runEditActivities,
}

// Register with parent in init()
func init() {
    editCmd.AddCommand(editActivityCmd)
    flags := editActivityCmd.Flags()
    flags.Int64Var(&editActivityOpts.id, "id", 0, "Activity ID")
    _ = editActivityCmd.MarkFlagRequired("id")
}

// Run function
func runEditActivities(_ *cobra.Command, _ []string) error {
    // Implementation
    return nil
}
```

### API Client Pattern

Standard pattern for authenticated API calls:

```go
savedToken, err := authhelper.LoadTokenFromViper()
if err != nil {
    return err
}
tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: savedToken.AccessToken})
auth := context.WithValue(context.Background(), swagger.ContextOAuth2, tokenSource)
config := swagger.NewConfiguration()
client := swagger.NewAPIClient(config)

// Use client.SomeAPI.SomeMethod(auth)...
```

### Table Output Pattern

Use tablewriter for CLI table output:

```go
table := tablewriter.NewWriter(os.Stdout)
table.SetHeader([]string{"ID", "Name", "Date"})
table.SetBorder(false)
table.AppendBulk(data)
table.Render()
```

### TUI Pattern (Bubble Tea)

Models implement the standard Bubble Tea interface:

```go
func (m Model) Init() tea.Cmd { return nil }
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) { ... }
func (m Model) View() string { ... }
```

## Key Dependencies

- `github.com/alexhokl/helper` - Auth, CLI, JSON utilities
- `github.com/spf13/cobra` - CLI framework
- `github.com/spf13/viper` - Configuration management
- `github.com/charmbracelet/bubbletea` - Terminal UI
- `github.com/charmbracelet/bubbles` - UI components
- `github.com/olekukonko/tablewriter` - Table output
- `golang.org/x/oauth2` - OAuth2 authentication
- `github.com/stretchr/testify` - Test assertions (assert/require)

## Configuration

The CLI uses Viper for configuration stored in `~/.strava-cli.yaml`. Required config:
- `clientId` - Strava OAuth client ID
- `clientSecret` - Strava OAuth client secret
- OAuth tokens are stored after `login` command
