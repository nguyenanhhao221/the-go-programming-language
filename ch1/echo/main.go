// Echo prints its command line arguments
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// TODO:Exercise 1.3:
	// Experiment to measure the difference in running time between our
	// potentially inefficient versions and the one that uses strings.Join.
	// (Section 1.6 illustrates part of the time package,
	// and Section 11.4 shows how to write benchmark tests
	// for systematic performance evaluation.)
	fmt.Println(strings.Join(os.Args[1:], " "))
}
