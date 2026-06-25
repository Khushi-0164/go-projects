package main

type Result struct {
	URL          string `json:"url"`
	StatusCode   int    `json:"status_code"`
	ResponseTime string `json:"response_time"`
	Error        string `json:"error"`
}
