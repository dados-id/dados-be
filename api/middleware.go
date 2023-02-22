package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/dados-id/dados-be/exception"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(firebaseClient auth.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, exception.ErrorResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, exception.ErrorResponse(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, exception.ErrorResponse(err))
			return
		}

		accessToken := fields[1]
		tokenInfo, err := firebaseClient.VerifyIDToken(ctx, accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, exception.ErrorResponse(err))
			return
		}

		userID := tokenInfo.Claims["user_id"].(string)

		ctx.Set(authorizationPayloadKey, userID)
		ctx.Next()
	}
}
