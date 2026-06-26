package main

import (
	"io/fs"
	"path/filepath"
)

// ScanFolder walks through the directory and returns all file paths.
func ScanFolder(root string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {

		// If there is an error accessing a file or directory, stop.
		if err != nil {
			return err
		}

		// Ignore directories.
		if d.IsDir() {
			return nil
		}

		// Add the file path to our slice.
		files = append(files, path)

		return nil
	})

	return files, err
}
