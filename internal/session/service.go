package session

import (
	"errors"
	"time"

	"github.com/alexlup06/authgate/internal/store"
)

var (
	ErrNoSession      = errors.New("no session")
	ErrSessionExpired = errors.New("session expired")
	ErrInvalidSession = errors.New("invalid session")
)

type Config struct {
	Store *store.Store

	// cookie / session settings (can be extended later)
	CookieName string
	TTL        time.Duration
	Secure     bool
}

type Service struct {
	store      *store.Store
	cookieName string
	ttl        time.Duration
	secure     bool
}

func New(cfg Config) *Service {
	cookieName := cfg.CookieName
	if cookieName == "" {
		cookieName = "authgate_session"
	}

	ttl := cfg.TTL
	if ttl == 0 {
		ttl = 24 * time.Hour
	}

	return &Service{
		store:      cfg.Store,
		cookieName: cookieName,
		ttl:        ttl,
		secure:     cfg.Secure,
	}
}

// // Create creates a new session for a user and sets the session cookie.
// func (s *Service) Create(
// 	ctx context.Context,
// 	w http.ResponseWriter,
// 	user domain.User,
// ) (*domain.Session, error) {
// 	session := domain.Session{
// 		ID:        generateSessionID(),
// 		UserID:    user.ID,
// 		ExpiresAt: time.Now().Add(s.ttl),
// 	}
//
// 	if err := s.store.CreateSession(ctx, session); err != nil {
// 		return nil, err
// 	}
//
// 	s.setCookie(w, session.ID, session.ExpiresAt)
//
// 	return &session, nil
// }
//
// // ValidateRequest validates the session from the incoming request
// // and returns the associated user.
// func (s *Service) ValidateRequest(
// 	r *http.Request,
// ) (*domain.User, error) {
// 	sessionID, err := s.readCookie(r)
// 	if err != nil {
// 		return nil, ErrNoSession
// 	}
//
// 	session, err := s.store.GetSession(r.Context(), sessionID)
// 	if err != nil {
// 		return nil, ErrInvalidSession
// 	}
//
// 	if session.ExpiresAt.Before(time.Now()) {
// 		_ = s.store.DeleteSession(r.Context(), sessionID)
// 		return nil, ErrSessionExpired
// 	}
//
// 	user, err := s.store.GetUserByID(r.Context(), session.UserID)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return &user, nil
// }
//
// // Destroy deletes the current session and clears the cookie.
// // It is intentionally idempotent.
// func (s *Service) Destroy(
// 	w http.ResponseWriter,
// 	r *http.Request,
// ) error {
// 	sessionID, err := s.readCookie(r)
// 	if err != nil {
// 		s.clearCookie(w)
// 		return nil
// 	}
//
// 	_ = s.store.DeleteSession(r.Context(), sessionID)
// 	s.clearCookie(w)
//
// 	return nil
// }
