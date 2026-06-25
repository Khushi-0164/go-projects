# 🔍 GitHub Profile Finder
A simple interactive command-line application that searches for GitHub user profiles and displays their repositories. Enter any GitHub username and instantly fetch user information and repository details from the GitHub API.

## 🚀 Features
* Interactive CLI interface with continuous search capability.
* Fetch detailed GitHub user profile information by username.
* Display all public repositories associated with a GitHub user.
* Input validation for non-empty usernames.
* Multiple user searches in a single session without restarting.
* Error handling for network issues and invalid usernames.

## 🛠️ Go Concepts Demonstrated
* Buffered input/output using `bufio.NewReader` for reading user input from stdin.
* String manipulation with `strings.TrimSpace` and `strings.ToLower` for input processing.
* Error handling with `err` return values and conditional checking.
* Control flow with loops (`for {}`) and conditional statements (`if`, `switch`).
* HTTP requests to GitHub REST API (in accompanying `profile.go` and `repositories.go` files).
* JSON unmarshaling for parsing API responses.

## 📖 How to Run

### Option 1: Using `go run`
```bash
go run main.go
```

### Option 2: Build and Execute
```bash
go build -o github-profile-finder
./github-profile-finder
```

### Example Usage:
```bash
$ go run main.go

Enter GitHub Username:
golang
Searching for: golang

User Profile:
Name: Go
Bio: The Go Programming Language
Public Repos: 500+
Followers: 15000+
Profile URL: https://github.com/golang

Repositories:
- go
- website
- proposal
- build
- tools

Search another user? (y/n): y
Enter GitHub Username:
torvalds
Searching for: torvalds

User Profile:
Name: Linus Torvalds
Bio: Linux kernel developer
Public Repos: 200+
Followers: 250000+
Profile URL: https://github.com/torvalds

Repositories:
- linux
- subsurface-for-android
- git

Search another user? (y/n): n
Goodbye!
```