package service

import (
	"twitter-go-api/internal/entity"
	"twitter-go-api/internal/repository"
)

type UserService interface {
	GetProfile(username string) (entity.User, map[string]int64, error)
	GetUserFollower(userID int) ([]entity.Follower, error)
	FollowUser(followerUserID int, userID int) error
	UnfollowUser(username string, userID int) error
}

type userService struct {
	userRepository repository.UserRepository
}

func (u userService) GetProfile(username string) (entity.User, map[string]int64, error) {
	profile, countDetail, err := u.userRepository.GetProfile(username)
	if err != nil {
		return entity.User{}, countDetail, err
	}

	return profile, countDetail, nil
}

func (u userService) GetUserFollower(userID int) ([]entity.Follower, error) {
	followers, err := u.userRepository.GetFollowers(userID)
	if err != nil {
		return nil, err
	}

	return followers, nil
}

func (u userService) FollowUser(followerUserID int, userID int) error {
	err := u.userRepository.Follow(followerUserID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (u userService) UnfollowUser(username string, userID int) error {
	//TODO implement me
	panic("implement me")
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		userRepository: repo,
	}
}
