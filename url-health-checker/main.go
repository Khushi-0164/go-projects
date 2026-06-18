package main

import (
	"fmt"
	"os"
	"sync"
	"url-health-checker/checker"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide at least one URL")
		return
	}
	var wg sync.WaitGroup
	// Example usage when a checker package is available:
	// 	for _, url := range os.Args[1:] {
	// 		checker.Check(url)
	// 	}
	// }
	for _, url := range os.Args[1:] {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()
			checker.Check(u)
		}(url)
	}

	wg.Wait()
}
