package html

import (
	"fmt"
	"io"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/rpoletaev/newsreader/internal"
)

type Rule struct {
	Root string // .l-row span[itemprop='itemListElement']
}

func (rule *Rule) Container(r io.Reader) (internal.Container, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	s := doc.Find(rule.Root)
	if s == nil {
		return nil, fmt.Errorf("unable to find container")
	}

	return &Container{container: s}, nil
}

type Container struct {
	container *goquery.Selection
}

type ArticleRule struct {
	selection *goquery.Selection
}

func (r *ArticleRule) FeedID() string   { return "" }
func (r *ArticleRule) SourceID() string { return "" }
func (r *ArticleRule) Caption() string {
	t, _ := r.selection.Find("meta[itemprop='name']").Attr("content")
	return t
}

func (r *ArticleRule) Href() string {
	v, _ := r.selection.Find("meta[itemprop='url']").Attr("content")
	return v
}
func (r *ArticleRule) MainContent() string { return "" }
func (r *ArticleRule) PubDate() *time.Time { return nil }

func (c *Container) get(s *goquery.Selection)
func (c *Container) Get() ArticleRule {

}
func (c *Container) Articles() ([]*internal.Article, error) {
	result := []*internal.Article{}
	c.container.Each(func(i int, s *goquery.Selection) {
		rule := ArticleRule{
			selection: s,
		}
		// For each item found, get the band and title
		fmt.Println("href", rule.Href())
		fmt.Println("title", rule.Caption())
	})

	return result, nil
}

// func (c *Container) FeedID() string      { return "" }
// func (c *Container) SourceID() string    { return "" }
// func (c *Container) Caption() string     { return "" }
// func (c *Container) Href() string        { return "" }
// func (c *Container) MainContent() string { return "" }
// func (c *Container) PubDate() *time.Time { return nil }
