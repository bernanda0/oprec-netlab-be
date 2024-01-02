package handlers

import (
	"log"

	"github.com/bernanda0/oprec-netlab-be/database/sqlc"
)

type Handler struct {
	l *log.Logger
	q *sqlc.Queries
	c *uint
	u *AuthedUser
}

type DriveApiHandler struct {
	h *Handler
}

type AuthedUser struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
}
