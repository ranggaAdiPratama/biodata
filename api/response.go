package api

import (
	"time"

	db "github.com/ranggaAdiPratama/go_biodata/db/sqlc"
)

// SECTION index
type errorDataResponse struct {
	Status  int64  `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// !SECTION index
// SECTION auth
type loginDataResponse struct {
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
	User                  userResponse `json:"user"`
}

type loginResponse struct {
	Status  int64             `json:"status"`
	Message string            `json:"message"`
	Data    loginDataResponse `json:"data"`
}

type registerResponse struct {
	Status  int64        `json:"status"`
	Message string       `json:"message"`
	Data    userResponse `json:"data"`
}

type userResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		ID:        user.ID,
		Username:  user.Username,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

// !SECTION auth
// SECTION user
type meResponse struct {
	Status  int64              `json:"status"`
	Message string             `json:"message"`
	Data    userDetailResponse `json:"data"`
}

type profileResponse struct {
	Status  int64                 `json:"status"`
	Message string                `json:"message"`
	Data    userDetailAllResponse `json:"data"`
}

type userDetailResponse struct {
	ID             int64     `json:"id"`
	Username       string    `json:"username"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	ProfilePicture string    `json:"profile_picture"`
	CreatedAt      time.Time `json:"created_at"`
}

type userDetailAllResponse struct {
	ID             int64     `json:"id"`
	Username       string    `json:"username"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	ProfilePicture string    `json:"profile_picture"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      string    `json:"updated_at"`
}

type userListResponse struct {
	Status  int64                          `json:"status"`
	Message string                         `json:"message"`
	Data    map[int]map[string]interface{} `json:"data"`
}

func newUserDetailResponse(user db.User) userDetailResponse {
	var profilePicture string

	if user.ProfilePicture.Valid {
		profilePicture = user.ProfilePicture.String
	} else {
		profilePicture = ""
	}

	return userDetailResponse{
		ID:             user.ID,
		Username:       user.Username,
		Name:           user.Name,
		Email:          user.Email,
		ProfilePicture: profilePicture,
		CreatedAt:      user.CreatedAt,
	}
}

func UserDetailAllResponse(user db.User) userDetailAllResponse {
	var profilePicture string
	var updatedAt string

	if user.ProfilePicture.Valid {
		profilePicture = user.ProfilePicture.String
	} else {
		profilePicture = ""
	}

	if user.UpdatedAt.Valid {
		updatedAt = user.UpdatedAt.Time.Format("2024-01-21T13:46:27.463394Z")
	} else {
		updatedAt = ""
	}

	return userDetailAllResponse{
		ID:             user.ID,
		Username:       user.Username,
		Name:           user.Name,
		Email:          user.Email,
		ProfilePicture: profilePicture,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      updatedAt,
	}
}

// !SECTION user
