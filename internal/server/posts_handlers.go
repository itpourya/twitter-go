package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"twitter-go-api/internal/repository"
	"twitter-go-api/internal/serilizers"
	"twitter-go-api/internal/service"
)

var (
	postRepository = repository.NewPostRepository(DB)
	postService    = service.NewPostService(postRepository)
)

func (s *Server) listUserPosts(ctx *gin.Context) {
	username := ctx.Param("username")

	if username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "username is empty",
		})
		return
	}

	posts, err := postService.GetListPost(username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user have not post",
		})
		return
	}

	ctx.JSON(http.StatusOK, posts)
}

func (s *Server) detailUserPosts(ctx *gin.Context) {
	username := ctx.Param("username")
	postId := ctx.Param("postId")
	convert, _ := strconv.Atoi(postId)

	post, err := postService.GetDetailPost(username, convert)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"ID":       post.ID,
		"content":  post.Content,
		"username": post.AuthorUsername,
		"like":     len(post.Likes),
		"comment":  len(post.Comments),
	})
}

func (s *Server) createUserPost(ctx *gin.Context) {
	var postRequest serilizers.CreatePostRequest

	err := ctx.ShouldBind(&postRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	username := ctx.Value("username")
	err = postService.CreatePost(postRequest, username.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "created",
	})
}

func (s *Server) updateUserPost(ctx *gin.Context) {
	var updatePostRequest serilizers.UpdatePostRequest
	username := ctx.Value("username")

	err := ctx.ShouldBind(&updatePostRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "enter valid post id or valid content",
		})
		return
	}

	err = postService.UpdatePost(updatePostRequest, username.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "you don't have permission to update",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "post updated",
	})

}

func (s *Server) removeUserPost(ctx *gin.Context) {
	postID, ok := strconv.Atoi(ctx.Param("postId"))
	if ok != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "enter valid post id",
		})
		return
	}
	username := ctx.Value("username")

	err := postService.DeletePost(postID, username.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "you don't have permission to delete",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "post deleted",
	})
}
