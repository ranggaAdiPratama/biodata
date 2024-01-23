package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ranggaAdiPratama/go_biodata/token"
	"github.com/ranggaAdiPratama/go_biodata/util"
)

func (server *Server) index(ctx *gin.Context) {
	entries, err := server.store.GetAllUser(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))

		return
	}

	users := make(map[int]map[string]interface{})

	for key, value := range entries {
		var profilePicture string

		if value.ProfilePicture.Valid {
			profilePicture = value.ProfilePicture.String
		} else {
			profilePicture = ""
		}

		users[key] = map[string]interface{}{
			"id":              value.ID,
			"username":        value.Username,
			"name":            value.Name,
			"email":           value.Email,
			"profile_picture": profilePicture,
			"created_at":      value.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	rsp := userListResponse{
		Status:  http.StatusOK,
		Message: "User list retrieved",
		Data:    users,
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) exporttoExcel(ctx *gin.Context) {
	entries, err := server.store.GetAllUser(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))

		return
	}

	userCount := util.CountUserDbStructs(entries)

	fmt.Printf("users are %d", userCount)

	ctx.JSON(http.StatusOK, entries)
}

func (server *Server) me(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	me, err := server.store.GetUser(ctx, authPayload.UserId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))

		return
	}

	rsp := meResponse{
		Status:  http.StatusOK,
		Message: "User list retrieved",
		Data:    newUserDetailResponse(me),
	}

	ctx.JSON(http.StatusOK, rsp)
}
