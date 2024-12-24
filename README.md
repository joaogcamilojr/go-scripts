# Golang Scripts

This repository contains a collection of scripts written in Go (Golang) for automation, data processing, and utility tasks.

## Table of Contents

- [About](#about)
- [Requirements](#requirements)
- [Installation](#installation)
- [Running Scripts](#running-scripts)
- [License](#license)

## About

A variety of scripts designed to simplify everyday tasks, including database handling, file operations, and API integrations.

## Requirements

1. **Golang** (recommended: Go 1.23+).
2. **Version Manager** (optional but recommended): [asdf](https://asdf-vm.com/).

## Installation

### Install Go

- Official Download: [https://golang.org/dl/](https://golang.org/dl/)
- Or use **asdf** for version management:
  ```bash
  asdf plugin add golang https://github.com/kennyp/asdf-golang.git
  asdf install golang latest
  asdf global golang latest
  ```

## Running Scripts

1. Clone the repository:
   ```bash
   git clone https://github.com/joaogcamilojr/go-scripts
   cd go-scripts
   ```
2. Execute a script:
   ```bash
   go run scripts/script-name.go
   ```
3. Build and run:
   ```bash
   go build -o bin/script-name scripts/script-name.go
   ./bin/script-name
   ```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
