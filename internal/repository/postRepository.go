package repository

import (
	"errors"
	"gorm.io/gorm"
	"twitter-go-api/internal/entity"
	"twitter-go-api/internal/serilizers"
)

type PostRepository interface {
	ListUserPosts(username string) ([]entity.Post, error)
	GetDetailPost(username string, postID int) (entity.Post, error)
	CreatePost(post entity.Post) error
	DeletePost(postID int, username string) error
	UpdatePost(postDetail serilizers.UpdatePostRequest, username string) error
}

type postRepository struct {
	session *gorm.DB
}

func (p postRepository) ListUserPosts(username string) ([]entity.Post, error) {
	var post []entity.Post

	//query := p.session.Where("author_username = ?", username).Find(&post)
	query := p.session.Preload("Bookmarks").Preload("User").Preload("Likes").Preload("Comments").Where("author_username = ?", username).Find(&post)
	//query := p.session.Model("Users").Where("author_username = ?", username).Find(&post)
	if query.Error != nil {
		return nil, errors.New("empty")
	}

	return post, nil
}

func (p postRepository) GetDetailPost(username string, postID int) (entity.Post, error) {
	var post []entity.Post

	query := p.session.Where("author_username = ?", username).Find(&post)
	if query.Error != nil {
		return entity.Post{}, errors.New("empty")
	}

	return post[postID-1], nil
}

func (p postRepository) CreatePost(post entity.Post) error {

	p.session.Save(&post)
	return nil
}

func (p postRepository) DeletePost(postID int, username string) error {
	var post entity.Post

	query := p.session.Model(&post).Where("id = ? AND author_username = ?", postID, username).Delete(&post)
	if query.Error != nil {
		return errors.New("can not delete post")
	}
	if query.RowsAffected < 1 {
		return errors.New("you don't have permission to delete")
	}

	return nil
}

func (p postRepository) UpdatePost(postDetail serilizers.UpdatePostRequest, username string) error {
	var post entity.Post

	query := p.session.Model(&post).Where("id = ? AND author_username = ?", postDetail.PostID, username).Update("content", postDetail.Content)
	if query.Error != nil {
		return errors.New("post not found")
	}

	return nil
}

func NewPostRepository(session *gorm.DB) PostRepository {
	return &postRepository{
		session: session,
	}
}
