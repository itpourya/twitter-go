package repository

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"twitter-go-api/internal/entity"
)

type AuthRepository interface {
	AddUser(user entity.User) (string, error)
	FindUserByEmail(email string) (entity.User, error)
}

type authRepository struct {
	session *gorm.DB
}

func NewAuthRepository(session *gorm.DB) AuthRepository {
	return &authRepository{
		session: session,
	}
}

func (a *authRepository) FindUserByEmail(email string) (entity.User, error) {
	var user entity.User

	query := a.session.Where("email == ?", email).Take(&user)
	if query.Error != nil {
		log.Println(query.Error)
	}

	return user, nil
}

func (a *authRepository) AddUser(user entity.User) (string, error) {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	}
	a.session.Save(&user)

	return "add user", nil
}

func hashAndSalt(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
