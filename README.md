# CSV to JSON Converter (Go CLI Tool)

This command-line tool converts a CSV file into a JSON file using Go. It supports custom separators (comma or semicolon) and optional pretty-printing of the JSON output.

---

## Features

- Read CSV data with a user-specified separator (`comma` or `semicolon`)
- Convert CSV rows to a list of JSON objects
- Supports pretty-printed JSON output
- Handles common file I/O and parsing errors
- Graceful exit on malformed rows or unsupported formats

---

## Requirements

- Go 1.20 or later
- A valid CSV file

---

## Installation

Clone the repository and build the binary:



git clone https://github.com/RewanshChoudhary/go-bench.git
cd go-bench
go build -o main
