package main

import (
	"flag"
)

func scan(path string) {
	print(path)
}

func stats(email string) {
	print(email)
}

func main() {
	var folder string
	var email string

	flag.StringVar(&folder, "add", "", "add a folder to scan for Git repositories")
	flag.StringVar(&email, "email", "your@email.com", "email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}

	stats(email)
}
