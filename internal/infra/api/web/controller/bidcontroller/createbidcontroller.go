package bidcontroller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhonathann10/leilao-fullcycle/configuration/rest_err"
	"github.com/jhonathann10/leilao-fullcycle/internal/infra/api/web/validation"
	"github.com/jhonathann10/leilao-fullcycle/internal/usecase/bidusecase"
)

type BidController struct {
	bidUseCase bidusecase.BidUseCaseInterface
}

func NewBidController(bidUseCase bidusecase.BidUseCaseInterface) *BidController {
	return &BidController{
		bidUseCase: bidUseCase,
	}
}

func (b *BidController) CreateBid(c *gin.Context) {
	var bidInputDTO bidusecase.BidInputDTO

	if err := c.ShouldBindJSON(&bidInputDTO); err != nil {
		restErr := validation.ValidateErr(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	err := b.bidUseCase.CreateBid(context.Background(), bidInputDTO)
	if err != nil {
		restErr := rest_err.ConvertError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(http.StatusCreated)
}
