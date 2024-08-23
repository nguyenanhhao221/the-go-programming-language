// Dup3 similar to dup2, but we read the entire file into memory rather than line by line with the Scan()
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// Readfile will read the whole file into memory, then we split this data line by line manually
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3 err: %v", err)
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}

	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
