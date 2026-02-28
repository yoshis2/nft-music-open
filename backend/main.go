// Package main は、アプリケーションの実行エントリポイントです。
package main

import (
	"os"

	"nft-music/docs"
	"nft-music/infrastructure/ethereum"
	"nft-music/infrastructure/logging"
	"nft-music/infrastructure/mysql"
	"nft-music/infrastructure/server"

	"github.com/go-playground/validator/v10"
)

// @contact.name NFT Music
// @contact.url https://nft.threenext.com
// @contact.email seki@threenext.com
// @termsOfService http://swagger.io/terms/

// @securityDefinitions.basic BasicAuth
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	swaggerSet()

	newMysql := mysql.NewMysql()
	client := newMysql.Open()

	logging := logging.NewZapLogging()

	validate := validator.New()

	newEthereumVirtualMachine := ethereum.NewEthereumVirtualMachine()
	etherClient, transactOpts, contracts, err := newEthereumVirtualMachine.Auth()
	if err != nil {
		panic(err)
	}

	server.Run(client, etherClient, transactOpts, contracts, logging, validate)
}

func swaggerSet() {
	docs.SwaggerInfo.Title = "NFT MusicのAPIドキュメント Swagger"
	docs.SwaggerInfo.Description = "NFT Musicの音楽に特化したNFT販売のAPI"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{os.Getenv("SWAGGER_SCHEMA")}
}
