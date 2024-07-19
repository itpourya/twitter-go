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

	users := r.Group("/users", middleware.Authorization(jwt))
	{
		users.GET("/:username", s.getUserProfile) // DONE
		users.GET("/show-followers/:userID", s.getUserFollowers)
		users.DELETE("/unfollow/:userID", s.UnfollowUser)
		users.GET("/show-followings/:userID", s.getUserFollowings)
		users.GET("/follow/:user_id", s.FollowUser) // DONE
		users.DELETE("/remove/:userID", s.removeFromFollowers)
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
		posts.PUT("/update", s.updateUserPost)             // DONE
	}

	return r
}
