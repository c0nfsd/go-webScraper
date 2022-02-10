package main

import (
	"crypto/tls"
	"fmt"
	"go-webScraper/helper"
	"net/http"
	"os"

	"github.com/steelx/extractlinks"
)

func main() {

	arguments := os.Args[1:]

	if len(arguments) == 0 {
		fmt.Println("Url cannot be empty, please enter a valid URL, example: https://wikileaks.org")
		os.Exit(1)
	}
	baseUrl := arguments[0]
	fmt.Println("baseUrl", baseUrl)

	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{
		TLSClientConfig: config,
	}
	netClient := &http.Client{
		Transport: transport,
	}

	res, err := netClient.Get(baseUrl)
	helper.ErrCheck(err)
	defer res.Body.Close()

	links, err := extractlinks.All(res.Body)
	helper.ErrCheck(err)

	for i, link := range links {
		fmt.Printf("Index %v --Link %+v\n", i+1, link)
	}

}
