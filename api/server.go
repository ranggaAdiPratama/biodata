package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/ranggaAdiPratama/go_biodata/db/sqlc"
	"github.com/ranggaAdiPratama/go_biodata/token"
	"github.com/ranggaAdiPratama/go_biodata/util"
)

// Server serves HTTP requests for our banking service
type Server struct {
	config     util.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymetricKey)

	if err != nil {
		return nil, err
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("currency", validCurrency)
	// }

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	authRoutes := router.Group("/api").Use(authMiddleware(server.tokenMaker))

	// router.GET("/public/:dir/*asset", func(c *gin.Context) {
	// 	dir := c.Param("dir")
	// 	asset := c.Param("asset")

	// 	if strings.TrimPrefix(asset, "/") == "" {
	// 		c.AbortWithStatus(http.StatusNotFound)
	// 		return
	// 	}

	// 	fullName := filepath.Join(dir, filepath.FromSlash(path.Clean("/"+asset)))

	// 	c.File(fullName)
	// })

	router.Static("/public", "./public")

	authRoutes.GET("/me", server.me)
	authRoutes.GET("/export/user", server.exportUsertoExcel)
	authRoutes.GET("/my-hobby", server.myHobby)
	authRoutes.GET("/users", server.userList)
	authRoutes.GET("/export/hobby", server.exportHobbytoExcel)
	authRoutes.POST("/hobby", server.storeHobby)
	authRoutes.POST("/profile", server.updateProfile)
	// authRoutes.GET("/accounts/:id", server.getAccount)

	apiRoutes := router.Group("/api")

	apiRoutes.POST("/auth/login", server.login)
	apiRoutes.POST("/auth/refresh", server.refreshToken)
	apiRoutes.POST("/auth/register", server.register)

	server.router = router
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponsewithString(msg string) errorDataResponse {
	return errorDataResponse{
		Status:  http.StatusInternalServerError,
		Message: "Error",
		Data:    msg,
	}
}

func errorResponse(err error) errorDataResponse {
	return errorDataResponse{
		Status:  http.StatusInternalServerError,
		Message: "Error",
		Data:    err.Error(),
	}
}
