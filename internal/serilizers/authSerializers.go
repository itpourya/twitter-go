package serilizers

type RegisterRequest struct {
	Firstname string `json:"firstname" form:"firstname" binding:"required,min=1"`
	Lastname  string `json:"lastname" form:"lastname" binding:"required,min=1"`
	Email     string `json:"email" form:"email" binding:"required,email"`
	Password  string `json:"password" form:"password" binding:"required,min=4"`
	Username  string `json:"username" form:"username" binding:"required,min=1"`
}

type LoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required,min=4"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
