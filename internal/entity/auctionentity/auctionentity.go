package auctionentity

import (
	"context"
	"time"

	"github.com/google/uuid"
	internalerrors "github.com/jhonathann10/leilao-fullcycle/internal/errors"
)

func CreateAuction(productName, category, description string, condition ProductCondition) (*Auction, *internalerrors.InternalError) {
	auction := &Auction{
		ID:          uuid.New().String(),
		ProductName: productName,
		Category:    category,
		Description: description,
		Condition:   condition,
		Status:      Active,
		TimeStamp:   time.Now(),
	}

	if err := auction.Validate(); err != nil {
		return nil, err
	}

	return auction, nil
}

func (au *Auction) Validate() *internalerrors.InternalError {
	if len(au.ProductName) <= 1 || len(au.Category) <= 2 || len(au.Description) <= 10 && (au.Condition != New || au.Condition != Used || au.Condition != Refurbished) {
		return internalerrors.NewBadRequestError("invalid auction")
	}

	return nil
}

type Auction struct {
	ID          string
	ProductName string
	Category    string
	Description string
	Condition   ProductCondition
	Status      AuctionStatus
	TimeStamp   time.Time
}

type ProductCondition int
type AuctionStatus int

const (
	Active AuctionStatus = iota
	Completed
)

const (
	New ProductCondition = iota
	Used
	Refurbished
)

type AuctionRepositoryInterface interface {
	CreateAuction(ctx context.Context, auction *Auction) *internalerrors.InternalError
	FindAuctionByID(ctx context.Context, id string) (*Auction, *internalerrors.InternalError)
	FindAuctions(ctx context.Context, status AuctionStatus, category, productName string) ([]Auction, *internalerrors.InternalError)
}
