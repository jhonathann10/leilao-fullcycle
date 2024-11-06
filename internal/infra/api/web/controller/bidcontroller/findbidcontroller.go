package bidcontroller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jhonathann10/leilao-fullcycle/configuration/rest_err"
)

func (b *BidController) FindBidByAuctionID(c *gin.Context) {
	auctionId := c.Param("auctionId")

	if err := uuid.Validate(auctionId); err != nil {
		errRest := rest_err.NewBadRequestError("invalid fields: ", rest_err.Causes{
			Field:   "auctionId",
			Message: "auctionId is not a valid UUID",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	bidOutputList, err := b.bidUseCase.FindBidByAuctionID(context.Background(), auctionId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, bidOutputList)
}
