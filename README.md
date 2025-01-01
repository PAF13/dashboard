# CLI Dashboard App

## Overview

The CLI Dashboard App is a powerful command-line interface (CLI) tool built with Golang. It leverages the `bubbletea` package for creating an interactive terminal user interface and integrates with `ripgrep` to efficiently search for patterns within files. The app provides JSON-formatted search results, making it suitable for both human and machine consumption.

## Features

- **Ripgrep Integration**: Utilizes `ripgrep` to perform high-speed, recursive search across directories.
- **JSON Parsing**: Processes `ripgrep`'s JSON output to provide structured and detailed information about search matches.
- **Interactive CLI**: Built with the `bubbletea` package to create an interactive and user-friendly command-line dashboard.
- **Detailed Match Information**: Outputs match details, including line numbers, offsets, and submatches.
- **Cross-platform Support**: Compatible with major operating systems, including Windows, macOS, and Linux.

## Installation

1. **Clone the Repository**:
   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. **Install Dependencies**:
   - Ensure you have [Golang](https://go.dev/dl/) installed.
   - Install `ripgrep`:
     - macOS: `brew install ripgrep`
     - Linux: Use your package manager (e.g., `apt-get install ripgrep` for Ubuntu).
     - Windows: Download from the [ripgrep releases page](https://github.com/BurntSushi/ripgrep/releases).

3. **Build the Application**:
   ```bash
   go build -o cli-dashboard
   ```

## Usage

1. **Run the App**:
   ```bash
   ./cli-dashboard
   ```

2. **Perform a Search**:
   The application requires two inputs:
   - `rootDir`: The directory where the search should start.
   - `text`: The pattern or text to search for.

3. **Example Output**:
   The application displays search results in the following format:
   ```json
   {
       "type": "match",
       "data": {
           "path": {
               "text": "example.txt"
           },
           "lines": {
               "text": "This is a matching line."
           },
           "line_number": 42,
           "absolute_offset": 128,
           "submatches": [
               {
                   "match": {
                       "text": "matching"
                   },
                   "start": 10,
                   "end": 18
               }
           ]
       }
   }
   ```

## Code Overview

### Main Functionality

The core functionality of the app is defined in the `rg` function. It:

- Executes `ripgrep` with the provided search text and root directory.
- Parses the JSON output line-by-line using `bufio.Scanner`.
- Extracts and structures match details into a `RipgrepMatch` type.

### Error Handling

- Handles errors during command execution and JSON parsing gracefully.
- Logs issues to the console to aid debugging.

## Requirements

- Go 1.18+
- Ripgrep

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your improvements.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

## Acknowledgments

- [BurntSushi/ripgrep](https://github.com/BurntSushi/ripgrep) for the amazing search tool.
- [Bubbletea](https://github.com/charmbracelet/bubbletea) for making terminal UIs a delight to build.

---

Feel free to enhance this README to better suit your app's unique features and functionalities!

