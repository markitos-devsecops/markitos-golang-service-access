package api

import (
	"log"
	"markitos-golang-service-access/internal/services"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Server) userCreateHandler(ctx *gin.Context) {
	var request services.UserCreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	var service services.UserCreateService = services.NewUserCreateService(s.repository, s.hasher)
	user, err := service.Execute(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (s *Server) userOneHandler(ctx *gin.Context) {
	var request services.UserOneRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	var service services.UserOneService = services.NewUserOneService(s.repository)
	user, err := service.Execute(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (s *Server) userUpdateHandler(ctx *gin.Context) {
	request, shouldExitByError := createUpdateRequestOrExitWithError(ctx)
	if shouldExitByError {
		return
	}

	var service services.UserUpdateService = services.NewUserUpdateService(s.repository)
	user, err := service.Execute(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func createUpdateRequestOrExitWithError(ctx *gin.Context) (services.UserUpdateRequest, bool) {
	var requestUri services.UserUpdateRequestUri
	if err := ctx.ShouldBindUri(&requestUri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return services.UserUpdateRequest{}, true
	}
	var requestBody services.UserUpdateRequestBody
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return services.UserUpdateRequest{}, true
	}

	var request services.UserUpdateRequest = services.UserUpdateRequest{
		Id:   requestUri.Id,
		Name: requestBody.Name,
	}

	return request, false
}

func (s *Server) userLoginHandler(ctx *gin.Context) {
	var request services.UserLoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	var service services.UserLoginService = services.NewUserLoginService(s.repository, s.hasher, s.tokener)
	log.Println("request", request)
	user, err := service.Execute(request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResonses(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (s *Server) userMotdHandler(ctx *gin.Context) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	ctx.JSON(http.StatusOK, "[ACCESS] - Marco Antonio - markitos say, Hi all!! at "+time.Now().Format(time.RFC3339)+" on "+hostname)
}
