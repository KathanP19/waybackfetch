package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// ANSI color codes for styling
const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	cyan   = "\033[36m"
)

// Banner with colors
func printBanner() {
	fmt.Println(string(cyan + `

	 _  _  _             _                 _     _______              _     
	| || || |           | |               | |   (_______)   _        | |    
	| || || | ____ _   _| | _   ____  ____| |  _ _____ ____| |_  ____| | _  
	| ||_|| |/ _  | | | | || \ / _  |/ ___) | / )  ___) _  )  _)/ ___) || \ 
	| |___| ( ( | | |_| | |_) | ( | ( (___| |< (| |  ( (/ /| |_( (___| | | |
	 \______|\_||_|\__  |____/ \_||_|\____)_| \_)_|   \____)\___)____)_| |_|
				  (____/                                                    

` + reset))
	fmt.Println(string(yellow + "              v1.0 Created by KathanP19" + reset))
	fmt.Println()
}

// WaybackResponse holds the snapshot timestamps returned from the Wayback Machine API
type WaybackResponse [][]string

// FetchSnapshotUrls fetches all snapshot URLs for a given URL
func FetchSnapshotUrls(url string, silent bool, output io.Writer) error {
	apiUrl := fmt.Sprintf("http://web.archive.org/cdx/search/cdx?url=%s&output=json&fl=timestamp", url)

	resp, err := http.Get(apiUrl)
	if err != nil {
		return fmt.Errorf(red+"error fetching data:"+reset+" %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf(red+"error reading response body:"+reset+" %v", err)
	}

	var timestamps WaybackResponse
	err = json.Unmarshal(body, &timestamps)
	if err != nil {
		return fmt.Errorf(red+"error parsing JSON:"+reset+" %v", err)
	}

	for _, ts := range timestamps[1:] {
		timestamp := ts[0]
		snapshotUrl := fmt.Sprintf("https://web.archive.org/web/%sif_/%s", timestamp, url)
		fmt.Fprintln(output, snapshotUrl)
	}
	return nil
}

func main() {
	url := flag.String("u", "", "Single URL to fetch snapshots for")
	list := flag.String("l", "", "File containing list of URLs to fetch snapshots for")
	silent := flag.Bool("silent", false, "Enable silent mode, only print URLs")
	outputFile := flag.String("o", "", "Output file to write results")

	// Custom help message
	flag.Usage = func() {
		printBanner()
		fmt.Println("Usage:")
		fmt.Println("  -u <url>       Fetch snapshots for a single URL")
		fmt.Println("  -l <file>      File containing list of URLs to fetch snapshots for")
		fmt.Println("  -o <file>      Output file to save the results")
		fmt.Println("  --silent       Enable silent mode, only print URLs")
		fmt.Println("  -h, --help     Show this help message and exit")
	}

	flag.Parse()

	// Print the banner unless silent mode is enabled
	if !*silent {
		printBanner()
	}

	var output io.Writer = os.Stdout
	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Printf("Error creating output file: %v\n", err)
			return
		}
		defer file.Close()
		output = io.MultiWriter(os.Stdout, file)
	}

	processUrls := func(url string) {
		if !*silent {
			fmt.Printf(green+"\nFetching snapshots for URL:"+reset+" %s\n", url)
		}
		if err := FetchSnapshotUrls(url, *silent, output); err != nil && !*silent {
			fmt.Println("Error:", err)
		}
	}

	if *url != "" {
		processUrls(*url)
		if *outputFile != "" && !*silent {
			fmt.Printf(green+"\nResults have been saved to:"+reset+" %s\n", *outputFile)
		}
		return
	}

	if *list != "" {
		file, err := os.Open(*list)
		if err != nil {
			if !*silent {
				fmt.Printf(red+"Error opening file:"+reset+" %v\n", err)
			}
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			processUrls(scanner.Text())
		}

		if err := scanner.Err(); err != nil && !*silent {
			fmt.Printf(red+"Error reading file:"+reset+" %v\n", err)
		}

		if *outputFile != "" && !*silent {
			fmt.Printf(green+"\nResults have been saved to:"+reset+" %s\n", *outputFile)
		}
		return
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			processUrls(scanner.Text())
		}
		if err := scanner.Err(); err != nil && !*silent {
			fmt.Printf(red+"Error reading stdin:"+reset+" %v\n", err)
		}

		if *outputFile != "" && !*silent {
			fmt.Printf(green+"\nResults have been saved to:"+reset+" %s\n", *outputFile)
		}
		return
	}

	if !*silent {
		fmt.Println(red + "Please provide -u <URL> for a single URL, -l <file> for a list of URLs, or input via stdin" + reset)
	}
}
