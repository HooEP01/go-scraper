package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Make an HTTP GET request to the target URL
	response, err := http.Get("https://example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Parse the HTML document
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the entire HTML document
	html, err := doc.Html()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(html)

	// Find and extract the desired data from the HTML document
	doc.Find("h1").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}
