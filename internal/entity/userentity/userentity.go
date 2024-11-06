package userentity

import (
	"context"

	internalerrors "github.com/jhonathann10/leilao-fullcycle/internal/errors"
)

type User struct {
	ID   string
	Name string
}

type UserRepositoryInterface interface {
	FindUserByID(ctx context.Context, userID string) (*User, *internalerrors.InternalError)
}
