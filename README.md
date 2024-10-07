# Wayback Machine Snapshot Fetcher

## Overview

The **Waybackfetchr** is a command-line tool written in Go that allows users to easily retrieve archived snapshots of web pages from the Internet Archive’s Wayback Machine. With this tool, users can fetch all available snapshot URLs for a given webpage or a list of web pages, enabling easy access to historical versions of content.

## Features

- **Single URL Fetching**: Quickly retrieve all snapshot URLs for a specific web page using the `-u` flag.
- **Batch Processing**: Process multiple URLs by providing a file containing a list of URLs with the `-l` flag.
- **Output Options**: Save the retrieved snapshot URLs to a specified output file using the `-o` flag while also printing the results to the console.
- **Silent Mode**: Enable a clean output experience with the `--silent` flag to display only the results without additional console messages.
- **Colorful Banner**: The tool features a vibrant ASCII art banner to enhance the command-line interface experience.
- **Input from Standard Input**: Supports reading URLs directly from standard input, allowing for flexible usage in scripts or pipelines.
- **Help Command**: Provides a helpful usage guide and flag descriptions when the `-h` or `--help` flags are used.

## Installation

1. Ensure you have Go installed on your machine.
2. Install the tool using the following command:
   ```bash
   go install github.com/KathanP19/waybackfetch@latest
   ```
OR
   ```bash
   git clone https://github.com/KathanP19/waybackfetch.git
   cd waybackfetch
   go install
   ```
   
## Usage
```
# From STDIN
echo "https://vulnweb.com" | waybackfetch

# Fetch snapshots for a single URL
waybackfetch -u <URL>

# Fetch snapshots for a list of URLs from a file
waybackfetch -l <filename>

# Save results to a file while printing to console
waybackfetch -u <URL> -o <outputfile>

# Enable silent mode
waybackfetch -u <URL> --silent

# Display help message
waybackfetch -h
```

## Todo
- [ ] Add Duplicate content check.

## Contributing
Contributions are welcome! Please feel free to submit issues or pull requests to enhance the functionality and performance of the tool.
