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

// WalletController はFirestoreコネクト用コントローラー
type WalletController struct {
	Interactor *interactor.WalletInteractor
	Presenter  *presenters.WalletPresenter
	Error      *presenters.ErrorPresenter
	Validator  *validator.Validate
}

// NewWalletController はfirestoreコネクト用Newコントローラー
func NewWalletController(interactor *interactor.WalletInteractor, logging logging.Logging, validate *validator.Validate) *WalletController {
	return &WalletController{
		Interactor: interactor,
		Presenter:  presenters.NewWalletPresenter(logging),
		Error:      presenters.NewErrorPresenter(logging),
		Validator:  validate,
	}
}

// List はDBにウォレット情報を配列で出力するハンドラー
// @Tags ウォレット情報
// Wallet godoc
// @Summary ウォレットの情報をデータベースから抽出する
// @Description 仮想通貨のウォレットの情報を出力する
// @Accept  json
// @Produce  json
// @Success 200 {object} []ports.WalletOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /wallets [get]
func (controller *WalletController) List(c echo.Context) error {
	ctx := c.Request().Context()

	outputs, err := controller.Interactor.List(ctx)
	if err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, outputs)
}

// Create はDBにウォレット情報を登録するハンドラー
// @Tags ウォレット情報
// Wallet godoc
// @Summary ウォレットアドレスをデータベースに格納する
// @Description ウォレットアドレスが登録されているかウォレットテーブルを確認し、なければウォレットアドレスをデータベースに格納する
// @Accept  json
// @Produce  json
// @Param wallet body ports.WalletInput true "ウォレットアドレス"
// @Success 200 {object} ports.WalletOutput
// @Success 201 {object} ports.CreatedObject
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /wallets [post]
func (controller *WalletController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var input ports.WalletInput

	if err := c.Bind(&input); err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	if err := controller.Validator.Struct(&input); err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	// 存在チェック
	output, err := controller.Interactor.Get(ctx, &input)
	if err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	if output.Address != "" {
		return controller.Presenter.Exist(c, output)
	}

	output, err = controller.Interactor.Create(ctx, &input)
	if err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	return controller.Presenter.Create(c, output)
}
