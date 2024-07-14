package repository

import (
	"errors"
	"gorm.io/gorm"
	"twitter-go-api/internal/entity"
)

type PostRepository interface {
	ListUserPosts(username string) ([]entity.Post, error)
	GetDetailPost(username string, postID int) (entity.Post, error)
	CreatePost(post entity.Post) error
	DeletePost(postID int) error
	UpdatePost(post entity.Post) error
}

type postRepository struct {
	session *gorm.DB
}

func (p postRepository) ListUserPosts(username string) ([]entity.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p postRepository) GetDetailPost(username string, postID int) (entity.Post, error) {
	var post entity.Post

	query := p.session.Where("ID = ?", postID, username).Take(&post)
	if query.Error != nil {
		return entity.Post{}, errors.New("empty")
	}

	return post, nil
}

func (p postRepository) CreatePost(post entity.Post) error {

	p.session.Save(&post)
	return nil
}

func (p postRepository) DeletePost(postID int) error {
	//TODO implement me
	panic("implement me")
}

func (p postRepository) UpdatePost(post entity.Post) error {
	//TODO implement me
	panic("implement me")
}

func NewPostRepository(session *gorm.DB) PostRepository {
	return &postRepository{
		session: session,
	}
}
