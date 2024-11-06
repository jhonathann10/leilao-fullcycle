package userusecase

import (
	"context"

	"github.com/jhonathann10/leilao-fullcycle/internal/entity/userentity"
	internalerrors "github.com/jhonathann10/leilao-fullcycle/internal/errors"
)

type UserUseCase struct {
	userRepository userentity.UserRepositoryInterface
}

type UserOutputDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewUserUseCase(userRepository userentity.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		userRepository: userRepository,
	}
}

type UserUseCaseInterface interface {
	FindUserByID(ctx context.Context, id string) (*UserOutputDTO, *internalerrors.InternalError)
}

func (u *UserUseCase) FindUserByID(ctx context.Context, id string) (*UserOutputDTO, *internalerrors.InternalError) {
	userEntity, err := u.userRepository.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &UserOutputDTO{
		ID:   userEntity.ID,
		Name: userEntity.Name,
	}, nil
}
