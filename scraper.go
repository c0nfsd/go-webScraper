package main

import (
	"fmt"
	"go-webScraper/helper"
	"os"
)

func main() {

	arguments := os.Args[1:]

	if len(arguments) == 0 {
		fmt.Println("Url cannot be empty, please enter a valid URL, example: https://wikileaks.org")
		os.Exit(1)
	}

	baseURL := arguments[0]
	go func() {
		helper.Queue <- baseURL
	}()

	for href := range helper.Queue {
		if !helper.Visited[href] && helper.IsSameDomain(href, baseURL) {
			helper.Urlcrawl(href)
		}

	}

}
