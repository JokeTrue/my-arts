package http

type SignUpRequest struct {
	Email     string `json:"email" binding:"required"`
	Password1 string `json:"password1" binding:"required"`
	Password2 string `json:"password2" binding:"required"`

	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Age       int    `json:"age" binding:"required"`
	Gender    string `json:"gender" binding:"required"`
	Location  string `json:"location" binding:"required"`
	Biography string `json:"biography" binding:"required"`
}
