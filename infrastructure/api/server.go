package api

import (
	"markitos-golang-service-access/internal/domain"

	"github.com/gin-gonic/gin"
)

type Server struct {
	address    string
	repository domain.UserRepository
	router     *gin.Engine
}

func (s *Server) Router() *gin.Engine {
	return s.router
}

func (s *Server) Repository() domain.UserRepository {
	return s.repository
}

func NewServer(address string, repository domain.UserRepository) *Server {
	server := &Server{
		address:    address,
		repository: repository,
	}
	server.router = server.createRouter()

	return server
}

func (s *Server) createRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/v1/users", s.userCreateHandler)
	router.GET("/v1/users/all", s.userListHandler)
	router.GET("/v1/users/:id", s.userOneHandler)
	router.PUT("/v1/users/:id", s.userUpdateHandler)
	router.GET("/v1/", s.userMotdHandler)
	router.GET("/v1/users", s.userSearchHandler)

	return router
}

func (s *Server) Run() error {
	return s.router.Run(s.address)
}

func errorResonses(err error) gin.H {
	return gin.H{"error": err.Error()}
}
