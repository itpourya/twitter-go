package serilizers

type CreatePostRequest struct {
	Content string `json:"content" form:"content" binding:"required,min=1"`
}
