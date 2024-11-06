package auctionusecase

import (
	"context"

	"github.com/jhonathann10/leilao-fullcycle/internal/entity/auctionentity"
	internalerrors "github.com/jhonathann10/leilao-fullcycle/internal/errors"
	"github.com/jhonathann10/leilao-fullcycle/internal/usecase/bidusecase"
)

func (au *AuctionUseCase) FindAuctionByID(ctx context.Context, id string) (*AuctionOutputDTO, *internalerrors.InternalError) {
	auctionEntity, err := au.auctionRepositoryInterface.FindAuctionByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &AuctionOutputDTO{
		ID:          auctionEntity.ID,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   ProductCondition(auctionEntity.Condition),
		Status:      AuctionStatus(auctionEntity.Status),
		TimeStamp:   auctionEntity.TimeStamp,
	}, nil
}

func (au *AuctionUseCase) FindAuctions(ctx context.Context, status AuctionStatus, category, productName string) ([]AuctionOutputDTO, *internalerrors.InternalError) {
	auctionEntities, err := au.auctionRepositoryInterface.FindAuctions(ctx, auctionentity.AuctionStatus(status), category, productName)
	if err != nil {
		return nil, err
	}

	var auctionOutputs []AuctionOutputDTO

	for _, value := range auctionEntities {
		auctionOutputs = append(auctionOutputs, AuctionOutputDTO{
			ID:          value.ID,
			ProductName: value.ProductName,
			Category:    value.Category,
			Description: value.Description,
			Condition:   ProductCondition(value.Condition),
			Status:      AuctionStatus(value.Status),
			TimeStamp:   value.TimeStamp,
		})
	}

	return auctionOutputs, nil
}

func (au *AuctionUseCase) FindWinningBidByAuctionID(ctx context.Context, auctionID string) (*WinningInfoOutputDTO, *internalerrors.InternalError) {
	auctionEntity, err := au.auctionRepositoryInterface.FindAuctionByID(ctx, auctionID)
	if err != nil {
		return nil, err
	}

	auctionOutputDTO := AuctionOutputDTO{
		ID:          auctionEntity.ID,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   ProductCondition(auctionEntity.Condition),
		Status:      AuctionStatus(auctionEntity.Status),
		TimeStamp:   auctionEntity.TimeStamp,
	}

	bidWinning, err := au.bidRespositoryInterface.FindWinningBidByAuctionID(ctx, auctionID)
	if err != nil {
		return &WinningInfoOutputDTO{
			AuctionID: auctionOutputDTO,
			Bid:       nil,
		}, nil
	}

	bidOutputDTO := &bidusecase.BidOutputDTO{
		ID:        bidWinning.ID,
		UserID:    bidWinning.UserID,
		AuctionID: bidWinning.AuctionID,
		Amount:    bidWinning.Amount,
		Timestamp: bidWinning.Timestamp,
	}

	return &WinningInfoOutputDTO{
		AuctionID: auctionOutputDTO,
		Bid:       bidOutputDTO,
	}, nil
}
