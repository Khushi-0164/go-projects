package main

import (
	"file-health-checker/checker"
	"fmt"
	"os"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide at least one file path")
		return
	}
	var wg sync.WaitGroup
	for _, path := range os.Args[1:] {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			checker.Check(p)
		}(path)
	}
	wg.Wait()
}
