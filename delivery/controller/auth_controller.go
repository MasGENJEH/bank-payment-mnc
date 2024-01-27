package controller

import (
	"net/http"
	"test-mnc/config"
	"test-mnc/delivery/middleware"
	"test-mnc/entity/dto"
	"test-mnc/shared/common"
	"test-mnc/usecase"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUc usecase.AuthUseCase
	rg *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (a *AuthController) loginHandler(ctx *gin.Context) {
	var payload dto.AuthRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	rsv, err := a.authUc.Login(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, rsv, "Ok")
}

func (a *AuthController) logoutHandler(ctx *gin.Context) {
	// Dapatkan token dari header Authorization
	token := ctx.GetHeader("Authorization")

	// Panggil metode Logout dari AuthUseCase
	if err := a.authUc.Logout(token); err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Kirim respons bahwa logout berhasil
	common.SendLogoutResponse(ctx, "Logout successful")
}

func (a *AuthController) Route() {
	a.rg.POST(config.AuthLogin, a.loginHandler)
	a.rg.POST(config.AuthLogout, a.authMiddleware.RequireToken(), a.logoutHandler)

}

func NewAuthController(authUc usecase.AuthUseCase, authMiddleware middleware.AuthMiddleware,rg *gin.RouterGroup) *AuthController {
	return &AuthController{authUc: authUc, rg: rg, authMiddleware: authMiddleware,}
}
