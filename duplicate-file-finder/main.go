package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <folder_path>")
		return
	}

	folderPath := os.Args[1]

	// Validate the folder
	info, err := os.Stat(folderPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if !info.IsDir() {
		fmt.Println("The given path is not a directory.")
		return
	}

	fmt.Println("Scanning folder:", folderPath)

	files, err := ScanFolder(folderPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	hashMap := make(map[string][]string)

	for _, file := range files {

		hash, err := HashFile(file)
		if err != nil {
			fmt.Println("Error hashing:", file)
			continue
		}

		hashMap[hash] = append(hashMap[hash], file)
	}

	fmt.Println("\nDuplicate Files:")

	found := false

	group := 1

	for _, fileList := range hashMap {
		if len(fileList) > 1 {
			found = true

			fmt.Printf("\nDuplicate Group %d\n", group)
			fmt.Println("--------------------")

			for _, file := range fileList {
				fmt.Println(file)
			}

			group++
		}
	}

	if !found {
		fmt.Println("No duplicate files found.")
	}
}
