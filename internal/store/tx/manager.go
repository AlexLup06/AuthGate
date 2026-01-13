package tx

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/alexlup06/authgate/internal/store"
	"gorm.io/gorm"
)

type Manager struct {
	store *store.Store
}

func New(store *store.Store) *Manager {
	return &Manager{store: store}
}

type contextKey string

const (
	dbKey contextKey = "store.tx.db"
	txKey contextKey = "store.tx.state"
)

type transactionState struct {
	committed  bool
	rolledBack bool
}

// ---------------------------------------------------------------------
// Public API
// ---------------------------------------------------------------------

func (m *Manager) WithTransaction(
	ctx context.Context,
	fn func(context.Context) error,
) error {
	txCtx, cancel, err := m.withCancel(ctx)
	if err != nil {
		return err
	}
	defer cancel()

	if err := fn(txCtx); err != nil {
		if !isFinished(txCtx) {
			_ = m.Rollback(txCtx)
		}
		return err
	}

	if !isFinished(txCtx) {
		return m.Commit(txCtx)
	}

	return nil
}

func (m *Manager) Commit(ctx context.Context) error {
	db, err := getDB(ctx)
	if err != nil {
		return err
	}

	if err := db.Commit().Error; err != nil {
		return err
	}

	setCommitted(ctx)
	return nil
}

func (m *Manager) Rollback(ctx context.Context) error {
	if isCommitted(ctx) {
		return errors.New("transaction already committed")
	}

	db, err := getDB(ctx)
	if err != nil {
		return err
	}

	if err := db.Rollback().Error; err != nil {
		return err
	}

	setRolledBack(ctx)
	return nil
}

// ---------------------------------------------------------------------
// Internal helpers
// ---------------------------------------------------------------------

func (m *Manager) withCancel(
	parent context.Context,
) (context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithCancel(parent)
	txCtx, err := m.withContext(ctx)
	return txCtx, cancel, err
}

func (m *Manager) withContext(
	parent context.Context,
) (context.Context, error) {
	session := m.store.DB().
		WithContext(parent).
		Begin(&sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
		})

	if session.Error != nil {
		return nil, session.Error
	}

	ctx := context.WithValue(parent, dbKey, session)
	ctx = context.WithValue(ctx, txKey, &transactionState{})

	go m.cleanup(ctx)

	return ctx, nil
}

func (m *Manager) cleanup(ctx context.Context) {
	<-ctx.Done()
	if !isFinished(ctx) {
		_ = m.Rollback(ctx)
	}
}

// ---------------------------------------------------------------------
// Context utilities (store-internal only)
// ---------------------------------------------------------------------

func getDB(ctx context.Context) (*gorm.DB, error) {
	db, ok := ctx.Value(dbKey).(*gorm.DB)
	if !ok {
		return nil, fmt.Errorf("no transaction in context")
	}
	return db, nil
}

func isCommitted(ctx context.Context) bool {
	state, ok := ctx.Value(txKey).(*transactionState)
	return ok && state.committed
}

func isFinished(ctx context.Context) bool {
	state, ok := ctx.Value(txKey).(*transactionState)
	return ok && (state.committed || state.rolledBack)
}

func setCommitted(ctx context.Context) {
	if state, ok := ctx.Value(txKey).(*transactionState); ok {
		state.committed = true
	}
}

func setRolledBack(ctx context.Context) {
	if state, ok := ctx.Value(txKey).(*transactionState); ok {
		state.rolledBack = true
	}
}
