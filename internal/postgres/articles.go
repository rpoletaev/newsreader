package postgres

import (
	"time"

	"github.com/pkg/errors"
	"github.com/rpoletaev/newsreader/internal"
)

const (
	articlesInsertTemplate = `INSERT INTO articles (
		feed_id,
		source_id,
		caption,
		href,
		main_content,
		pub_date
	) VALUES (
		:feed_id, 
		:source_id, 
		:caption,
		:href,
		:main_content,
		:pub_date
	)`

	articlesGetByIDTemplate = `SELECT * FROM articles WHERE id = $1 AND deleted_at = NULL`

	articleGetByFeedTemplate = `SELECT * FROM articles WHERE feed_id = $1 AND deleted_at = NULL`

	ArticlesDeleteTemplate = `UPDATE articles SET deleted_at = :deleted WHERE id = :id`

	articleSchema = `CREATE TABLE IF NOT EXISTS articles (
		id integer PRIMARY KEY,
		feed_id integer NOT NULL,
		source_id NOT NULL,
		caption NOT NULL,
		href NOT NULL,
		main_content,
		pub_date  NOT NULL,
		CONSTRAINT unique_address UNIQUE(address),
		CONSTRAINT FOREIGN KEY (feed_id) REFERENCES feeds(id)
		)`
)

type ArticlesRepository Store

func (s *Store) Articles() internal.ArticlesRepository {
	return (*ArticlesRepository)(s)
}

func (s *ArticlesRepository) Init() error {
	_, err := s.db.Exec(articleSchema)
	return err
}

func (r *ArticlesRepository) Get(id uint) (*internal.Article, error) {
	a := &internal.Article{}
	if err := r.db.Get(a, articlesGetByIDTemplate, id); err != nil {
		return nil, err
	}

	return a, nil
}

func (r *ArticlesRepository) GetByFeed(id uint) ([]*internal.Article, error) {
	list := []*internal.Article{}
	if err := r.db.Select(&list, articleGetByFeedTemplate, id); err != nil {
		return nil, err
	}
	return list, nil
}

func (r *ArticlesRepository) Insert(articles ...internal.Article) error {
	tx, err := r.db.Begin()
	if err != nil {
		return errors.Wrap(err, "unable to get tx on insert articles")
	}

	for _, article := range articles {
		tx.Exec(articlesInsertTemplate, article)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "unable commit transaction on insert articles")
	}

	return nil
}

func (r *ArticlesRepository) Delete(id uint) error {
	if _, err := r.db.Exec(ArticlesDeleteTemplate, map[string]interface{}{
		"id":      id,
		"deleted": time.Now(),
	}); err != nil {
		return err
	}

	return nil
}
