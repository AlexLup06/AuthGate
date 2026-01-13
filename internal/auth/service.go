package auth

import (
	"context"
	"errors"

	"github.com/alexlup06/authgate/internal/domain"
	"github.com/alexlup06/authgate/internal/store"
	"github.com/alexlup06/authgate/internal/store/tx"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserAlreadyExists  = errors.New("user already exists")
)

type Config struct {
	Store *store.Store
	Tx    *tx.Manager
}

type Service struct {
	store *store.Store
	tx    *tx.Manager
}

func New(cfg Config) *Service {
	return &Service{
		store: cfg.Store,
		tx:    cfg.Tx,
	}
}

// Signup creates a user and password auth provider atomically.
func (s *Service) Signup(
	ctx context.Context,
	email string,
	password string,
) (*domain.User, error) {
	var user domain.User
	//
	// 	err := s.tx.WithTransaction(ctx, func(txCtx context.Context) error {
	// 		exists, err := s.store.UserExistsByEmail(txCtx, email)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		if exists {
	// 			return ErrUserAlreadyExists
	// 		}
	//
	// 		hash, err := hashPassword(password)
	// 		if err != nil {
	// 			return err
	// 		}
	//
	// 		user = domain.User{
	// 			ID:    generateUserID(),
	// 			Email: email,
	// 		}
	//
	// 		if err := s.store.CreateUser(txCtx, user); err != nil {
	// 			return err
	// 		}
	//
	// 		provider := domain.AuthProvider{
	// 			ID:           generateProviderID(),
	// 			UserID:       user.ID,
	// 			Type:         domain.ProviderPassword,
	// 			PasswordHash: hash,
	// 		}
	//
	// 		if err := s.store.CreateAuthProvider(txCtx, provider); err != nil {
	// 			return err
	// 		}
	//
	// 		return nil
	// 	})
	//
	// 	if err != nil {
	// 		return nil, err
	// 	}
	//
	return &user, nil
}

//
// // Login validates credentials and returns the user.
// func (s *Service) Login(
// 	ctx context.Context,
// 	email string,
// 	password string,
// ) (*domain.User, error) {
// 	var (
// 		user     domain.User
// 		provider domain.AuthProvider
// 	)
//
// 	err := s.tx.WithTransaction(ctx, func(txCtx context.Context) error {
// 		var err error
//
// 		user, err = s.store.GetUserByEmail(txCtx, email)
// 		if err != nil {
// 			return ErrInvalidCredentials
// 		}
//
// 		provider, err = s.store.GetPasswordProviderByUserID(txCtx, user.ID)
// 		if err != nil {
// 			return ErrInvalidCredentials
// 		}
//
// 		if !verifyPassword(provider.PasswordHash, password) {
// 			return ErrInvalidCredentials
// 		}
//
// 		return nil
// 	})
//
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return &user, nil
// }
