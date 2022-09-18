package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey = "Authorization"
	authorizationBearer    = "Bearer"
	authorizationPayload   = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHAiOiJSZWFjdGpzIiwibmFtZSI6Im1vcG9lbiIsInZlcnNpb24iOiJyZW1ha2UifQ.nOzZaARtqU0Oo4r-NLmze8Uubqmj8MqOdOcKIAsabW4"
)

func Bearear() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": errors.New("authorization required")})
			return
		}

		authorizationIssued := strings.Fields(authorizationHeader)
		if len(authorizationIssued) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": errors.New("invalid authorization header format")})
			return
		}

		if authorizationIssued[0] != authorizationBearer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": errors.New("invalid authorization")})
			return
		}

		if authorizationIssued[1] != authorizationPayload {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": errors.New("invalid credential")})
			return
		}

		ctx.Next()
	}
}
