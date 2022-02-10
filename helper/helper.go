package helper

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/steelx/extractlinks"
)

var (
	config = &tls.Config{
		InsecureSkipVerify: true,
	}

	transport = &http.Transport{
		TLSClientConfig: config,
	}
	netClient = &http.Client{
		Transport: transport,
	}
)

func ErrCheck(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Urlcrawl(href string) {
	fmt.Printf("Crawling url -----> %v \n", href)
	res, err := netClient.Get(href)
	ErrCheck(err)
	defer res.Body.Close()

	links, err := extractlinks.All(res.Body)
	ErrCheck(err)

	for _, link := range links {
		Urlcrawl(ToFixedURL(link.Href, href))
	}
}

func ToFixedURL(href, baseURL string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}

	base, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}

	toFixedUri := base.ResolveReference(uri)

	return toFixedUri.String()
}
