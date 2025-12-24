# Repository Guidelines

## Overview

This repository hosts Advent of Code solutions in Go, organized by day and
year. Each day's solution is implemented as a separate command in the `cmd/`
directory, with shared utilities in `utils/`. Puzzle inputs are stored in the
`data/` directory. The goal of the exercise is fun and learning about language features, so solutions may vary in style and complexity. Agents (that's you) should bias for explaining code clearly and being very clear about rationale.

## Project Structure & Module Organization
- `aoc.go` and `cmd/` define the Cobra CLI (`aoc`) and per-day commands like `day2101`.
- `utils/` contains shared helpers (file loading, parsing, etc.).
- `data/` holds puzzle inputs and samples referenced by `cmd/*.go` (e.g., `data/2101.txt`).
- `2021/` and `2025/` contain older standalone solutions and notes for specific years.
- `scripts/` includes utility scripts such as `new_day.sh` for scaffolding a new day.

## Build, Test, and Development Commands
- `go run .` shows CLI help.
- `go run . day2101` runs a specific day command (use any `dayYYDD` in `cmd/`).
- `go build -o aoc` builds the CLI binary.
- `scripts/new_day.sh YYDD` scaffolds a new `cmd/YYDD.go` command and input hook.

## Coding Style & Naming Conventions
- Use standard Go formatting (`gofmt`); tabs for indentation.
- Keep day commands named `dayYYDD` and file names `cmd/YYDD.go`.
- Input files live in `data/YYDD.txt`; sample data can be `data/YYDD-sample.txt`.
- Prefer small helpers in `utils/` over duplicating parsing logic in each day.

## Testing Guidelines
- There are currently no Go test files (`*_test.go`).
- If adding tests, use `go test ./...` and keep test names aligned with day or helper (e.g., `TestDay2101`).

## Commit & Pull Request Guidelines
- Commit history uses short, informal messages tied to day progress (e.g., “2502 phase2 working!”).
- Keep messages brief, mention the day (`YYDD`) and status.
- PRs should describe the day/part implemented, inputs touched in `data/`, and any notable runtime considerations.

## Agent Notes
- Keep new day scaffolds consistent with `scripts/new_day.sh` output.
- Avoid large refactors; this repo favors incremental, day-by-day additions.
