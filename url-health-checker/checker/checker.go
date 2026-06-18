package checker

import (
	"fmt"
	"net/http"
	"time"
)

// func Check(url string) {
// 	client := http.Client{
// 		Timeout: 3 * time.Second,
// 	}

// 	res, err := client.Get(url)
// 	if err != nil {
// 		fmt.Println(url, "-> DOWN")
// 		return
// 	}

// 	defer res.Body.Close()

// 	fmt.Println(url, "-> UP (", res.StatusCode, ")")

// }

func Check(url string) {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	start := time.Now()

	res, err := client.Get(url)
	if err != nil {
		fmt.Println(url, "-> DOWN")
		return
	}
	defer res.Body.Close()

	duration := time.Since(start)

	fmt.Println(url, "-> UP (", res.StatusCode, ")", "[", duration, "]")
}
