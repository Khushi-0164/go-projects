package main

import "time"

type GitHubUser struct {
	Login       string    `json:"login"`
	Name        string    `json:"name"`
	Bio         string    `json:"bio"`
	Followers   int       `json:"followers"`
	Following   int       `json:"following"`
	PublicRepos int       `json:"public_repos"`
	Location    string    `json:"location"`
	Company     string    `json:"company"`
	HTMLURL     string    `json:"html_url"`
	CreatedAt   time.Time `json:"created_at"`
}

type Repository struct {
	Name            string `json:"name"`
	Language        string `json:"language"`
	StargazersCount int    `json:"stargazers_count"`
}
