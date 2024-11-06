package usercontroller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jhonathann10/leilao-fullcycle/configuration/rest_err"
	"github.com/jhonathann10/leilao-fullcycle/internal/usecase/userusecase"
)

type UserController struct {
	userUseCase userusecase.UserUseCaseInterface
}

func NewUserController(userUseCase userusecase.UserUseCaseInterface) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func (u *UserController) FindUserById(c *gin.Context) {
	userId := c.Param("userId")

	if err := uuid.Validate(userId); err != nil {
		errRest := rest_err.NewBadRequestError("invalid fields: ", rest_err.Causes{
			Field:   "userId",
			Message: "userId is not a valid UUID",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	userData, err := u.userUseCase.FindUserByID(context.Background(), userId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, userData)
}
