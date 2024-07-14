package repository

import (
	"gorm.io/gorm"
	"twitter-go-api/internal/entity"
)

type UserRepository interface {
	FindById(userID int) (entity.User, error)
	GetFollowers(userID int) ([]entity.User, error)
	Unfollow(userID int, followerUserID int) error
	GetFollowings(userID int) ([]entity.User, error)
	Follow(userID int) error
	RemoveFollowing(userID int) error
}

type userRepository struct {
	session *gorm.DB
}

func (u userRepository) FindById(userID int) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetFollowers(userID int) ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Unfollow(userID int, followerUserID int) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetFollowings(userID int) ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Follow(userID int) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) RemoveFollowing(userID int) error {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(session *gorm.DB) UserRepository {
	return &userRepository{
		session: session,
	}
}
