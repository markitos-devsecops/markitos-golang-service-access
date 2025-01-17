package api

import (
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/domain/dependencies"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAuthenticatedUser(ctx *gin.Context) (string, error) {
	authPayload, exists := ctx.Get("auth_payload")
	if !exists || authPayload == nil {
		return "", domain.NewUnauthorizedError()
	}

	payload, ok := authPayload.(*dependencies.Payload)
	if !ok {
		return "", domain.NewUnauthorizedError()
	}

	return payload.User(), nil
}

func bearerTokenMiddleware(tokener dependencies.Tokener) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerHeader := ctx.GetHeader("Authorization")
		if len(bearerHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, domain.NewUnauthorizedError())
			return
		}

		bearerItems := strings.Fields(bearerHeader)
		if len(bearerItems) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, domain.NewUnauthorizedError())
			return
		}

		if bearerItems[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, domain.NewUnauthorizedError())
			return
		}

		token := bearerItems[1]
		payload, err := tokener.Validate(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, domain.NewUnauthorizedError())
			return
		}

		ctx.Set("auth_payload", payload)
		ctx.Next()
	}
}
