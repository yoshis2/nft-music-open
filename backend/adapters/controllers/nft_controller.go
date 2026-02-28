// Package controllersは、HTTPリクエストのハンドリングとレスポンス制御を実装します。
package controllers

import (
	"log"
	"net/http"
	"strconv"

	"nft-music/usecases/interactor"
	"nft-music/usecases/ports"

	"github.com/labstack/echo/v4"
)

type NftController struct {
	IpfsInteractor *interactor.IpfsInteractor
	NftInteractor  *interactor.NftInteractor
}

func NewNftController(nftInteractor *interactor.NftInteractor, ipfsInteractor *interactor.IpfsInteractor) *NftController {
	return &NftController{
		NftInteractor:  nftInteractor,
		IpfsInteractor: ipfsInteractor,
	}
}

// List はブロックチェーンにNFTを複数出力するハンドラー
// @Tags NFT情報
// Nft godoc
// @Summary ブロックチェーンにNFTを複数出力する
// @Description はブロックチェーンにNFTを複数出力する
// @Accept  json
// @Produce  json
// @Param limit query int false "取得件数"
// @Success 200 {object} []ports.TransactionOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /nfts [get]
func (controller *NftController) List(c echo.Context) error {
	ctx := c.Request().Context()

	limitQuery := c.QueryParam("limit")
	limit, err := strconv.Atoi(limitQuery)
	if err != nil || limit <= 0 {
		limit = 0 // パラメータが無い、または無効な場合は0（無制限）とする
	}

	outputs, err := controller.NftInteractor.List(ctx, limit)
	if err != nil {
		return controller.NftInteractor.Error.ErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, outputs)
}

// ListByWallet はウォレットアドレスでNFTを複数出力するハンドラー
// @Tags NFT情報
// @Summary ウォレットアドレスでNFTを複数出力する
// @Description ウォレットアドレスに紐づくNFTを複数出力する
// @Accept  json
// @Produce  json
// @Param wallet path string true "ウォレットアドレス"
// @Success 200 {object} []ports.TransactionOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /nfts/{wallet} [get]
func (controller *NftController) ListByWallet(c echo.Context) error {
	ctx := c.Request().Context()

	wallet := c.Param("wallet")

	outputs, err := controller.NftInteractor.ListByWallet(ctx, wallet)
	if err != nil {
		return controller.NftInteractor.Error.ErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, outputs)
}

// Search はキーワードでNFTを複数出力するハンドラー
// @Tags NFT情報
// @Summary キーワードでNFTを複数出力する
// @Description キーワードに一致するNFTを複数出力する
// @Accept  json
// @Produce  json
// @Param q query string false "検索キーワード"
// @Param genre query string false "ジャンルID"
// @Param min_price query int false "最小価格"
// @Param max_price query int false "最大価格"
// @Param sort query string false "ソート順"
// @Success 200 {object} []ports.TransactionOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /nfts/search [get]
func (controller *NftController) Search(c echo.Context) error {
	ctx := c.Request().Context()
	query := c.QueryParam("q")
	genre := c.QueryParam("genre")
	minPrice, err := strconv.Atoi(c.QueryParam("min_price"))
	if err != nil {
		minPrice = 0 // or handle error appropriately
	}
	maxPrice, err := strconv.Atoi(c.QueryParam("max_price"))
	if err != nil {
		maxPrice = 0 // or handle error appropriately
	}
	sort := c.QueryParam("sort")

	outputs, err := controller.NftInteractor.Search(ctx, query, genre, minPrice, maxPrice, sort)
	if err != nil {
		return controller.NftInteractor.Error.ErrorResponse(c, err)
	}
	return c.JSON(http.StatusOK, outputs)
}

// GetByTransactionid はトランザクションIDでNFTを1件出力するハンドラー
// @Tags NFT情報
// @Summary トランザクションIDでNFTを1件出力する
// @Description トランザクションIDに紐づくNFTを1件出力する
// @Accept  json
// @Produce  json
// @Param transaction_id path string true "トランザクションID"
// @Success 200 {object} ports.TransactionOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /nfts/detail/{transaction_id} [get]
func (controller *NftController) GetByTransactionid(c echo.Context) error {
	ctx := c.Request().Context()

	transactionID := c.Param("transaction_id")

	output, err := controller.NftInteractor.GetByTransactionid(ctx, transactionID)
	if err != nil {
		return controller.NftInteractor.Error.ErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, output)
}

// Mint はブロックチェーンにNFTをで登録するハンドラー
// @Tags NFT情報
// Nft godoc
// @Summary NFTの情報をブロックチェーンに登録する
// @Description NFTの情報をブロックチェーンに登録する
// @Accept  json
// @Produce  json
// @Param wallet body ports.NftInput true "ジャンルマスター"
// @Success 200 {object} ports.TransactionOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /nfts [post]
func (controller *NftController) Mint(c echo.Context) error {
	ctx := c.Request().Context()

	var input ports.NftInput
	if err := c.Bind(&input); err != nil {
		return err
	}

	if err := controller.NftInteractor.Validator.Struct(&input); err != nil {
		return controller.NftInteractor.Error.ErrorResponse(c, err)
	}

	metaJSON := ports.IpfsMetaInput{
		Name:        input.Name,
		Description: input.Description,
		FileType:    input.FileType,
		ImageCid:    input.ImageCid,
		AudioCid:    input.AudioCid,
		VideoCid:    input.VideoCid,
		Insentive:   input.Insentive,
	}

	// meta json をIPFSに登録
	token, err := controller.IpfsInteractor.MetaJSON(ctx, metaJSON)
	if err != nil {
		log.Printf("エラーーの内容 : %v", err)
		return controller.NftInteractor.Error.ErrorResponse(c, err)
	}

	// NFTをブロックチェーンに登録
	output, err := controller.NftInteractor.Mint(ctx, &input, token.Cid)
	if err != nil {
		return controller.NftInteractor.Error.ErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, output)
}
