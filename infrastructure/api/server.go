package api

import (
	"markitos-golang-service-access/internal/domain/dependencies"

	"github.com/gin-gonic/gin"
)

type Server struct {
	address    string
	repository dependencies.UserRepository
	tokener    dependencies.Tokener
	hasher     dependencies.Hasher
	router     *gin.Engine
}

func (s *Server) Router() *gin.Engine {
	return s.router
}

func (s *Server) Repository() dependencies.UserRepository {
	return s.repository
}

func NewServer(
	address string,
	repository dependencies.UserRepository,
	tokener dependencies.Tokener,
	hasher dependencies.Hasher) *Server {

	server := &Server{
		address:    address,
		repository: repository,
		tokener:    tokener,
		hasher:     hasher,
	}

	server.router = server.createRouter()

	return server
}

func (s *Server) createRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/v1/users/motd", s.userMotdHandler)
	router.POST("/v1/users/register", s.userRegisterHandler)
	router.POST("/v1/users/login", s.userLoginHandler)

	protectedRoutes := router.Group("/").Use(bearerTokenMiddleware(s.tokener))
	protectedRoutes.PUT("/v1/users/me", s.userUpdateMeHandler)
	protectedRoutes.GET("/v1/users/me", s.userMeHandler)

	return router
}

func (s *Server) Run() error {
	return s.router.Run(s.address)
}

func errorResonses(err error) gin.H {
	return gin.H{"error": err.Error()}
}
