# GitHub Scraper

Scrapes the HTML from a GitHub issues page, parses issues, and
generates a templated HTML table displaying the results.

Given a working installation of golang with workspace paths properly set,
cd into this directory and run the following commands from the terminal:

```
go build
./github repo:golang/go commenter:gopherbot json encoder > issues.html
```
