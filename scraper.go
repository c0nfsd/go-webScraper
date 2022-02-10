package main

import (
	"crypto/tls"
	"fmt"
	"go-webScraper/helper"
	"io/ioutil"
	"net/http"
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

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
	res.Body.Close()

}
