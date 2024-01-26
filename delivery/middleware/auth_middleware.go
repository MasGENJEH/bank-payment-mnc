package middleware

import (
	"log"
	"net/http"
	"strings"
	"test-mnc/shared/service"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	RequireToken() gin.HandlerFunc
}

type authMiddleware struct {
	jwtService service.JwtService
}

type AuthHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func (a *authMiddleware) RequireToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var autHeader AuthHeader
		if err := ctx.ShouldBindHeader(&autHeader); err != nil {
			log.Printf("RequireToken.autHeader: %v \n", err.Error())
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenHeader := strings.Replace(autHeader.AuthorizationHeader, "Bearer ", "", -1)
		if tokenHeader == "" {
			log.Printf("RequireToken.tokenHeader \n")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, err := a.jwtService.ParseToken(tokenHeader)
		if err != nil {
			log.Printf("RequireToken.ParseToken: %v \n", err.Error())
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		
		ctx.Set("user", claims["userId"])

		ctx.Next()
	}
}

func NewAuthMiddleware(jwtService service.JwtService) AuthMiddleware {
	return &authMiddleware{jwtService: jwtService}
}