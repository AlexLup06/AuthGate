package handlers

import (
	"net/http"

	"github.com/alexlup06/authgate/internal/auth"
	"github.com/alexlup06/authgate/internal/session"
)

type AuthHandler struct {
	auth    *auth.Service
	session *session.Service
}

func NewAuthHandler(
	authService *auth.Service,
	sessionService *session.Service,
) *AuthHandler {
	return &AuthHandler{
		auth:    authService,
		session: sessionService,
	}
}

func (h *AuthHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("login page (placeholder)"))
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("login action (placeholder)"))
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("logout action (placeholder)"))
}

func (h *AuthHandler) SignupPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("signup page (placeholder)"))
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("signup action (placeholder)"))
}
