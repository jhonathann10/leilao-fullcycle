package bidusecase

import (
	"context"

	internalerrors "github.com/jhonathann10/leilao-fullcycle/internal/errors"
)

func (bu *BidUseCase) FindBidByAuctionID(ctx context.Context, auctionID string) ([]BidOutputDTO, *internalerrors.InternalError) {
	bidList, err := bu.bidRepository.FindBidByAuctionID(ctx, auctionID)
	if err != nil {
		return nil, err
	}

	var bidOutputList []BidOutputDTO
	for _, value := range bidList {
		bidOutputList = append(bidOutputList, BidOutputDTO{
			ID:        value.ID,
			UserID:    value.UserID,
			AuctionID: value.AuctionID,
			Amount:    value.Amount,
			Timestamp: value.Timestamp,
		})
	}

	return bidOutputList, nil
}
func (bu *BidUseCase) FindWinningBidByAuctionID(ctx context.Context, auctionID string) (*BidOutputDTO, *internalerrors.InternalError) {
	bidEntity, err := bu.bidRepository.FindWinningBidByAuctionID(ctx, auctionID)
	if err != nil {
		return nil, err
	}

	return &BidOutputDTO{
		ID:        bidEntity.ID,
		UserID:    bidEntity.UserID,
		AuctionID: bidEntity.AuctionID,
		Amount:    bidEntity.Amount,
		Timestamp: bidEntity.Timestamp,
	}, nil
}
