package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) listUserPosts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "list user posts",
	})
}

func (s *Server) detailUserPosts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "user post detail",
	})
}

func (s *Server) createUserPost(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "create post",
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
