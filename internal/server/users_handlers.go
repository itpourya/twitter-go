package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"twitter-go-api/internal/database"
	"twitter-go-api/internal/entity"
	jwt2 "twitter-go-api/internal/pkg/jwt"
	"twitter-go-api/internal/repository"
	"twitter-go-api/internal/serilizers"
	"twitter-go-api/internal/service"
)

var (
	DB             = database.New()
	authRepository = repository.NewAuthRepository(DB)
	authService    = service.NewAuthService(authRepository)
)

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
	var signupRequest serilizers.RegisterRequest

	err := ctx.ShouldBind(&signupRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	authService, err := authService.AddUserService(signupRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": authService,
	})

}

func (s *Server) LoginUser(ctx *gin.Context) {
	var loginRequest serilizers.LoginRequest
	var user entity.User
	err := ctx.ShouldBind(&loginRequest)
	if err != nil {
		return
	}

	_, err = authService.VerifyLogin(loginRequest.Username, loginRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userEmail, _ := authService.FindEmailService(loginRequest.Username)
	user.Username = loginRequest.Username
	user.Email = userEmail

	jwt := jwt2.Jwt{}
	token, _ := jwt.CreateToken(user)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
