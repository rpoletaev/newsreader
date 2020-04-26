package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rpoletaev/newsreader/internal"
	"github.com/rs/zerolog"
)

// Config for postgres connection
type Config struct {
	Driver string `envconfig:"DRIVER"`
	URI    string `envconfig:"URI"`
}

// Store implements internal Store with postgres db
type Store struct {
	db     *sqlx.DB
	Logger zerolog.Logger
}

var _ internal.Store = (*Store)(nil)

// Connect fulfill db connection
func (s *Store) Connect() error {
	db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	if err != nil {
		return errors.Wrap(err, "on connect to postgres")
	}

	s.db = db

	return nil
}

func (s *Store) Init() error {
	if err := s.Feeds().Init(); err != nil {
		return errors.Wrap(err, "on init feeds")
	}

	if err := s.Articles().Init(); err != nil {
		return errors.Wrap(err, "on init articles")
	}

	return nil
}
