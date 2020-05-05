package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	scrapHTTP("https://www.rbc.ru/story/5e7ceff09a7947293fe73b75")
	ExampleScrape()
}

func scrapHTTP(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".l-row span[itemprop='itemListElement']").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		href, _ := s.Find("meta[itemprop='url']").Attr("content")
		title, _ := s.Find("meta[itemprop='name']").Attr("content")
		fmt.Println("href", href)
		fmt.Println("title", title)
	})
}

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})
}
