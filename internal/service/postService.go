package service

import (
	"errors"
	"twitter-go-api/internal/entity"
	"twitter-go-api/internal/repository"
	"twitter-go-api/internal/serilizers"
)

type PostService interface {
	CreatePost(request serilizers.CreatePostRequest, email string) error
	GetListPost(userID int) ([]entity.Post, error)
	GetDetailPost(username string, postID int) (entity.Post, error)
	DeletePost(postID int) error
	UpdatePost(post entity.Post) error
}

type postService struct {
	postRepository repository.PostRepository
}

func (p postService) CreatePost(request serilizers.CreatePostRequest, email string) error {
	var post entity.Post
	if request.Content == "" {
		return errors.New("empty")
	}

	post.AuthorEmail = email
	post.Content = request.Content

	err := p.postRepository.CreatePost(post)
	if err != nil {
		return errors.New("cant create post")
	}

	return nil
}

func (p postService) GetListPost(userID int) ([]entity.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p postService) GetDetailPost(username string, postID int) (entity.Post, error) {
	post, err := p.postRepository.GetDetailPost(username, postID)
	if err != nil {
		return entity.Post{}, errors.New("empty")
	}

	return post, nil
}

func (p postService) DeletePost(postID int) error {
	//TODO implement me
	panic("implement me")
}

func (p postService) UpdatePost(post entity.Post) error {
	//TODO implement me
	panic("implement me")
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{
		postRepository: repo,
	}
}
