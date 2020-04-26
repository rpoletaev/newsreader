package postgres

import (
	"time"

	"github.com/rpoletaev/newsreader/internal"
)

const (
	FeedInsertTemplate     = `INSERT INTO feeds (address, rule, created_at) VALUES (:address, :rule, :created_at)`
	FeedSelectByIDTemplate = `SELECT * FROM feeds WHERE id = $1 AND deleted_at = NULL`
	FeedSelectAllTemplate  = `SELECT * FROM feeds WHERE deleted_at = NULL`
	FeedDeleteTemplate     = `UPDATE feeds SET deleted_at = :deleted WHERE id = :id`

	feedSchema = `CREATE TABLE IF NOT EXISTS feeds (
		id integer PRIMARY KEY,
		address text UNIQUE,
		rule text,
		created_at date,
		CONSTRAINT unique_address UNIQUE(address)
		)`
)

type FeedsRepository Store

func (s *Store) Feeds() internal.FeedsRepository {
	return (*FeedsRepository)(s)
}

func (r *FeedsRepository) Init() error {
	_, err := r.db.Exec(feedSchema)
	return err
}

func (r *FeedsRepository) New(f internal.Feed) error {
	_, err := r.db.NamedExec(FeedInsertTemplate, &f)
	return err
}

func (r *FeedsRepository) Get(id uint) (*internal.Feed, error) {
	f := &internal.Feed{}
	if err := r.db.Get(f, FeedSelectByIDTemplate, id); err != nil {
		return nil, err
	}

	return f, nil
}

func (r *FeedsRepository) GetAll() ([]*internal.Feed, error) {
	feeds := []*internal.Feed{}
	if err := r.db.Select(&feeds, FeedSelectAllTemplate); err != nil {
		return nil, err
	}
	return feeds, nil
}

func (r *FeedsRepository) Delete(id uint) error {
	_, err := r.db.NamedExec(FeedDeleteTemplate, map[string]interface{}{
		"deleted": time.Now(),
		"id":      id,
	})
	return err
}
