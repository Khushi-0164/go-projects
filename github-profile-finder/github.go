package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"
)

func FetchRepositories(username string) {
	url := fmt.Sprintf(
		"https://api.github.com/users/%s/repos",
		username,
	)
	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	req.Header.Set(
		"User-Agent",
		"Go-GitHub-Finder",
	)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var repos []Repository
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&repos)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\nTop 5 Repositories")
	fmt.Println("-------------------")
	sort.Slice(repos, func(i, j int) bool {
		return repos[i].StargazersCount > repos[j].StargazersCount
	})
	limit := 5

	if len(repos) < limit {
		limit = len(repos)
	}
	for i := 0; i < limit; i++ {
		repo := repos[i]

		fmt.Printf(
			"%s | %s | ⭐ %d\n",
			repo.Name,
			displayValue(repo.Language),
			repo.StargazersCount,
		)
	}
}

func FetchProfile(username string) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	fmt.Println(url)

	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	req.Header.Set(
		"User-Agent",
		"Go-GitHub-Finder",
	)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == http.StatusNotFound {
		fmt.Println("User not found")
		return
	}

	var user GitHubUser
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	bio := displayValue(user.Bio)
	company := displayValue(user.Company)
	location := displayValue(user.Location)

	fmt.Println("========================================")
	fmt.Println("          GitHub Profile                ")
	fmt.Println("========================================")
	fmt.Println(" ")
	fmt.Printf("%-15s : %s\n", "Name", user.Name)
	fmt.Printf("%-15s : %s\n", "Username", user.Login)
	fmt.Printf("%-15s : %s\n", "Bio", bio)
	fmt.Printf("%-15s : %s\n", "Location", location)
	fmt.Printf("%-15s : %s\n", "Company", company)

	fmt.Println()

	fmt.Printf("%-15s : %d\n", "Followers", user.Followers)
	fmt.Printf("%-15s : %d\n", "Following", user.Following)
	fmt.Printf("%-15s : %d\n", "Public Repos", user.PublicRepos)

	fmt.Println()

	fmt.Printf("%-15s : %s\n", "Joined", user.CreatedAt.Format("2006-01-02"))
	fmt.Printf("%-15s : %s\n", "Profile URL", user.HTMLURL)
	fmt.Println(" ")
	fmt.Println("========================================")
}
func displayValue(value string) string {
	if value == "" {
		return "Not provided"
	}
	return value
}
