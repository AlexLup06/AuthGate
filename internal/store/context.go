package store

import (
	"context"

	"gorm.io/gorm"
)

type contextKey string

const dbKey contextKey = "store.tx.db"

// dbFromContext returns a transactional DB if present, otherwise base DB.
func (s *Store) dbFromContext(ctx context.Context) *gorm.DB {
	if txDB, ok := ctx.Value(dbKey).(*gorm.DB); ok {
		return txDB
	}
	return s.db
}
