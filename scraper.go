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

	func() {
		helper.Queue <- arguments[0]
	}()

	for href := range helper.Queue {
		helper.Urlcrawl(href)

	}

}
