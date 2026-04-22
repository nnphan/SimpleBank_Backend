package api

import (
	db "simplebank/internal/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store *db.SQLStore
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.SQLStore) (*Server, error) {
	server := &Server{store: store}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	//server.router = router.NewRouter()
	router := gin.Default()
	router.POST("/users", server.createUser)
	router.GET("/listusers", server.listUsers)

	router.POST("/accounts", server.createAccount)
	router.POST("/addaccountbalance", server.addAccountBalance)

	router.POST("/transfers", server.createTransfer)
	
	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}