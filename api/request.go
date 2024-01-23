package api

// SECTION auth
type loginRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type refreshRequest struct {
	Token string `form:"token" binding:"required"`
}

type registerRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

// !SECTION auth
// SECTION user
// !SECTION user
