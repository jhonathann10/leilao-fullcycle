package bid

import (
	"context"
	"fmt"
	"time"

	"github.com/jhonathann10/leilao-fullcycle/configuration/logger"
	"github.com/jhonathann10/leilao-fullcycle/internal/entity/bidentity"
	internalerrors "github.com/jhonathann10/leilao-fullcycle/internal/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (br *BidRepository) FindBidByAuctionID(ctx context.Context, auctionID string) ([]bidentity.Bid, *internalerrors.InternalError) {
	filter := bson.M{
		"auction_id": auctionID,
	}

	cursor, err := br.Collection.Find(ctx, filter)
	if err != nil {
		errMsg := fmt.Sprintf("Error trying to find bid by auction id: %s", err)
		logger.Error(errMsg, err)
		return nil, internalerrors.NewInternalServerError(errMsg)
	}

	var bidEntitiesMongo []BidEntityMongo
	if err := cursor.All(ctx, &bidEntitiesMongo); err != nil {
		errMsg := fmt.Sprintf("Error trying to find bid by auction id: %s", err)
		logger.Error(errMsg, err)
		return nil, internalerrors.NewInternalServerError(errMsg)
	}

	var bidEntities []bidentity.Bid
	for _, bidEntityMongo := range bidEntitiesMongo {
		bidEntities = append(bidEntities, bidentity.Bid{
			ID:        bidEntityMongo.ID,
			UserID:    bidEntityMongo.UserID,
			AuctionID: bidEntityMongo.AuctionID,
			Amount:    bidEntityMongo.Amount,
			Timestamp: time.Unix(bidEntityMongo.Timestamp, 0),
		})
	}

	return bidEntities, nil
}

func (br *BidRepository) FindWinningBidByAuctionID(ctx context.Context, auctionID string) (*bidentity.Bid, *internalerrors.InternalError) {
	filter := bson.M{
		"auction_id": auctionID,
	}

	var bidEntityMongo BidEntityMongo
	opts := options.FindOne().SetSort(bson.D{{"amount", -1}})
	if err := br.Collection.FindOne(ctx, filter, opts).Decode(&bidEntityMongo); err != nil {
		errMsg := fmt.Sprintf("Error trying to find winning bid by auction id: %s", err)
		logger.Error(errMsg, err)
		return nil, internalerrors.NewInternalServerError(errMsg)
	}

	return &bidentity.Bid{
		ID:        bidEntityMongo.ID,
		UserID:    bidEntityMongo.UserID,
		AuctionID: bidEntityMongo.AuctionID,
		Amount:    bidEntityMongo.Amount,
		Timestamp: time.Unix(bidEntityMongo.Timestamp, 0),
	}, nil
}
