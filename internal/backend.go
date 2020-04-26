package internal

import "time"

// Store allows interact with DB
type Store interface {
	Init() error
	Feeds() FeedsRepository
	Articles() ArticlesRepository
}

// Feed represents a source of articles
type Feed struct {
	ID        uint
	Address   string
	Rule      string
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
