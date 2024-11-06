package auction

import (
	"context"

	"github.com/jhonathann10/leilao-fullcycle/configuration/logger"
	"github.com/jhonathann10/leilao-fullcycle/internal/entity/auctionentity"
	internalerrors "github.com/jhonathann10/leilao-fullcycle/internal/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	ID          string                         `bson:"_id"`
	ProductName string                         `bson:"product_name"`
	Category    string                         `bson:"category"`
	Description string                         `bson:"description"`
	Condition   auctionentity.ProductCondition `bson:"condition"`
	Status      auctionentity.AuctionStatus    `bson:"status"`
	TimeStamp   int64                          `bson:"timestamp"`
}

type AuctionRepository struct {
	Collection *mongo.Collection
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection: database.Collection("auctions"),
	}
}

func (ar *AuctionRepository) CreateAuction(ctx context.Context, auction *auctionentity.Auction) *internalerrors.InternalError {
	auctionEntityMongo := AuctionEntityMongo{
		ID:          auction.ID,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   auction.Condition,
		Status:      auction.Status,
		TimeStamp:   auction.TimeStamp.Unix(),
	}

	_, err := ar.Collection.InsertOne(ctx, auctionEntityMongo)
	if err != nil {
		logger.Error("Error inserting auction in database", err)
		return internalerrors.NewInternalServerError("Error inserting auction in database")
	}

	return nil
}
