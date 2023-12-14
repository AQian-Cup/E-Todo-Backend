package user

type CreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
