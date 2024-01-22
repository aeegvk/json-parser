# JSON Parser

This is a simple JSON parser written in Go. It reads a JSON file, prettifies the JSON, and writes the result to another file. This was built to parse sensitive json logs/files.

TODO: Add functionality to also parse invalid files.

## Requirements

- Go 1.21.6 or later

## Usage

1. Build the project:

```bash
go build main.go
```

2. Run the program with a filename as a command-line argument:

```bash
./main input.json
```

This will read the JSON from `input.json`, prettify it, and write the result to `output.json`.

## Note

This program uses the `json` package from the Go standard library, which strictly follows the JSON specification. If the input JSON is invalid, the program will return an error.