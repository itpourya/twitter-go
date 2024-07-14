package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"twitter-go-api/internal/repository"
	"twitter-go-api/internal/serilizers"
	"twitter-go-api/internal/service"
)

var (
	postRepository repository.PostRepository = repository.NewPostRepository(DB)
	postService    service.PostService       = service.NewPostService(postRepository)
)

func (s *Server) listUserPosts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "list user posts",
	})
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

	log.Println(post)

	ctx.JSON(http.StatusOK, gin.H{
		"ID":       post.ID,
		"content":  post.Content,
		"username": post.AuthorEmail,
		"like":     post.Likes,
		"comment":  post.Comments,
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
	userEmail := ctx.Value("userEmail")
	err = postService.CreatePost(postRequest, userEmail.(string))
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
	ctx.JSON(http.StatusOK, gin.H{
		"data": "update user post",
	})
}

func (s *Server) removeUserPost(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "delete user post",
	})
}
