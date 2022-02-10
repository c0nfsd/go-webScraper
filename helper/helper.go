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
	Queue   = make(chan string)
	Visited = make(map[string]bool)
)

func errCheck(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func Urlcrawl(href string) {
	Visited[href] = true
	fmt.Printf("Crawling url -----> %v \n", href)
	resp, err := netClient.Get(href)
	errCheck(err)
	defer resp.Body.Close()

	links, err := extractlinks.All(resp.Body)
	errCheck(err)

	for _, link := range links {
		absoluteURL := ToFixedURL(link.Href, href)
		go func(url string) {
			Queue <- url
		}(absoluteURL)

	}
}

func IsSameDomain(href, baseURL string) bool {
	uri, err := url.Parse(href)
	if err != nil {
		return false
	}

	parentUri, err := url.Parse(baseURL)
	if err != nil {
		return false
	}

	if uri.Host != parentUri.Host {
		return false
	}
	return true
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
