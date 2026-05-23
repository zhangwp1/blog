package service

import (
	"blog/internal/dto"
	"blog/internal/model"
	"blog/internal/repository"
	"blog/internal/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return nil, errors.New("生成token失败")
	}

	return &dto.LoginResponse{
		Token:     token,
		ExpiresIn: 86400,
	}, nil
}

func (s *AuthService) GetProfile(userID uint) (*dto.UserProfile, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	return &dto.UserProfile{
		ID:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

var _ = model.User{}
