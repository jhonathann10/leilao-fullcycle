package auction

import (
	"context"
	"fmt"
	"time"

	"github.com/jhonathann10/leilao-fullcycle/configuration/logger"
	"github.com/jhonathann10/leilao-fullcycle/internal/entity/auctionentity"
	internalerrors "github.com/jhonathann10/leilao-fullcycle/internal/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ar *AuctionRepository) FindAuctionByID(ctx context.Context, id string) (*auctionentity.Auction, *internalerrors.InternalError) {
	var auctionEntityMongo AuctionEntityMongo
	filter := bson.M{"_id": id}
	if err := ar.Collection.FindOne(ctx, filter).Decode(&auctionEntityMongo); err != nil {
		errMsg := fmt.Sprintf("Error finding auction by id %s in database", id)
		logger.Error(errMsg, err)
		return nil, internalerrors.NewInternalServerError(errMsg)
	}

	return &auctionentity.Auction{
		ID:          auctionEntityMongo.ID,
		ProductName: auctionEntityMongo.ProductName,
		Category:    auctionEntityMongo.Category,
		Description: auctionEntityMongo.Description,
		Condition:   auctionEntityMongo.Condition,
		Status:      auctionEntityMongo.Status,
		TimeStamp:   time.Unix(auctionEntityMongo.TimeStamp, 0),
	}, nil
}

func (ar *AuctionRepository) FindAuctions(ctx context.Context, status auctionentity.AuctionStatus, category, productName string) ([]auctionentity.Auction, *internalerrors.InternalError) {
	filter := bson.M{}
	if status != 0 {
		filter["status"] = status
	}

	if category != "" {
		filter["category"] = category
	}

	if productName != "" {
		filter["productName"] = primitive.Regex{
			Pattern: productName,
			Options: "i",
		}
	}

	cursor, err := ar.Collection.Find(ctx, filter)
	if err != nil {
		errMsg := "Error finding auctions in database"
		logger.Error(errMsg, err)
		return nil, internalerrors.NewInternalServerError(errMsg)
	}
	defer cursor.Close(ctx)

	var auctionEntityMongo []AuctionEntityMongo
	if err := cursor.All(ctx, &auctionEntityMongo); err != nil {
		errMsg := "Error decoding auctions from database"
		logger.Error(errMsg, err)
		return nil, internalerrors.NewInternalServerError(errMsg)
	}

	var auctionEntity []auctionentity.Auction
	for _, auctionMongo := range auctionEntityMongo {
		auctionEntity = append(auctionEntity, auctionentity.Auction{
			ID:          auctionMongo.ID,
			ProductName: auctionMongo.ProductName,
			Category:    auctionMongo.Category,
			Description: auctionMongo.Description,
			Condition:   auctionMongo.Condition,
			Status:      auctionMongo.Status,
			TimeStamp:   time.Unix(auctionMongo.TimeStamp, 0),
		})
	}

	return auctionEntity, nil
}
