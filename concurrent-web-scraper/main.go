package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func checkURL(url string) Result {
	start := time.Now()

	resp, err := http.Get(url)

	if err != nil {
		return Result{
			URL:   url,
			Error: err.Error(),
		}
	}
	defer resp.Body.Close()

	duration := time.Since(start)

	return Result{
		URL:          url,
		StatusCode:   resp.StatusCode,
		ResponseTime: duration.String(),
	}

}
func worker(jobs <-chan string, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range jobs {
		fmt.Println("Worker processing:", url)
		result := checkURL(url)
		results <- result
	}
}

func main() {
	start := time.Now()
	urls, err := readURLs("urls.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var wg sync.WaitGroup
	workers := 3
	jobs := make(chan string)
	results := make(chan Result)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}
	go func() {
		for _, url := range urls {
			jobs <- url
		}

		close(jobs)
	}()
	var allResults []Result
	for i := 0; i < len(urls); i++ {
		result := <-results

		allResults = append(allResults, result)

		fmt.Printf("%+v\n", result)
	}
	err = saveReport(allResults)
	if err != nil {
		fmt.Println("Error saving report:", err)
		return
	}
	wg.Wait()
	fmt.Println("Total Time:", time.Since(start))
}
