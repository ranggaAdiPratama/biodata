package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ranggaAdiPratama/go_biodata/token"
)

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
