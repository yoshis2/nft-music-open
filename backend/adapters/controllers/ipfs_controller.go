// Package controllersは、HTTPリクエストのハンドリングとレスポンス制御を実装します。
package controllers

import (
	"net/http"

	"nft-music/adapters/presenters"
	"nft-music/usecases/interactor"
	"nft-music/usecases/logging"
	"nft-music/usecases/ports"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// IpfsController はIPFSのアップロード用のコントローラー
type IpfsController struct {
	Interactor *interactor.IpfsInteractor
	Error      *presenters.ErrorPresenter
	Validator  *validator.Validate
}

// NewIpfsController のデータを格納する
func NewIpfsController(interactor *interactor.IpfsInteractor, logging logging.Logging, validate *validator.Validate) *IpfsController {
	return &IpfsController{
		Interactor: interactor,
		Error:      presenters.NewErrorPresenter(logging),
		Validator:  validate,
	}
}

// Upload はイメージをIPFS登録するハンドラー
// @Tags IPFS
// Ipfs godoc
// @Summary IPFSノードにイメージデータを登録
// @Description 分散型ストレージIPFSに画像を登録する
// @Accept multipart/form-data
// @Produce  json
// @Param	file	formData file true	"this is a test file"
// @Param wallet formData string true "ウォレットアドレス"
// @Success 200 {object} ports.IpfsOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /ipfs [post]
func (controller *IpfsController) Upload(c echo.Context) error {
	ctx := c.Request().Context()

	wallet := c.FormValue("wallet")
	header, err := c.FormFile("file")
	if err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	form := ports.IpfsInput{
		Wallet: wallet,
		File:   header.Filename,
	}

	output, err := controller.Interactor.Upload(ctx, header, form)
	if err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, output)
}

// MetaUpload はMetaJSONをIPFS登録するハンドラー
// @Tags IPFS
// Ipfs godoc
// @Summary IPFSノードにJSONデータを登録
// @Description 分散型ストレージIPFSにJSONデータを登録する
// @Accept json
// @Produce  json
// @Param json body ports.IpfsMetaInput true "MetaJSON"
// @Success 200 {object} ports.IpfsOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /ipfs/meta [post]
func (controller *IpfsController) MetaUpload(c echo.Context) error {
	ctx := c.Request().Context()
	var input ports.IpfsMetaInput

	if err := c.Bind(&input); err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	if input.ImageCid == "" {
		input.FileType = "video"
	} else {
		input.FileType = "audio"
	}

	if err := controller.Validator.Struct(&input); err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	output, err := controller.Interactor.MetaJSON(ctx, input)
	if err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, output)
}
