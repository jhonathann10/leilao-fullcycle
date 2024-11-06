package auctioncontroller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhonathann10/leilao-fullcycle/configuration/rest_err"
	"github.com/jhonathann10/leilao-fullcycle/internal/infra/api/web/validation"
	"github.com/jhonathann10/leilao-fullcycle/internal/usecase/auctionusecase"
)

type AuctionController struct {
	auctionUseCase auctionusecase.AuctionUseCaseInterface
}

func NewAuctionController(auctionUseCase auctionusecase.AuctionUseCaseInterface) *AuctionController {
	return &AuctionController{
		auctionUseCase: auctionUseCase,
	}
}

func (a *AuctionController) CreateAuction(c *gin.Context) {
	var auctionInputDTO auctionusecase.AuctionInputDTO

	if err := c.ShouldBindJSON(&auctionInputDTO); err != nil {
		restErr := validation.ValidateErr(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	err := a.auctionUseCase.CreateAuction(context.Background(), auctionInputDTO)
	if err != nil {
		restErr := rest_err.ConvertError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(http.StatusCreated)
}
