package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const weblioJeURL = "http://ejje.weblio.jp/content/"

func getSearchWords(words []string) string {
	return strings.Join(words, "+")
}

const usage = `goje is a Japanese<->English dictionary using weblio.

Example (simple):
	goje hello

Example (how to search idiom):
	goje first of all
`

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println(usage)
		return
	}

	args := flag.Args()
	words := getSearchWords(args)
	searchURL := weblioJeURL + words

	doc, err := goquery.NewDocument(searchURL)
	if err != nil {
		fmt.Println("url scarapping failed")
	}

	notfound := true
	doc.Find(".content-explanation").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Text())
		notfound = false
	})
	if notfound {
		fmt.Println("Not found.")
	}
}
