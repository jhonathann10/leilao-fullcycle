package bidentity

import (
	"context"
	"time"

	"github.com/google/uuid"
	internalerrors "github.com/jhonathann10/leilao-fullcycle/internal/errors"
)

type Bid struct {
	ID        string
	UserID    string
	AuctionID string
	Amount    float64
	Timestamp time.Time
}

type BidEntityRepository interface {
	CreateBid(ctx context.Context, bidEntities []Bid) *internalerrors.InternalError
	FindBidByAuctionID(ctx context.Context, auctionID string) ([]Bid, *internalerrors.InternalError)
	FindWinningBidByAuctionID(ctx context.Context, auctionID string) (*Bid, *internalerrors.InternalError)
}

func CreateBid(userID, auctionID string, amount float64) (*Bid, *internalerrors.InternalError) {
	bid := &Bid{
		ID:        uuid.New().String(),
		UserID:    userID,
		AuctionID: auctionID,
		Amount:    amount,
		Timestamp: time.Now(),
	}

	if err := bid.Validate(); err != nil {
		return nil, err
	}

	return bid, nil
}

func (b *Bid) Validate() *internalerrors.InternalError {
	if err := uuid.Validate(b.UserID); err != nil {
		return internalerrors.NewBadRequestError("Invalid userID")
	}

	if err := uuid.Validate(b.AuctionID); err != nil {
		return internalerrors.NewBadRequestError("Invalid auctionID")
	}

	if b.Amount <= 0 {
		return internalerrors.NewBadRequestError("Invalid Amount")
	}

	return nil
}
