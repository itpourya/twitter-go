package service

import (
	"errors"
	"twitter-go-api/internal/entity"
	"twitter-go-api/internal/repository"
	"twitter-go-api/internal/serilizers"
)

type PostService interface {
	CreatePost(request serilizers.CreatePostRequest, username string) error
	GetListPost(username string) ([]entity.Post, error)
	GetDetailPost(username string, postID int) (entity.Post, error)
	DeletePost(postID int, username string) error
	UpdatePost(postDetail serilizers.UpdatePostRequest, username string) error
}

type postService struct {
	postRepository repository.PostRepository
}

func (p postService) CreatePost(request serilizers.CreatePostRequest, username string) error {
	var post entity.Post
	if request.Content == "" {
		return errors.New("empty")
	}

	post.AuthorUsername = username
	post.Content = request.Content

	err := p.postRepository.CreatePost(post)
	if err != nil {
		return errors.New("cant create post")
	}

	return nil
}

func (p postService) GetListPost(username string) ([]entity.Post, error) {
	posts, err := p.postRepository.ListUserPosts(username)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (p postService) GetDetailPost(username string, postID int) (entity.Post, error) {
	post, err := p.postRepository.GetDetailPost(username, postID)
	if err != nil {
		return entity.Post{}, errors.New("empty")
	}

	return post, nil
}

func (p postService) DeletePost(postID int, username string) error {
	ok := p.postRepository.DeletePost(postID, username)
	if ok != nil {
		return ok
	}

	return nil
}

func (p postService) UpdatePost(postDetail serilizers.UpdatePostRequest, username string) error {
	err := p.postRepository.UpdatePost(postDetail, username)
	if err != nil {
		return err
	}

	return nil
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{
		postRepository: repo,
	}
}
