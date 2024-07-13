package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) getUserProfile(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "user profile",
	})
}

func (s *Server) deleteUserProfile(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "delete user profile",
	})
}

func (s *Server) getUserFollowers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "get user profile",
	})
}

func (s *Server) removeFromFollowers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "remove user profile",
	})
}

func (s *Server) getUserFollowings(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "user followings",
	})
}

func (s *Server) FollowUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "follow user",
	})
}

func (s *Server) UnfollowUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "unfollow user",
	})
}

func (s *Server) SignupUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "signup user",
	})
}

func (s *Server) LoginUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "login user",
	})
}
