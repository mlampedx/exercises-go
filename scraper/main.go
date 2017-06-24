package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

const (
	twitterURL      = "https://twitter.com"
	tweetSelector   = "content"                                                      // Class of Top-Level Tweet Element as well as "You May Also Like" section
	dateSelector    = "_timestamp js-short-timestamp"                                // Format example: Jun 2
	textSelector    = "TweetTextSize TweetTextSize--normal js-tweet-text tweet-text" // string
	replySelector   = "ProfileTweet-action--reply u-hiddenVisually"                  // int
	upvoteSelector  = "ProfileTweet-action--favorite u-hiddenVisually"               // int
	retweetSelector = "ProfileTweet-action--retweet u-hiddenVisually"                // int
	urlSelector     = "tweet-timestamp js-permalink js-nav js-tooltip"               // id
)

// Profile type
type Profile struct {
	Handle string
	URL    string
	Tweets []*Tweet
}

// Tweet type
type Tweet struct {
	Text     string
	Date     string
	URL      string
	Replies  int
	Upvotes  int
	Retweets int
}

func main() {
	p := &Profile{}
	p.Handle = strings.Join(os.Args[1:], "")
	p.URL = twitterURL + "/" + p.Handle
	s, err := fetch(p.URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Scraper failed: Twitter handle not found.\n")
		os.Exit(1)
	}
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Scraper failed: %v\n", err)
		os.Exit(1)
	}
	tweet := &Tweet{}
	tweets := []*Tweet{}
	p.Tweets = traverse(doc, tweet, tweets)
	renderer(p)
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

func traverse(n *html.Node, currentTweet *Tweet, tweets []*Tweet) []*Tweet {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		currentTweet = buildTweet(c, currentTweet)
		if isPopulated(currentTweet) {
			tweets = append(tweets, currentTweet)
			currentTweet = &Tweet{}
		}
		tweets = traverse(c, currentTweet, tweets)
	}
	return tweets
}

func buildTweet(n *html.Node, t *Tweet) *Tweet {
	if n.Type == html.ElementNode {
		class, ok := getAttribute(n, "class")
		if !ok {
			return t
		}
		switch class {
		case dateSelector:
			t.Date = getInnerText(n)
		case textSelector:
			t.Text = strings.Trim(getInnerText(n), "")
		case replySelector:
			t.Replies = getInnerData(n, "data-tweet-stat-count")
		case upvoteSelector:
			t.Upvotes = getInnerData(n, "data-tweet-stat-count")
		case retweetSelector:
			t.Retweets = getInnerData(n, "data-tweet-stat-count")
		case urlSelector:
			t.URL = getURL(n)
		}
	}
	return t
}

func isPopulated(t *Tweet) bool {
	return t.Date != "" &&
		t.Text != "" &&
		t.Replies != 0 &&
		t.Upvotes != 0 &&
		t.Retweets != 0 &&
		t.URL != ""
}

func getAttribute(n *html.Node, key string) (attr string, ok bool) {
	for _, attr := range n.Attr {
		if compareAttributes(attr.Key, key) {
			return attr.Val, true
		}
	}
	return "", false
}

func compareAttributes(class, desiredClass string) bool {
	return class == desiredClass
}

func getInnerText(n *html.Node) string {
	c := n.FirstChild
	if c != nil && c.Type == html.TextNode && len(c.Data) > 0 {
		return c.Data
	}
	return ""
}

func getInnerData(n *html.Node, attrName string) int {
	c := n.FirstChild
	if c != nil {
		for _, attr := range c.Attr {
			if attr.Key == attrName {
				i, err := strconv.Atoi(attr.Val)
				if err == nil {
					return i
				}
			}
		}
	}
	return 0
}

func getURL(n *html.Node) string {
	for _, attr := range n.Attr {
		if attr.Key == "href" && attr.Val != "" {
			return twitterURL + attr.Val
		}
	}
	return ""
}

func renderer(p *Profile) {
	table := template.Must(template.New("tweettable").Parse(TweetTable))
	if err := table.Execute(os.Stdout, p); err != nil {
		log.Fatal(err)
	}
}
