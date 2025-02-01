## **Kisumu Programming Language**

[![](https://github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/actions/workflows/dev.yml/badge.svg?workflow:ci)](https://github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/actions/workflows/dev.yml?query=workflow:ci "Builds")
[![](https://img.shields.io/discord/322538954119184384.svg?logo=discord&logoColor=white&label=Discord&color=5865F2)](https://discord.gg/amrst3npC8 "Join the Discord chat at https://discord.gg/amrst3npC8")
[![](https://img.shields.io/github/license/Zone01-Kisumu-Open-Source-Projects/kisumu-lang)](https://opensource.org/license/bsd-3-clause "License: BSD 3-Clause")
[![](https://www.codetriage.com/zone01-kisumu-open-source-projects/kisumu-lang/badges/users.svg)](https://www.codetriage.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang "Help Contribute to Open Source")
[![](https://pkg.go.dev/badge/github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang?status.svg)](https://pkg.go.dev/github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang "GoDoc")

Welcome to **Kisumu**, a modern statically-typed programming language inspired by Go, Rust, and Python. Kisumu is designed for simplicity, performance, and scalability, making it an excellent choice for both beginner programmers and experienced developers building robust, efficient applications.

### **Key Features**

- **C-style Syntax**: Easy-to-read and familiar syntax for developers.
- **Statically Typed**: Ensures type safety and performance.
- **First-Class Concurrency**: Powerful models like goroutines, actors, and channels.
- **Extensibility**: Modular structure with packages for scalable codebases.
- **Garbage Collection**: Automatic memory management for safety and ease.
- **Interoperability**: FFI support for integration with C, Go, and more.

### **Getting Started**

#### **1. Install Kisumu**

Installation instructions will be provided closer to the public release. Stay tuned for updates.

#### **2. Write Your First Program**

```ksm
fn main() {
    print("Hello, Kisumu!")
}
```

#### **3. Explore the Documentation**

Dive deeper into Kisumu's features and capabilities:

- [Introduction](docs/README.md)
- [Design Philosophy](docs/specs/architecture.md)
- [Quick Start Guide](docs/development/setup.md)

<!-- #### **4. Try the Examples**

Explore working examples:

- [Hello World](examples/hello_world.ksm)
- [Concurrency](examples/concurrency_example.ksm)
- [Modules](examples/modules_example.ksm) -->

### **Development and Testing**

#### **Running Unit Tests**

To ensure code quality and maintain reliability, Kisumu uses Go's built-in testing framework. Here's how to run the tests locally:

1. **Run All Tests**
   ```bash
   make test
   ```
   This command executes all unit tests in the project.

2. **Run Tests with Coverage**
   ```bash
   make coverage
   ```
   This generates a coverage report and displays the coverage statistics for each package.

3. **Run Tests for Specific Package**
   ```bash
   go test ./path/to/package
   ```
   Replace `path/to/package` with the package you want to test.

4. **Run Tests with Verbose Output**
   ```bash
   go test -v ./...
   ```
   This shows detailed test output including individual test cases.

For development, we recommend running the complete test suite before submitting any changes:
```bash
make pre-commit
```
This command runs tests along with code formatting and linting checks.

### **Contributing to Kisumu**

We welcome contributions to make Kisumu even better!  
Learn how to get involved by reading our [Contribution Guidelines](docs/development/contribution-guidelines.md).

### **Community and Feedback**

Join the growing Kisumu community:

- Share feedback in [docs/community/feedback.md](docs/community/communication.md).
- Check out our [Code Of Conduct](docs/community/code-of-conduct.md).

### **License**

This project is licensed under the BSD 3-Clause License. See the [LICENSE](/LICENSE) file for details.

## :clap: Supporters

[![Stargazers repo roster for @Zone01-Kisumu-Open-Source-Projects/kisumu-lang](https://reporoster.com/stars/dark/Zone01-Kisumu-Open-Source-Projects/kisumu-lang)](https://github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/stargazers)

[![Forkers repo roster for @Zone01-Kisumu-Open-Source-Projects/kisumu-lang](https://reporoster.com/forks/dark/Zone01-Kisumu-Open-Source-Projects/kisumu-lang)](https://github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/network/members)

<p align="center"><a href="#"><img src="https://img.shields.io/badge/Back%20to%20top--lightgrey?style=social" alt="Back to top" height="20"/></a></p>
