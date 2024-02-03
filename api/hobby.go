package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/ranggaAdiPratama/go_biodata/db/sqlc"
	"github.com/ranggaAdiPratama/go_biodata/token"
	"github.com/ranggaAdiPratama/go_biodata/util"
)

func (server *Server) myHobby(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	_, err := server.store.GetUser(ctx, authPayload.UserId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	entries, err := server.store.GetHobbyByUserId(ctx, authPayload.UserId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	if util.CountHobbyDbStructs(entries) > 0 {
		hobbies := make(map[int]map[string]interface{})

		for key, value := range entries {
			var updatedAt string

			if value.UpdatedAt.Valid {
				updatedAt = value.UpdatedAt.Time.Format("2006-01-02 15:04:05")
			} else {
				updatedAt = ""
			}

			hobbies[key] = map[string]interface{}{
				"id":         value.ID,
				"name":       value.Name,
				"created_at": value.CreatedAt.Format("2006-01-02 15:04:05"),
				"updated_at": updatedAt,
			}
		}

		rsp := ListResponse{
			Status:  http.StatusOK,
			Message: "No Hobby Found",
			Data:    hobbies,
		}

		ctx.JSON(http.StatusOK, rsp)

		return
	}

	rsp := noDataResponse{
		Status:  http.StatusOK,
		Message: "No Hobby Found",
	}

	ctx.JSON(http.StatusOK, rsp)

	return
}

func (server *Server) storeHobby(ctx *gin.Context) {
	var req hobbyRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))

		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	hobbies, err := server.store.GetHobby(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	if util.CountHobbyDbStructs(hobbies) > 0 {
		hobbyExists := false

		for _, value := range hobbies {
			if req.Name == value.Name && authPayload.UserId == value.UserID {
				hobbyExists = true
			}
		}

		if hobbyExists {
			ctx.JSON(http.StatusInternalServerError, errorResponsewithString("Hobby is registered already"))

			return
		}
	}

	arg := db.CreateHobbyParams{
		Name:   req.Name,
		UserID: authPayload.UserId,
	}

	newHobby, err := server.store.CreateHobby(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := createdHobbyResponse{
		Status:  http.StatusOK,
		Message: "Hobby created successfully",
		Data:    newHobbyResponse(newHobby),
	}

	ctx.JSON(http.StatusOK, rsp)

	return
}
