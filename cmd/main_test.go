package main

import (
	"testing"
	"fmt"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func TestRBC(t *testing.T) {
	res, err := http.Get("https://www.rbc.ru/story/5e7ceff09a7947293fe73b75")
	if err != nil {
		t.Error(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Errorf("status code error: %d %s", res.StatusCode, res.Status)
		return
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		t.Error(err)
		return
	}

	// Find the review items
	doc.Find(".l-row span[itemprop='itemListElement']").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		href,_ := s.Find("meta[itemprop='url']").Attr("content")
		title,_ := s.Find("meta[itemprop='name']").Attr("content")
		fmt.Println("href", href)
		fmt.Println("title", title)
	})
}