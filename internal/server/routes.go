package server

import (
	"net/http"
	"twitter-go-api/internal/middleware"
	jwt2 "twitter-go-api/internal/pkg/jwt"

	"github.com/gin-gonic/gin"
)

var (
	jwt jwt2.Jwt
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.HandleMethodNotAllowed = true

	users := r.Group("/profile", middleware.AthorizationJWT(jwt))
	{
		users.GET("/:userId", s.getUserProfile)
		users.GET("/:userId/followers", s.getUserFollowers)
		users.DELETE("/:userId/followers/:follower_userId", s.UnfollowUser)
		users.GET("/:userId/followings", s.getUserFollowings)
		users.POST("/:userId/followings", s.FollowUser)
		users.DELETE("/:userId/followings/:followingUserId", s.UnfollowUser)
	}

	auth := r.Group("/api/v1")
	{
		auth.POST("/create", s.SignupUser)
		auth.POST("/login", s.LoginUser)
	}

	posts := r.Group("/posts", middleware.AthorizationJWT(jwt))
	{
		posts.GET("/:username", s.listUserPosts)
		posts.GET("/:username/:postId", s.detailUserPosts)
		posts.POST("/create-post", s.createUserPost)
		posts.DELETE("/:postId", s.removeUserPost)
		posts.PUT("/update", s.updateUserPost)
	}

	return r
}
