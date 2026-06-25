package main

import (
	"bufio"
	"os"
	"strings"
)

func IsValidURL(url string) bool {
	return strings.HasPrefix(url, "http")
}
func readURLs(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	return urls, scanner.Err()

}
