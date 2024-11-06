package auctioncontroller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jhonathann10/leilao-fullcycle/configuration/rest_err"
	"github.com/jhonathann10/leilao-fullcycle/internal/usecase/auctionusecase"
)

func (a *AuctionController) FindAuctionById(c *gin.Context) {
	auctionId := c.Param("auctionId")

	if err := uuid.Validate(auctionId); err != nil {
		errRest := rest_err.NewBadRequestError("invalid fields: ", rest_err.Causes{
			Field:   "auctionId",
			Message: "auctionId is not a valid UUID",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	auctionData, err := a.auctionUseCase.FindAuctionByID(context.Background(), auctionId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auctionData)
}

func (a *AuctionController) FindAuctions(c *gin.Context) {
	status := c.Param("status")
	category := c.Param("category")
	productName := c.Param("productName")

	statusNumber, errConv := strconv.Atoi(status)
	if errConv != nil {
		errRest := rest_err.NewBadRequestError("Error trying to validate auction status param")
		c.JSON(errRest.Code, errRest)
		return
	}

	auctions, err := a.auctionUseCase.FindAuctions(context.Background(), auctionusecase.AuctionStatus(statusNumber), category, productName)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auctions)
}

func (a *AuctionController) FindWinningBidByAuctionId(c *gin.Context) {
	auctionId := c.Param("auctionId")

	if err := uuid.Validate(auctionId); err != nil {
		errRest := rest_err.NewBadRequestError("invalid fields: ", rest_err.Causes{
			Field:   "auctionId",
			Message: "auctionId is not a valid UUID",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	auctionData, err := a.auctionUseCase.FindWinningBidByAuctionID(context.Background(), auctionId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auctionData)
}
