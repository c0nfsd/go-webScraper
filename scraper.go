package main

import (
	"crypto/tls"
	"fmt"
	"go-webScraper/helper"
	"net/http"

	"github.com/steelx/extractlinks"
)

func main() {

	baseUrl := "https://youtube.com/jsfunc"

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

	fmt.Println(links)

	//res.Body.Close()

}
