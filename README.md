# Wayback Machine Snapshot Fetcher

## Overview

The **Waybackfetch** is a command-line tool written in Go that allows users to easily retrieve archived snapshots of web pages from the Internet Archive’s Wayback Machine. With this tool, users can fetch all available snapshot URLs for a given webpage or a list of web pages, enabling easy access to historical versions of content.

## Features

```console
└─# waybackfetch -h


         _  _  _             _                 _     _______              _     
        | || || |           | |               | |   (_______)   _        | |    
        | || || | ____ _   _| | _   ____  ____| |  _ _____ ____| |_  ____| | _  
        | ||_|| |/ _  | | | | || \ / _  |/ ___) | / )  ___) _  )  _)/ ___) || \ 
        | |___| ( ( | | |_| | |_) | ( | ( (___| |< (| |  ( (/ /| |_( (___| | | |
         \______|\_||_|\__  |____/ \_||_|\____)_| \_)_|   \____)\___)____)_| |_|
                                  (____/                                                    


              v1.0 Created by KathanP19

Usage:
  -u <url>       Fetch snapshots for a single URL
  -l <file>      File containing list of URLs to fetch snapshots for
  -o <file>      Output file to save the results
  --silent       Enable silent mode, only print URLs
  -h, --help     Show this help message and exit
```


- **Single URL Fetching**: Quickly retrieve all snapshot URLs for a specific web page using the `-u` flag.
- **Batch Processing**: Process multiple URLs by providing a file containing a list of URLs with the `-l` flag.
- **Output Options**: Save the retrieved snapshot URLs to a specified output file using the `-o` flag while also printing the results to the console.
- **Silent Mode**: Enable a clean output experience with the `--silent` flag to display only the results without additional console messages.
- **Input from Standard Input**: Supports reading URLs directly from standard input, allowing for flexible usage in scripts or pipelines.
- **Help Command**: Provides a helpful usage guide and flag descriptions when the `-h` or `--help` flags are used.

## Installation

1. Ensure you have Go installed on your machine.
2. You can install the tool using the following commands:
   ```console
   go install github.com/KathanP19/waybackfetch@latest
   ```
   OR
   ```console
   git clone https://github.com/KathanP19/waybackfetch.git
   cd waybackfetch
   go install
   ```
   
## Usage
```console
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
