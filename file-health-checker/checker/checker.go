package checker

import (
	"fmt"
	"os"
)

func Check(path string) {
	info, err := os.Stat(path)
	if err != nil {
		fmt.Printf("%s -> NOT FOUND\n", path)
		return
	}

	fmt.Println("----------")
	fmt.Println("File:", path)
	fmt.Println("Status: FOUND")
	fmt.Println("Size:", info.Size(), "bytes")
	fmt.Println("Last Modified:", info.ModTime())
}
