# fmon

`fmon` is a command-line tool to monitor file changes in a directory and automatically rerun a specified command when changes are detected. It's designed to improve development workflows by automating the restart of processes, such as application servers, when code is modified.

## Features

- **Directory Monitoring**: Watches a specified directory (defaults to the current directory) for file changes.
- **Automatic Command Restart**: Restarts a specified command whenever a change is detected.
- **Customizable Watch Directory**: Use a flag to set a custom directory to monitor, accepting both relative and absolute paths.
- **Colorful Console Output** (WIP): Displays monitored commands, directories, and file changes in a readable neon-green format.

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/fmon.git
cd fmon
```

2. Build the application:

```bash
go build -o fmon
```

3. Run the binary:

```bash
./fmon
```

## Usage

The basic syntax of the `fmon` command is:

```bash
fmon run "<command>" [flags]
```

### Examples

1. Monitor Current Directory:

To monitor the current directory and rerun a command (e.g., `go run main.go`) whenever a file changes:

```bash
fmon run "go run main.go"
```

2. Monitor a Custom Directory:

To specify a different directory to monitor (relative or absolute path), use the `--dir` or `-d` flag:

```bash
fmon run "go run main.go" --dir "./path/to/dir"
```

## Flags

- `--dir`, `-d`: Specify a directory to watch. Defaults to the current directory.

## Example Output

```bash
Command: go run main.go
Watching directory: /path/to/project

File modified: /path/to/project/main.go
```

## Development

This project uses [watcher](https://github.com/radovskyb/watcher) to monitor file changes and [Cobra](https://github.com/spf13/cobra) for CLI functionality. To modify or extend `fmon`, explore the main files and directories:

- **cmd/**: Contains CLI commands and flags.
- **watcher/**: Manages file monitoring and detects changes.
- **process/**: Handles running and restarting commands.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any features, bug fixes, or improvements.
