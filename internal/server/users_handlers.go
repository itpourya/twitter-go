package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"twitter-go-api/internal/database"
	jwt2 "twitter-go-api/internal/pkg/jwt"
	"twitter-go-api/internal/repository"
	"twitter-go-api/internal/serilizers"
	"twitter-go-api/internal/service"
)

var (
	DB             = database.New()
	authRepository = repository.NewAuthRepository(DB)
	authService    = service.NewAuthService(authRepository)
	userRepository = repository.NewUserRepository(DB)
	userService    = service.NewUserService(userRepository)
)

func (s *Server) getUserProfile(ctx *gin.Context) {
	username := ctx.Param("username")

	profile, countDetail, err := userService.GetProfile(username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}

	if profile.IsActive == false {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Account is disable or not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"firstname":  profile.Firstname,
		"lastname":   profile.Lastname,
		"username":   profile.Username,
		"email":      profile.Email,
		"IsActive":   profile.IsActive,
		"Posts":      countDetail["postsCount"],
		"Followers":  countDetail["followerCount"],
		"Followings": countDetail["followingCount"],
	})
}

func (s *Server) deleteUserProfile(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "delete user profile",
	})
}

func (s *Server) getUserFollowers(ctx *gin.Context) {
	userID := ctx.Param("userID")
	convert, _ := strconv.Atoi(userID)

	follower, err := userService.GetUserFollower(convert)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, follower)
}

func (s *Server) removeFromFollowers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "remove user profile",
	})
}

func (s *Server) getUserFollowings(ctx *gin.Context) {
	userID := ctx.Param("userID")
	convert, _ := strconv.Atoi(userID)

	followings, err := userService.GetUserFollowing(convert)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, followings)
}

func (s *Server) FollowUser(ctx *gin.Context) {
	user := ctx.Param("user_id")
	userID, _ := strconv.Atoi(user)
	followerID := ctx.Value("user_id")

	err := userService.FollowUser(followerID.(int), userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "can not follow",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "followed",
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

	err := ctx.ShouldBind(&loginRequest)
	if err != nil {
		return
	}

	user, err := authService.VerifyLogin(loginRequest.Username, loginRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	jwt := jwt2.Jwt{}
	token, _ := jwt.CreateToken(user)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
