package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"twitter-go-api/internal/entity"
	"twitter-go-api/internal/repository"
	"twitter-go-api/internal/serilizers"
)

type AuthService interface {
	AddUserService(request serilizers.RegisterRequest) (string, error)
	VerifyLogin(username string, password string) (entity.User, error)
	FindEmailService(username string) (string, error)
}

type authService struct {
	authRepository repository.AuthRepository
}

func (a authService) FindEmailService(username string) (string, error) {
	if username == "" {
		return "", errors.New("username is empty")
	}

	Username, err := a.authRepository.FindUserByUsername(username)
	if err != nil {
		return "", err
	}

	return Username.Email, nil
}

func (a authService) AddUserService(request serilizers.RegisterRequest) (string, error) {
	var user entity.User

	userExists, _ := a.authRepository.FindUserByEmail(request.Email)
	if userExists.Email != "" {
		return "", errors.New("user has already exist")
	}

	user.Firstname = request.Firstname
	user.Password = request.Password
	user.Lastname = request.Lastname
	user.Email = request.Email
	user.Username = request.Username
	user.IsActive = true

	status, err := a.authRepository.AddUser(user)
	if err != nil {
		return "", errors.New("can not add user")
	}

	return status, nil
}

func (a authService) VerifyLogin(username string, password string) (entity.User, error) {
	userExists, _ := a.authRepository.FindUserByUsername(username)

	if userExists.Email == "" {
		return entity.User{}, errors.New("email not found")
	}

	isValidPassword := comparePasswords(userExists.Password, []byte(password))
	if !isValidPassword {
		return entity.User{}, errors.New("failed to login, because password is not matched")
	}
	return userExists, nil
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{
		authRepository: repo,
	}
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
