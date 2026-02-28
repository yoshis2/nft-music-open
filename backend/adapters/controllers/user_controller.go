// Package controllersは、HTTPリクエストのハンドリングとレスポンス制御を実装します。
package controllers

import (
	"net/http"

	"nft-music/usecases/interactor"
	"nft-music/usecases/ports"
	"nft-music/util"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	Interactor *interactor.UserInteractor
}

func NewUserController(interactor *interactor.UserInteractor) *UserController {
	return &UserController{
		Interactor: interactor,
	}
}

// Get はNFTミュージックのアカウント情報を取得
// @Tags アカウント
// User godoc
// @Summary NFTミュージックのアカウントを登録する
// @Description NFTミュージックのアカウントを登録する
// @Accept  json
// @Produce  json
// @Param id path string true "ユーザーID"
// @Success 200 {object} ports.UserOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /users/{id} [get]
func (controller *UserController) Get(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	idUUID, err := uuid.Parse(id) // idをUUIDに変換
	if err != nil {
		return controller.Interactor.Error.ErrorResponse(c, err)
	}

	output, err := controller.Interactor.Get(ctx, idUUID)
	if err != nil {
		return controller.Interactor.Error.ErrorResponse(c, err)
	}

	user := ports.UserOutput{
		ID:         output.ID,
		Name:       output.Name,
		Email:      output.Email,
		Wallet:     output.Wallet,
		Address:    *util.EmptyString(output.Address),
		BusinessID: output.BusinessID,
		Website:    *util.EmptyString(output.Website),
		FaceImage:  *util.EmptyString(output.FaceImage),
		Eyecatch:   *util.EmptyString(output.Eyecatch),
		Profile:    *util.EmptyString(output.Profile),
		Role:       output.Role,
		CreatedAt:  output.CreatedAt,
		UpdatedAt:  output.UpdatedAt,
	}

	return c.JSON(http.StatusOK, user)
}

// GetByWallet はウォレットアドレスからアカウント情報を取得
// @Tags アカウント
// User godoc
// @Summary ウォレットアドレスからアカウント情報を取得する
// @Description ウォレットアドレスをキーにNFTミュージックのアカウント情報を取得する
// @Accept  json
// @Produce  json
// @Param wallet path string true "ウォレットアドレス"
// @Success 200 {object} ports.UserOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /users/wallet/{wallet} [get]
func (controller *UserController) GetByWallet(c echo.Context) error {
	ctx := c.Request().Context()

	wallet := c.Param("wallet")

	output, err := controller.Interactor.GetByWallet(ctx, wallet)
	if err != nil {
		return controller.Interactor.Error.ErrorResponse(c, err)
	}

	user := ports.UserOutput{
		ID:         output.ID,
		Name:       output.Name,
		Email:      output.Email,
		Wallet:     output.Wallet,
		Address:    *util.EmptyString(output.Address),
		BusinessID: output.BusinessID,
		Website:    *util.EmptyString(output.Website),
		FaceImage:  *util.EmptyString(output.FaceImage),
		Eyecatch:   *util.EmptyString(output.Eyecatch),
		Profile:    *util.EmptyString(output.Profile),
		Role:       output.Role,
		CreatedAt:  output.CreatedAt,
		UpdatedAt:  output.UpdatedAt,
	}

	return c.JSON(http.StatusOK, user)
}

// List はNFTミュージックのアカウントリストを取得
// @Tags アカウント
// User godoc
// @Summary NFTミュージックのアカウントを登録する
// @Description NFTミュージックのアカウントを登録する
// @Accept  json
// @Produce  json
// @Success 200 {object} []ports.UserOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /users [get]
func (controller *UserController) List(c echo.Context) error {
	ctx := c.Request().Context()

	outputs, err := controller.Interactor.List(ctx)
	if err != nil {
		return controller.Interactor.Error.ErrorResponse(c, err)
	}

	users := []ports.UserOutput{}
	for _, output := range outputs {
		user := ports.UserOutput{
			ID:         output.ID,
			Name:       output.Name,
			Email:      output.Email,
			Wallet:     output.Wallet,
			Address:    *util.EmptyString(output.Address),
			BusinessID: output.BusinessID,
			Website:    *util.EmptyString(output.Website),
			FaceImage:  *util.EmptyString(output.FaceImage),
			Eyecatch:   *util.EmptyString(output.Eyecatch),
			Profile:    *util.EmptyString(output.Profile),
			Role:       output.Role,
			CreatedAt:  output.CreatedAt,
			UpdatedAt:  output.UpdatedAt,
		}
		users = append(users, user)
	}

	return c.JSON(http.StatusOK, users)
}

// Create はNFTミュージックのアカウント作成
// @Tags アカウント
// User godoc
// @Summary NFTミュージックのアカウントを登録する
// @Description NFTミュージックのアカウントを登録する
// @Accept  json
// @Produce  json
// @Param users body ports.UserInput true "ユーザー情報"
// @Success 200 {object} ports.UserOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /users [post]
func (controller *UserController) Create(c echo.Context) error {
	ctx := c.Request().Context()

	var input ports.UserInput
	if err := c.Bind(&input); err != nil {
		return controller.Interactor.Error.ErrorResponse(c, err)
	}

	if err := controller.Interactor.Create(ctx, input); err != nil {
		return controller.Interactor.Error.ErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, "OK")
}

// Update はNFTミュージックのアカウント変更
// @Tags アカウント
// User godoc
// @Summary NFTミュージックのアカウントの情報を変更する
// @Description NFTミュージックのアカウントの情報を変更する
// @Accept  json
// @Produce  json
// @Param id path string true "ユーザーID"
// @Param users body ports.UserInput true "ユーザー情報"
// @Success 200 {object} ports.UserOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /users/{id} [put]
func (controller *UserController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	idUUID, err := uuid.Parse(id) // idをUUIDに変換
	if err != nil {
		return controller.Interactor.Error.ErrorResponse(c, err)
	}

	var input ports.UserInput
	if err := c.Bind(&input); err != nil {
		return controller.Interactor.Error.ErrorResponse(c, err)
	}

	if err := controller.Interactor.Update(ctx, idUUID, input); err != nil {
		return controller.Interactor.Error.ErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, "OK")
}

// Delete はNFTミュージックのアカウント削除
// @Tags アカウント
// User godoc
// @Summary NFTミュージックのアカウントを削除する
// @Description NFTミュージックのアカウントを削除する
// @Accept  json
// @Produce  json
// @Param id path string true "ユーザーID"
// @Success 200 {object} ports.UserOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /users/{id} [delete]
func (controller *UserController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	idUUID, err := uuid.Parse(id) // idをUUIDに変換
	if err != nil {
		return controller.Interactor.Error.ErrorResponse(c, err)
	}

	if err := controller.Interactor.Delete(ctx, idUUID); err != nil {
		return controller.Interactor.Error.ErrorResponse(c, err)
	}
	return c.JSON(http.StatusOK, "OK")
}
