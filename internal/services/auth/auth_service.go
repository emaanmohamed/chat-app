package auth

import (
	"github.com/emaanmohamed/chat-app/internal/db"
	"github.com/emaanmohamed/chat-app/internal/models"
	"github.com/emaanmohamed/chat-app/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (authService *AuthService) Register(request utils.AuthRequest) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := models.User{
		Username: request.Username,
		Password: string(hash),
	}
	if err := db.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
func (authService *AuthService) Login(request utils.AuthRequest) (string, *models.User, error) {
	var user models.User
	if err := db.DB.Where("username = ?", request.Username).First(&user).Error; err != nil {
		return "", nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return "", nil, err
	}

	token, err := utils.GenerateJWT(user.Username, user.ID.String())
	if err != nil {
		return "", nil, err
	}
	return token, &user, nil

}
