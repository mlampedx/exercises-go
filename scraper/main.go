package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

const tweetClass = "TweetTextSize TweetTextSize--normal js-tweet-text tweet-text"

// Tweets type
type Tweets struct {
	Handle string
	List   []*Tweet
}

// Tweet type
type Tweet struct {
	Text string
}

func main() {
	handle := strings.Join(os.Args[1:], "")
	url := "https://twitter.com/" + handle
	s, err := fetch(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Scraper failed: Twitter handle not found.\n")
		os.Exit(1)
	}
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Scraper failed: %v\n", err)
		os.Exit(1)
	}
	output := &Tweets{handle, scrape(nil, doc)}
	renderer(output)
}

func fetch(url string) (bodyString string, err error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	bodyString = string(bodyBytes)
	res.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	return bodyString, nil
}

func scrape(list []*Tweet, n *html.Node) []*Tweet {
	if isTweetNode(n) {
		// Todo: Escape n.Data string
		newTweet := &Tweet{n.Data}
		list = append(list, newTweet)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		list = scrape(list, c)
	}
	return list
}

func isTweetNode(n *html.Node) bool {
	if n.Type == html.TextNode {
		// Todo: Figure out how to parse class from html.TextNode or text from html.ElementNode to ensure tweets are parsed
		// class, ok := getAttribute(n, "class")
		// if !ok {
		// 	return false
		// }
		// return class == tweetClass
		return true
	}
	return false
}

func getAttribute(n *html.Node, key string) (attr string, ok bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func renderer(tweets *Tweets) {
	table := template.Must(template.New("tweettable").Parse(TweetTable))
	if err := table.Execute(os.Stdout, tweets); err != nil {
		log.Fatal(err)
	}
}
