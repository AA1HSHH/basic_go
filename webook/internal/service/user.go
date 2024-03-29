package service

import (
	"basic-go/webook/internal/domain"
	"basic-go/webook/internal/repository"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

var (
	ErrUserDuplicateEmail    = repository.ErrUserDuplicateEmail
	ErrInvalidUserOrPassword = errors.New("用户或密码错误")
	ErrUserNotFound          = repository.ErrUserNotFound
)

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}
func (svc *UserService) Signup(ctx context.Context, u domain.User) error {
	encrypt, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(encrypt)
	return svc.repo.Create(ctx, u)
}

func (svc *UserService) Login(ctx context.Context, email string, password string) (domain.User, error) {
	u, err := svc.repo.FindByEmail(ctx, email)
	if errors.Is(err, repository.ErrUserNotFound) {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return u, nil
}

func (svc *UserService) Edit(ctx context.Context, u domain.User) error {
	return svc.repo.UpdateInfo(ctx, u)
}

func (svc *UserService) Profile(ctx *gin.Context, userId int64) (domain.User, error) {
	u, err := svc.repo.FindById(ctx, userId)
	if errors.Is(err, repository.ErrUserNotFound) {
		return domain.User{}, ErrUserNotFound
	}
	if err != nil {
		return domain.User{}, err
	}
	return u, err
}
