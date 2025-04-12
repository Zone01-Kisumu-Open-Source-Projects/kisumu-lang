# Code Structure

The Kisumu programming language project follows a well-organized directory structure to ensure maintainability, scalability, and clarity. Below is a detailed explanation of the structure and the purpose of each directory and file.

## Project Structure

``` bash
> tree .
.
├── LICENSE
├── README.md
├── cmd
│   ├── kisumu
│   │   └── main.go
│   └── repl
│       └── main.go
├── docs
│   ├── README.md
│   ├── development
│   │   └── contribution-guidelines.md
│   ├── specs
│   │   ├── architecture.md
│   │   ├── features.md
│   │   ├── introduction.md
│   │   └── roadmap.md
│   └── technical
│       ├── data-structures.md
│       └── lexer.md
├── go.mod
├── internal
│   ├── interpreter
│   │   ├── interpreter.go
│   │   └── interpreter_test.go
│   ├── lexer
│   │   ├── lexer.go
│   │   └── lexer_test.go
│   ├── parser
│   │   ├── parser.go
│   │   └── parser_test.go
│   ├── repl
│   │   ├── file_reader.go
│   │   └── repl.go
│   └── token
│       └── tokens.go
├── pkg
│   └── common
│       └── errors.go
└── tests
    └── lexer
        └── token_tests.ksm
```

## Directory and File Descriptions

### Root Level
- **`LICENSE`**: Contains the licensing information for the project.
- **`README.md`**: Provides an overview of the project, including its purpose, usage, and key features.
- **`go.mod`**: Manages dependencies and module information for the Go project.

### `cmd/`
Contains the entry points for the Kisumu project.
- **`kisumu/main.go`**: The main entry point for the Kisumu compiler.
- **`repl/main.go`**: The main entry point for the REPL (Read-Eval-Print Loop).

### `docs/`
Houses all documentation related to the project.
- **`README.md`**: Overview of the documentation structure.
- **`development/contribution-guidelines.md`**: Guidelines for contributors.
- **`specs/`**: Contains language specifications:
  - **`architecture.md`**: Describes the system's architecture.
  - **`features.md`**: Lists the language features.
  - **`introduction.md`**: Introduction to Kisumu.
  - **`roadmap.md`**: Outlines the development timeline.
- **`technical/`**: Technical documentation:
  - **`data-structures.md`**: Details on data structures in Kisumu.
  - **`lexer.md`**: Explanation of the lexical analysis component.

### `internal/`
Contains the core components of the language implementation. This directory is not intended for public APIs.
- **`interpreter/`**: Implements the interpreter logic.
  - **`interpreter.go`**: Core interpreter functionality.
  - **`interpreter_test.go`**: Tests for the interpreter.
- **`lexer/`**: Handles lexical analysis.
  - **`lexer.go`**: Lexer implementation.
  - **`lexer_test.go`**: Tests for the lexer.
- **`parser/`**: Responsible for parsing.
  - **`parser.go`**: Parser implementation.
  - **`parser_test.go`**: Tests for the parser.
- **`repl/`**: Implements the REPL functionality.
  - **`file_reader.go`**: Reads files in the REPL environment.
  - **`repl.go`**: Core REPL logic.
- **`token/`**: Defines tokens used in lexical analysis.
  - **`tokens.go`**: Token definitions.

### `pkg/`
Contains reusable packages.
- **`common/errors.go`**: Defines common error structures and utilities.

### `tests/`
Contains test files for the project.
- **`lexer/token_tests.ksm`**: Test cases for token-related functionality.
