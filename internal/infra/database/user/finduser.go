package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/jhonathann10/leilao-fullcycle/configuration/logger"
	"github.com/jhonathann10/leilao-fullcycle/internal/entity/userentity"
	internalerrors "github.com/jhonathann10/leilao-fullcycle/internal/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserEntityMongo struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: database.Collection("users"),
	}
}

func (ur *UserRepository) FindUserByID(ctx context.Context, userID string) (*userentity.User, *internalerrors.InternalError) {
	filter := bson.M{"_id": userID}
	var userEntityMongo UserEntityMongo
	err := ur.Collection.FindOne(ctx, filter).Decode(&userEntityMongo)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errMsg := fmt.Sprintf("user not found with this id = %s", userID)
			logger.Error(errMsg, err)
			return nil, internalerrors.NewNotFoundError(errMsg)
		}
		errMsg := "Error trying to find user by userID"
		logger.Error(errMsg, err)
		return nil, internalerrors.NewNotFoundError(errMsg)
	}
	userEntity := &userentity.User{
		ID:   userEntityMongo.ID,
		Name: userEntityMongo.Name,
	}

	return userEntity, nil
}
