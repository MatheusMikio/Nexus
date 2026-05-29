package service

import (
	"time"

	authLogin "github.com/MatheusMikio/Nexus/internal/auth"
	"github.com/MatheusMikio/Nexus/internal/domain/dtos/auth"
	"github.com/MatheusMikio/Nexus/internal/domain/models"
	"github.com/MatheusMikio/Nexus/internal/repository"
)

type ILoginService interface {
	Login(email, password string) (*auth.LoginResponse, *models.ErrorMessage)
}

type LoginService struct {
	Repository repository.IUserRepository
}

func NewLoginService(repo repository.IUserRepository) ILoginService {
	return &LoginService{
		Repository: repo,
	}
}

func (l *LoginService) Login(email string, password string) (*auth.LoginResponse, *models.ErrorMessage) {
	userDb, err := l.Repository.GetByEmail(email)
	if err != nil {
		return nil, models.NewErrorMessage("Login", "Email or password invalid")
	}

	if !authLogin.VerifyPassword(password, userDb.GetPassword()) {
		return nil, models.NewErrorMessage("Login", "Email or password invalid")
	}

	accessToken, err := authLogin.GenerateAccessToken(userDb.PublicID, userDb.Role)
	if err != nil {
		return nil, models.NewErrorMessage("Token", "Falha ao gerar")
	}

	return &auth.LoginResponse{
		AccessToken: accessToken,
		ExpiresIn:   int64((2 * time.Hour).Seconds()),
		User: auth.AuthUser{
			Name:  userDb.GetName(),
			Email: userDb.GetEmail(),
			Phone: userDb.GetPhone(),
		},
	}, nil
}
