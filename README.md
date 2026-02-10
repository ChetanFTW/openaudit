# OpenAudit

OpenAudit is a lightweight, Go-based repository auditing tool that helps developers and open-source projects evaluate code quality, CI setup consistency, and repository best practices — all from the command line.

## Features

```bash
 Detects if the project has a Go module 
 Checks for CI configuration (GitHub Actions, GitLab CI, etc.)
 Detects common repository files like README.md
 Provides a quick summary report
 Easy to extend with plugins or extra analyzers
 Ready to include in CI/CD
```

## Why OpenAudit?

Modern open-source projects rely on consistent tooling and quality checks. OpenAudit helps you assess a repo’s hygiene instantly — ideal for maintainers, contributors, and learners alike.

### It’s:

Fast — Written in Go

Simple — Zero dependencies required to run

Extensible — Add more checks over time

Portfolio-ready — Shows architecture, CI, and tooling skills

## Quick Start
#### Clone the Project

```bash
git clone https://github.com/yourusername/openaudit.git
cd openaudit

```
#### Run the Scanner
```bash
go run ./cmd/openaudit

```
Scan a specific path:
```bash
go run ./cmd/openaudit /path/to/repo

```
## Testing

Add tests as you expand the project:
```bash
go test ./...

```

### License

This project is open source — feel free to choose an appropriate license such as MIT or Apache 2.0.
