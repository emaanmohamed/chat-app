package types

import "github.com/emaanmohamed/chat-app/internal/models"

type UserStore interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(models.User) error
}
