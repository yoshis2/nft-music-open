// Package controllersは、HTTPリクエストのハンドリングとレスポンス制御を実装します。
package controllers

import (
	"net/http"

	"nft-music/contracts"
	"nft-music/usecases/interactor"
	"nft-music/usecases/logging"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/labstack/echo/v4"
)

type EvmController struct {
	Interactor *interactor.EvmInteractor
	Auth       *bind.TransactOpts
	Contracts  *contracts.Contracts
	Logging    logging.Logging
}

func NewBlockChainController(interactor *interactor.EvmInteractor, auth *bind.TransactOpts, contracts *contracts.Contracts, logging logging.Logging) *EvmController {
	return &EvmController{
		Interactor: interactor,
		Auth:       auth,
		Contracts:  contracts,
		Logging:    logging,
	}
}

// Signer はEVMにログインする
// @Tags EVM
// Evm godoc
// @Summary EVMにログインを取得
// @Description Ethereum Virtual Machineのログイン情報を取得する
// @Accept  json
// @Produce  json
// @Param wallet body ports.WalletInput true "ウォレット"
// @Success 200 {object} ports.SignerOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /evm [post]
func (controller *EvmController) Signer(c echo.Context) error {
	ctx := c.Request().Context()

	controller.Logging.Debug("通過しました")

	// controller.Evm.
	output, err := controller.Interactor.Signer(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, output)
}
