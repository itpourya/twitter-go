package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"twitter-go-api/internal/entity"
)

type UserRepository interface {
	FindById(userID int) (entity.User, error)
	GetFollowers(userID int) ([]entity.Follower, error)
	Unfollow(userID int, followerUserID int) error
	GetFollowings(userID int) ([]entity.User, error)
	Follow(followerUserID int, userID int) error
	RemoveFollowing(userID int) error
	GetProfile(username string) (entity.User, map[string]int64, error)
}

type userRepository struct {
	session *gorm.DB
}

func (u userRepository) GetProfile(username string) (entity.User, map[string]int64, error) {
	var user entity.User
	var posts []entity.Post
	var postCount int64
	var countDetail = make(map[string]int64)

	query := u.session.Model(&user).Where("username = ?", username).Find(&user)
	if query.Error != nil {
		return entity.User{}, countDetail, errors.New("user not found")
	}

	query = u.session.Model(&posts).Where("author_username = ?", username).Find(&posts).Count(&postCount)
	followerCount, err := u.getFollowerCount(user.ID)
	if err != nil {
		return entity.User{}, countDetail, err
	}
	followingCount, err := u.getFollowingCount(user.ID)
	if err != nil {
		return entity.User{}, countDetail, err
	}

	countDetail["postsCount"] = postCount
	countDetail["followerCount"] = followerCount
	countDetail["followingCount"] = followingCount

	return user, countDetail, nil
}

func (u userRepository) FindById(userID int) (entity.User, error) {
	var user entity.User

	query := u.session.Model(&user).Where("id = ?", userID).Take(&user)
	if query.Error != nil {
		return entity.User{}, errors.New("user not found")
	}

	return user, nil
}

func (u userRepository) GetFollowers(userID int) ([]entity.Follower, error) {
	var follower []entity.Follower

	query := u.session.Model("Follower").Where("user_id", userID).Find(&follower)
	if query.Error != nil {
		return nil, errors.New("users have not follower")
	}

	fmt.Println(follower)

	return follower, nil
}

func (u userRepository) Unfollow(userID int, followerUserID int) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetFollowings(userID int) ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Follow(followerUserID int, userID int) error {
	var follower entity.Follower

	check := !u.checkUserFollowings(userID, followerUserID)
	if check == false {
		return errors.New("error")
	}

	follower.FollowerUserID = followerUserID
	follower.UserID = userID

	u.session.Create(&follower)
	return nil
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

func (u userRepository) checkUserFollowings(userID int, followerID int) bool {
	var follower entity.Follower

	query := u.session.Model(&follower).Where("follower_user_id = ? AND user_id = ?", followerID, userID).Take(&follower)
	if query.Error != nil {
		return false
	}

	return true
}

func (u userRepository) getFollowerCount(userID int) (int64, error) {
	var followerCount int64

	query := u.session.Model(&entity.Follower{}).Where("user_id == ?", userID).Count(&followerCount)
	if query.Error != nil {
		return 0, errors.New("empty")
	}

	return followerCount, nil
}

func (u userRepository) getUserByUsername(username string) (entity.User, error) {
	var user entity.User

	query := u.session.Where("username = ?", username).Take(&user)
	if query.Error != nil {
		return entity.User{}, errors.New("not found")
	}

	return user, nil
}

func (u userRepository) getFollowingCount(userID int) (int64, error) {
	var followingCount int64

	query := u.session.Model(&entity.Following{}).Where("user_id == ?", userID).Count(&followingCount)
	if query.Error != nil {
		return 0, errors.New("empty")
	}

	return followingCount, nil
}
