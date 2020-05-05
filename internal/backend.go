package internal

import (
	"io"
	"time"
)

type FeedType string

type FeedSubtype string

const (
	FeedTypeHTML FeedType = "html"
	FeedTypeXML  FeedType = "xml"

	FeedSubtypeAtom FeedSubtype = "atom"
)

// Store allows interact with DB
type Store interface {
	Init() error
	Feeds() FeedsRepository
	Articles() ArticlesRepository
}

// type FeedParser struct {
// 	Articles(io.Reader, Rule) []Article
// }

type Rule interface {
	Container(io.Reader) (Container, error)
	ArticleRule
}

type Container interface {
	Articles() ([]*Article, error)
}

type ArticleRuleBuilder interface {
	Get() func() ArticleRule
}
type ArticleRule interface {
	FeedID() string
	SourceID() string
	Caption() string
	Href() string
	MainContent() string
	PubDate() *time.Time
}

// Feed represents a source of articles
type Feed struct {
	ID        uint
	Address   string
	Rule      string
	Type      FeedType
	Subtype   FeedSubtype
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

// FeedsRepository allows to interact with Feeds store
type FeedsRepository interface {
	Init() error
	New(f Feed) error
	Get(id uint) (*Feed, error)
	GetAll() ([]*Feed, error)
	Delete(id uint) error
}

// Article represents one news item
type Article struct {
	ID          uint
	FeedID      uint   `db:"feed_id"`
	SourceID    string `db:"source_id"`
	Caption     string
	Href        string
	MainContent string     `db:"main_content"`
	PubDate     *time.Time `db:"pub_date"`
	DeletedAt   *time.Time `db:"deleted_at"`
}

// ArticlesRepository allows to interact with ArticlesRepository
type ArticlesRepository interface {
	Init() error
	Get(id uint) (*Article, error)
	GetByFeed(id uint) ([]*Article, error)
	Insert(list ...Article) error
	Delete(id uint) error
}
