package service

import (
	"database/sql"

	"github.com/hako/branca"
)

// Service contains the core logic. You can use it to back a REST, GraphQL or RPC API :)
type Service struct {
	db     *sql.DB
	codec  *branca.Branca
	origin string
}

// New service implementation.
func New(db *sql.DB, codec *branca.Branca, origin string) *Service {
	return &Service{
		db:     db,
		codec:  codec,
		origin: origin,
	}
}
