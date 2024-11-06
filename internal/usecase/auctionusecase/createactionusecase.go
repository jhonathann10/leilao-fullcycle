package auctionusecase

import (
	"context"
	"time"

	"github.com/jhonathann10/leilao-fullcycle/internal/entity/auctionentity"
	"github.com/jhonathann10/leilao-fullcycle/internal/entity/bidentity"
	internalerrors "github.com/jhonathann10/leilao-fullcycle/internal/errors"
	"github.com/jhonathann10/leilao-fullcycle/internal/usecase/bidusecase"
)

type AuctionInputDTO struct {
	ProductName string           `json:"product_name" binding:"required,min=1"`
	Category    string           `json:"category" binding:"required,min=2"`
	Description string           `json:"description" binding:"required,min=10,max=200"`
	Condition   ProductCondition `json:"condition"`
}

type AuctionOutputDTO struct {
	ID          string           `json:"id"`
	ProductName string           `json:"product_name"`
	Category    string           `json:"category"`
	Description string           `json:"description"`
	Condition   ProductCondition `json:"condition"`
	Status      AuctionStatus    `json:"status"`
	TimeStamp   time.Time        `json:"time_stamp" time_format:"2006-01-02 15:04:05"`
}

type WinningInfoOutputDTO struct {
	AuctionID AuctionOutputDTO         `json:"auction"`
	Bid       *bidusecase.BidOutputDTO `json:"bid,omitempty"`
}

type ProductCondition int64
type AuctionStatus int64

type AuctionUseCase struct {
	auctionRepositoryInterface auctionentity.AuctionRepositoryInterface
	bidRespositoryInterface    bidentity.BidEntityRepository
}

type AuctionUseCaseInterface interface {
	CreateAuction(ctx context.Context, auctionInput AuctionInputDTO) *internalerrors.InternalError
	FindAuctionByID(ctx context.Context, id string) (*AuctionOutputDTO, *internalerrors.InternalError)
	FindAuctions(ctx context.Context, status AuctionStatus, category, productName string) ([]AuctionOutputDTO, *internalerrors.InternalError)
	FindWinningBidByAuctionID(ctx context.Context, auctionID string) (*WinningInfoOutputDTO, *internalerrors.InternalError)
}

func NewAuctionUseCase(auctionRepositoryInterface auctionentity.AuctionRepositoryInterface, bidEntityRepository bidentity.BidEntityRepository) AuctionUseCaseInterface {
	return &AuctionUseCase{
		auctionRepositoryInterface: auctionRepositoryInterface,
		bidRespositoryInterface:    bidEntityRepository,
	}
}

func (au *AuctionUseCase) CreateAuction(ctx context.Context, auctionInput AuctionInputDTO) *internalerrors.InternalError {
	auction, err := auctionentity.CreateAuction(auctionInput.ProductName, auctionInput.Category, auctionInput.Description, auctionentity.ProductCondition(auctionInput.Condition))
	if err != nil {
		return err
	}

	if err := au.auctionRepositoryInterface.CreateAuction(ctx, auction); err != nil {
		return err
	}

	return nil
}
