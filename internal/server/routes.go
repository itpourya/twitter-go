package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/", s.HelloWorldHandler)

	r.GET("/health", s.healthHandler)

	users := r.Group("/profile")
	{
		users.GET("/:user_id", s.getUserProfile)
		users.GET("/:user_id/followers", s.getUserFollowers)
		users.DELETE("/:user_+id/followers/:follower_user_id", s.UnfollowUser)
		users.GET("/:user_id/followings", s.getUserFollowings)
		users.POST("/:user_id/followings", s.FollowUser)
		users.DELETE("/:user_id/followings/:following_user_id", s.UnfollowUser)
	}

	auth := r.Group("/api/v1")
	{
		auth.POST("/create", s.SignupUser)
		auth.POST("/login", s.LoginUser)
	}

	posts := r.Group("/posts")
	{
		posts.GET("/:username", s.listUserPosts)
		posts.GET("/:username/:post_id", s.detailUserPosts)
		posts.POST("/create-post", s.createUserPost)
		posts.DELETE("/:post_id", s.removeUserPost)
		posts.PUT("/update", s.updateUserPost)
	}

	return r
}
