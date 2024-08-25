// Fetching  a URL in go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const httpPrefix string = "http://"

func main() {
	for _, url := range os.Args[1:] {

		// Check for prefix in input url
		if !strings.HasPrefix(url, httpPrefix) {
			url = httpPrefix + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fail fetch: %v\n", err)
			os.Exit(1)
		}
		// Use io.Copy instead of io.Readall , Copy doesn't need to load everything in memory, so it become more efficient
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "Fail fetch copy: %v\n", err)
			os.Exit(1)
		}
		// Also print the status
		fmt.Println("status:", resp.Status)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stdout, "Fail reading: %s, err: %v\n", url, err)
			os.Exit(1)
		}

	}
}
