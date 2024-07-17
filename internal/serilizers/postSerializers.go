package serilizers

import "twitter-go-api/internal/entity"

type CreatePostRequest struct {
	Content string `json:"content" form:"content" binding:"required,min=1"`
}

type UpdatePostRequest struct {
	PostID  int    `json:"postId" form:"postId" binding:"required,min=1"`
	Content string `json:"content" form:"content" binding:"required,min=1"`
}

type PostResponse struct {
	ID        int              `json:"id" form:"id" binding:"required,min=1"`
	Content   string           `json:"content" form:"content" binding:"required,min=1"`
	User      *entity.User     `json:"user" form:"user" binding:"required,min=1"`
	Bookmarks []entity.User    `json:"bookmarks" form:"bookmarks" binding:"required,min=1"`
	Likes     []entity.User    `json:"likes" form:"likes" binding:"required,min=1"`
	Comments  []entity.Comment `json:"comments" form:"comments" binding:"required,min=1"`
}
