package dto

type LoginDTO struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type RegisterDTO struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required"`
	AvatarUrl   string `json:"avatarUrl" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Description string `json:"description" binding:"required"`
	Gender      string `json:"gender" binding:"required,oneof=MALE FEMALE"`
}

type LoginResponseDTO struct {
	Token     string `json:"token"`
	Username  string `json:"username"`
	AvatarUrl string `json:"avatarUrl"`
}

type RegisterResponseDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
