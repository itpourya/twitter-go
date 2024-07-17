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

	users := r.Group("/profile", middleware.Authorization(jwt))
	{
		users.GET("/:username", s.getUserProfile)
		users.GET("/:username/followers", s.getUserFollowers)
		users.DELETE("/:username/followers/:follower_userId", s.UnfollowUser)
		users.GET("/:username/followings", s.getUserFollowings)
		users.POST("/:username/followings", s.FollowUser)
		users.DELETE("/:username/followings/:followingUsername", s.removeFromFollowers)
	}

	auth := r.Group("/api/v1") // DONE
	{
		auth.POST("/create", s.SignupUser) // DONE
		auth.POST("/login", s.LoginUser)   // DONE
	}

	posts := r.Group("/posts", middleware.Authorization(jwt))
	{
		posts.GET("/:username", s.listUserPosts)           // DONE
		posts.GET("/:username/:postId", s.detailUserPosts) // DONE
		posts.POST("/create-post", s.createUserPost)       // DONE
		posts.DELETE("/:postId", s.removeUserPost)         // DONE
		posts.PUT("/update", s.updateUserPost)
	}

	return r
}
