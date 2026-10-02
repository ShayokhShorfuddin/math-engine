# Math Engine

A small mathematical expression evaluation engine written in Go. It evaluates arithmetic expressions with operator precedence and includes an interactive REPL. The primary goal of writing this engine was to gain a better understanding of Go and Cobra CLI, not to build a feature-rich engine.

## Requirements

- Go 1.27.1 or later

## Usage

Run an expression directly:

```bash
go run . eval "2 + 3 * 4"
# 14
```

Start the interactive calculator:

```bash
go run . repl
```

Enter expressions at the prompt and type `quit` to exit. Check the version with:

```bash
go run . version
```

The calculator supports decimal numbers and these operators:

- Addition: `+`
- Subtraction: `-`
- Multiplication: `*`
- Division: `/`

Multiplication and division are evaluated before addition and subtraction. 

## Development

Run the test suite with:

```bash
go test ./...
```
