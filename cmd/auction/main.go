package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jhonathann10/leilao-fullcycle/configuration/database/mongodb"
	"github.com/jhonathann10/leilao-fullcycle/internal/infra/api/web/controller/auctioncontroller"
	"github.com/jhonathann10/leilao-fullcycle/internal/infra/api/web/controller/bidcontroller"
	"github.com/jhonathann10/leilao-fullcycle/internal/infra/api/web/controller/usercontroller"
	"github.com/jhonathann10/leilao-fullcycle/internal/infra/database/auction"
	"github.com/jhonathann10/leilao-fullcycle/internal/infra/database/bid"
	"github.com/jhonathann10/leilao-fullcycle/internal/infra/database/user"
	"github.com/jhonathann10/leilao-fullcycle/internal/usecase/auctionusecase"
	"github.com/jhonathann10/leilao-fullcycle/internal/usecase/bidusecase"
	"github.com/jhonathann10/leilao-fullcycle/internal/usecase/userusecase"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	ctx := context.Background()
	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Error trying to load env variables")
		return
	}

	databaseConnection, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	router := gin.Default()

	userController, bidController, auctionController := initDependencies(databaseConnection)

	router.GET("/auctions", auctionController.FindAuctions)
	router.GET("/auctions/:auctionId", auctionController.FindAuctionById)
	router.POST("/auctions", auctionController.CreateAuction)
	router.GET("/auction/winner/:auctionId", auctionController.FindWinningBidByAuctionId)
	router.GET("/bid/:auctionId", bidController.FindBidByAuctionID)
	router.POST("/bid", bidController.CreateBid)
	router.GET("/user/:userId", userController.FindUserById)

	router.Run(":8080")

}

func initDependencies(database *mongo.Database) (userController *usercontroller.UserController, bidController *bidcontroller.BidController, auctionController *auctioncontroller.AuctionController) {
	auctionRepository := auction.NewAuctionRepository(database)
	bidRepository := bid.NewBidRepository(database, auctionRepository)
	userRepository := user.NewUserRepository(database)

	userController = usercontroller.NewUserController(userusecase.NewUserUseCase(userRepository))
	auctionController = auctioncontroller.NewAuctionController(auctionusecase.NewAuctionUseCase(auctionRepository, bidRepository))
	bidController = bidcontroller.NewBidController(bidusecase.NewBidUseCase(bidRepository))

	return
}
