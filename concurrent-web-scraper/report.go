package main

import (
	"encoding/json"
	"os"
)

func saveReport(results []Result) error {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("report.json", data, 0644)
}
