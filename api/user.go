package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	db "github.com/ranggaAdiPratama/go_biodata/db/sqlc"
	"github.com/ranggaAdiPratama/go_biodata/token"
	"github.com/ranggaAdiPratama/go_biodata/util"
)

func (server *Server) userList(ctx *gin.Context) {
	entries, err := server.store.GetAllUser(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))

		return
	}

	users := make(map[int]map[string]interface{})

	for key, value := range entries {
		var profilePicture string
		var updatedAt string

		if value.ProfilePicture.Valid {
			profilePicture = "/public/images/users/" + value.ProfilePicture.String
		} else {
			profilePicture = ""
		}

		if value.UpdatedAt.Valid {
			updatedAt = value.UpdatedAt.Time.Format("2006-01-02 15:04:05")
		} else {
			updatedAt = ""
		}

		users[key] = map[string]interface{}{
			"id":              value.ID,
			"username":        value.Username,
			"name":            value.Name,
			"email":           value.Email,
			"profile_picture": profilePicture,
			"created_at":      value.CreatedAt.Format("2006-01-02 15:04:05"),
			"updated_at":      updatedAt,
		}
	}

	rsp := ListResponse{
		Status:  http.StatusOK,
		Message: "User list retrieved",
		Data:    users,
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) exportUsertoExcel(ctx *gin.Context) {
	entries, err := server.store.GetAllUser(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))

		return
	}

	userCount := util.CountUserDbStructs(entries)

	fmt.Printf("users are %d", userCount)

	xlsx := excelize.NewFile()

	sheet1Name := "Sheet One"

	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	xlsx.SetCellValue(sheet1Name, "A1", "Name")
	xlsx.SetCellValue(sheet1Name, "B1", "Username")
	xlsx.SetCellValue(sheet1Name, "C1", "Email")

	err = xlsx.AutoFilter(sheet1Name, "A1", "C1", "")

	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	for i, value := range entries {
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+2), value.Name)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+2), value.Username)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+2), value.Email)
	}

	err = xlsx.SaveAs("public/xlsxs/file1.xlsx")

	if err != nil {
		fmt.Println(err)
	}

	rsp := userExportResponse{
		Status:  http.StatusOK,
		Message: "Export success",
		Data:    "/public/xlsxs/file1.xlsx",
	}

	ctx.JSON(http.StatusOK, rsp)
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

func (server *Server) updateProfile(ctx *gin.Context) {
	var req updateProfileRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	users, err := server.store.GetAllUser(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))

		return
	}

	me, err := server.store.GetUser(ctx, authPayload.UserId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))

		return
	}

	if req.Name == "" {
		req.Name = me.Name
	}

	if req.Email == "" {
		req.Email = me.Email
	}

	if req.Username == "" {
		req.Username = me.Username
	}

	userEmailExists := false
	userNameExists := false

	for _, value := range users {
		if req.Username != me.Username {
			if value.Username == req.Username && authPayload.UserId != value.ID {
				userNameExists = true
			}
		}

		if req.Email != me.Email {
			if value.Email == req.Email && authPayload.UserId != value.ID {
				userEmailExists = true
			}
		}
	}

	if userNameExists {
		ctx.JSON(http.StatusInternalServerError, errorResponsewithString("Username is registered already"))

		return
	}

	if userEmailExists {
		ctx.JSON(http.StatusInternalServerError, errorResponsewithString("Email is registered already"))

		return
	}

	file, err := ctx.FormFile("profile_picture")

	if err != nil {
		profilePicture := me.ProfilePicture

		arg := db.UpdateUserParams{
			ID:             me.ID,
			Name:           req.Name,
			Username:       req.Username,
			Email:          req.Email,
			ProfilePicture: profilePicture,
			UpdatedAt: sql.NullTime{
				Valid: true,
				Time:  time.Now(),
			},
		}

		me, err = server.store.UpdateUser(ctx, arg)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))

			return
		}

		rsp := profileResponse{
			Status:  http.StatusOK,
			Message: "Profile updated successfully",
			Data:    UserDetailAllResponse(me),
		}

		ctx.JSON(http.StatusOK, rsp)

		return
	}

	ext := util.GetFileExtension(file.Filename)

	fileName := authPayload.Username + "_" + time.Now().Format("20060102150405") + ext

	if me.ProfilePicture.Valid {
		oldFilePath := "public/images/users/" + me.ProfilePicture.String

		if util.FileExists(oldFilePath) {
			err = util.DeleteFile(oldFilePath)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))

				return
			}
		}
	}

	path := "public/images/users/" + fileName

	err = ctx.SaveUploadedFile(file, path)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, errorResponse(err))
		return
	}

	arg := db.UpdateUserParams{
		ID:       me.ID,
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		ProfilePicture: sql.NullString{
			Valid:  true,
			String: fileName,
		},
		UpdatedAt: sql.NullTime{
			Valid: true,
			Time:  time.Now(),
		},
	}

	me, err = server.store.UpdateUser(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))

		return
	}

	rsp := profileResponse{
		Status:  http.StatusOK,
		Message: "Profile updated successfully",
		Data:    UserDetailAllResponse(me),
	}

	ctx.JSON(http.StatusOK, rsp)
}
