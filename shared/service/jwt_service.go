package service

import (
	"fmt"
	"test-mnc/config"
	"test-mnc/entity"
	"test-mnc/entity/dto"
	"test-mnc/shared/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	CreateToken(user entity.Customers) (dto.AuthResponseDto, error)
	ParseToken(tokenHeader string) (jwt.MapClaims, error)
	InvalidateToken(token string) error
}

type jwtService struct {
	cfg config.TokenConfig
	invalidTokens map[string]bool
}

func (j *jwtService) CreateToken(user entity.Customers) (dto.AuthResponseDto, error) {
	claims := model.MyCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfg.IssuerName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.cfg.JwtExpiresTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserId: user.Id,
	}

	token := jwt.NewWithClaims(j.cfg.JwtSigningMethod, claims)
	ss, err := token.SignedString(j.cfg.JwtSignatureKy)
	if err != nil {
		return dto.AuthResponseDto{}, fmt.Errorf("oops, failed to create token")
	}
	return dto.AuthResponseDto{Token: ss}, nil
}

func (j *jwtService) ParseToken(tokenHeader string) (jwt.MapClaims, error) {
	if j.invalidTokens[tokenHeader] {
		return nil, fmt.Errorf("oops, token has been invalidated")
	}
	
	token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
		return j.cfg.JwtSignatureKy, nil
	})

	if err != nil {
		return nil, fmt.Errorf("oops, failed to verify token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("oops, failed to claim token")
	}
	return claims, nil
}

func (j *jwtService) InvalidateToken(token string) error {
	// Tandai token sebagai tidak valid (logout)
	j.invalidTokens[token] = true
	return nil
}
func NewJwtService(cfg config.TokenConfig) JwtService {
	return &jwtService{cfg: cfg, invalidTokens: make(map[string]bool)}
}