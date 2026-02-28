// Package server は、HTTPサーバーのセットアップとルーティングを定義します。
package server

import (
	"net/http"
	"os"

	"nft-music/adapters/controllers"
	"nft-music/adapters/gateways"
	"nft-music/contracts"
	"nft-music/usecases/interactor"
	"nft-music/usecases/logging"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
)

// Run はHTTPサーバーを起動し、ルートを設定します。
func Run(
	db *gorm.DB,
	etherClient *ethclient.Client,
	etherAuth *bind.TransactOpts,
	contracts *contracts.Contracts,
	logging logging.Logging,
	validate *validator.Validate,
) {
	e := echo.New()

	// ミドルウェア
	// e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper:          middleware.DefaultSkipper,
		Format:           `{"error":"${error}","method":"${method}","status":${status},"uri":"${uri}"}`,
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}))

	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://music.threenext.com", "http://music.threenext.com"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := e.Group("/api/v1")
	{
		walletGateway := gateways.NewWalletGateway(db)
		walletInteractor := interactor.NewWalletInteractor(walletGateway)
		walletController := controllers.NewWalletController(walletInteractor, logging, validate)
		v1.GET("/wallets", walletController.List)
		v1.POST("/wallets", walletController.Create)

		businessGateway := gateways.NewBusinessGateway(db)
		businessInteractor := interactor.NewBusinessInteractor(businessGateway)
		businessController := controllers.NewBusinessController(businessInteractor, logging, validate)
		v1.POST("/businesses", businessController.Create)
		v1.GET("/businesses/:id", businessController.Get)
		v1.GET("/businesses", businessController.List)
		v1.PUT("/businesses/:id", businessController.Update)
		v1.DELETE("/businesses/:id", businessController.Delete)

		genreGateway := gateways.NewGenreGateway(db)
		genreInteractor := interactor.NewGenreInteractor(genreGateway)
		genreController := controllers.NewGenreController(genreInteractor, logging, validate)
		v1.POST("/genres", genreController.Create)
		v1.GET("/genres/:id", genreController.Get)
		v1.GET("/genres", genreController.List)
		v1.PUT("/genres/:id", genreController.Update)
		v1.DELETE("/genres/:id", genreController.Delete)

		ipfsGateway := gateways.NewIpfsGateway(db)
		userGateway := gateways.NewUserGateway(db)
		ipfsInteractor := interactor.NewIpfsInteractor(ipfsGateway, userGateway)
		ipfsController := controllers.NewIpfsController(ipfsInteractor, logging, validate)
		v1.POST("/ipfs", ipfsController.Upload)
		v1.POST("/ipfs/meta", ipfsController.MetaUpload)

		collectionGateway := gateways.NewCollectionGateway(db)
		collectionInteractor := interactor.NewCollectionInteractor(collectionGateway)
		collectionController := controllers.NewCollectionController(collectionInteractor, logging, validate)
		v1.POST("/collections", collectionController.Create)
		v1.GET("/collections/:id", collectionController.Get)
		v1.GET("/collections", collectionController.List)
		v1.PUT("/collections/:id", collectionController.Update)
		v1.DELETE("/collections/:id", collectionController.Delete)

		transactionGateway := gateways.NewTransactionGateway(db)
		nftInteractor := interactor.NewNftInteractor(userGateway, transactionGateway, ipfsGateway, etherClient, etherAuth, contracts, logging, validate)
		nftController := controllers.NewNftController(nftInteractor, ipfsInteractor)
		v1.GET("/nfts/search", nftController.Search)
		v1.GET("/nfts", nftController.List)
		v1.GET("/nfts/:wallet", nftController.ListByWallet)
		v1.GET("/nfts/detail/:transaction_id", nftController.GetByTransactionid)
		v1.POST("/nfts", nftController.Mint)

		userInteractor := interactor.NewUserInteractor(userGateway, logging)
		userController := controllers.NewUserController(userInteractor)
		v1.GET("/users", userController.List)
		v1.GET("/users/:id", userController.Get)
		v1.GET("/users/wallet/:wallet", userController.GetByWallet) // ウォレットアドレスで取得するための明確なパス
		v1.POST("/users", userController.Create)
		v1.PUT("/users/:id", userController.Update)
		v1.DELETE("/users/:id", userController.Delete)

		evmInteractor := interactor.NewEvmInteractor(logging)
		blockChainController := controllers.NewBlockChainController(evmInteractor, etherAuth, contracts, logging)
		v1.POST("/evm", blockChainController.Signer)
	}
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "1323"
	}

	// サーバー起動
	e.Logger.Fatal(e.Start("0.0.0.0:" + port))
}
