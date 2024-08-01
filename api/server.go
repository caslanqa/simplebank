package api

import (
	"fmt"
	db "github.com/CharlieAlphaQA/simplebank/db/sqlc"
	"github.com/CharlieAlphaQA/simplebank/token"
	"github.com/CharlieAlphaQA/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
	config     util.Config
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker : %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("currency", validCurrency)
		if err != nil {
			return nil, fmt.Errorf("cannot register validation: %w", err)
		}
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authGroup := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authGroup.POST("/accounts", server.createAccount)
	authGroup.GET("/accounts/:id", server.getAccount)
	authGroup.GET("/accounts", server.listAccounts)
	authGroup.DELETE("/accounts/:id", server.deleteAccount)
	authGroup.POST("/transfers", server.createTransfer)

	server.router = router
}
