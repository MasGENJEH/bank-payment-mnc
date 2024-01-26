package usecase

import (
	"test-mnc/entity/dto"
	"test-mnc/repository"
	"test-mnc/shared/service"
)

type AuthUseCase interface {
	Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error)
	Logout(user dto.AuthResponseDto) error
}

type authUseCase struct {
	userUC CustomersUseCase
	jwtService service.JwtService
	Repository repository.CustomerRepository
}

func (a *authUseCase) Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	user, err := a.userUC.FindCustomerForLogin(payload.User, payload.Password)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}
	token, err := a.jwtService.CreateToken(user)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	return token, nil
}

func (a *authUseCase) Logout(user dto.AuthResponseDto) error{
	if err := a.userUC.DeleteTokenForLogout(user); err != nil {
		return err
	}
	return nil
}

func NewAuthUseCase(userUC CustomersUseCase, jwtService service.JwtService) AuthUseCase {
	return &authUseCase{userUC: userUC, jwtService: jwtService}
}