package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const weblioJeURL = "http://ejje.weblio.jp/content/"

func getSearchWords(words []string) string {
	return strings.Join(words, "+")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			`Usage: Enter the search words
        [Example] come from`)
	}
	flag.Parse()

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
