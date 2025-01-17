package api

import (
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/services"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Server) userRegisterHandler(ctx *gin.Context) {
	var request services.UserRegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	var service services.UserRegisterService = services.NewUserRegisterService(s.repository, s.hasher)
	user, err := service.Execute(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (s *Server) userMeHandler(ctx *gin.Context) {
	userId, err := GetAuthenticatedUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResonses(err))
		return
	}
	securedId, err := domain.NewUserId(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	var request services.UserMeRequest = services.NewUserMeRequest(securedId.Value())
	var service services.UserMeService = services.NewUserMeService(s.repository)
	user, err := service.Execute(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (s *Server) userUpdateMeHandler(ctx *gin.Context) {
	request, shouldExitByError := createUpdateRequestOrExitWithError(ctx)
	if shouldExitByError {
		return
	}

	var service services.UserUpdateMeService = services.NewUserUpdateMeService(s.repository)
	user, err := service.Execute(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func createUpdateRequestOrExitWithError(ctx *gin.Context) (services.UserUpdateMeRequest, bool) {
	userId, err := GetAuthenticatedUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResonses(err))
		return services.UserUpdateMeRequest{}, true
	}

	var requestBody services.UserUpdateMeRequestBody
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return services.UserUpdateMeRequest{}, true
	}

	var request services.UserUpdateMeRequest = services.UserUpdateMeRequest{
		Id:   userId,
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
