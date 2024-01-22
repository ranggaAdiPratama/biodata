package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/ranggaAdiPratama/go_biodata/db/sqlc"
	"github.com/ranggaAdiPratama/go_biodata/token"
	"github.com/ranggaAdiPratama/go_biodata/util"
)

func (server *Server) index(ctx *gin.Context) {
	users, err := server.store.GetAllUser(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))

		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (server *Server) exporttoExcel(ctx *gin.Context) {
	entries, err := server.store.GetAllUser(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))

		return
	}

	users := []db.User{entries}

	userCount := util.CountUserDbStructs(users)

	fmt.Printf("users are %d", userCount)

	ctx.JSON(http.StatusOK, users)
}

func (server *Server) me(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	me, err := server.store.GetUser(ctx, authPayload.UserId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))

		return
	}

	rsp := userResponse{
		ID:        me.ID,
		Username:  me.Username,
		Name:      me.Name,
		Email:     me.Email,
		CreatedAt: me.CreatedAt,
	}

	ctx.JSON(http.StatusOK, rsp)
}
