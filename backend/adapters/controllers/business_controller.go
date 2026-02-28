// Package controllersは、HTTPリクエストのハンドリングとレスポンス制御を実装します。
package controllers

import (
	"net/http"

	"nft-music/adapters/presenters"
	"nft-music/usecases/interactor"
	"nft-music/usecases/logging"
	"nft-music/usecases/ports"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// BusinessController はビジネスマスターの構造体
type BusinessController struct {
	Interactor *interactor.BusinessInteractor
	Error      *presenters.ErrorPresenter
	Validator  *validator.Validate
}

// NewBusinessController はビジネスマスターのコンストラクタ
func NewBusinessController(interactor *interactor.BusinessInteractor, logging logging.Logging, validate *validator.Validate) *BusinessController {
	return &BusinessController{
		Interactor: interactor,
		Error:      presenters.NewErrorPresenter(logging),
		Validator:  validate,
	}
}

// Create は職種マスターの情報を1件作成する
// @Tags 職種マスター
// Businesses godoc
// @Summary 職種マスターの情報を1件作成する
// @Description 職種マスターの情報を1件作成する
// @Accept  json
// @Produce  json
// @Param wallet body ports.BusinessMasterInput true "職種マスター"
// @Success 200 {object} ports.BusinessMasterOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /businesses [post]
func (controller *BusinessController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var input ports.BusinessMasterInput

	if err := c.Bind(&input); err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	if err := controller.Validator.Struct(&input); err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	output, err := controller.Interactor.Create(ctx, &input)
	if err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, output)
}

// Get は職種マスターの情報を1件取得する
// @Tags 職種マスター
// Businesses godoc
// @Summary 職種マスターから特定の１レコードを抽出する
// @Description 職種マスターの情報を出力し
// @Accept  json
// @Produce  json
// @Param id path string true "職種マスターID"
// @Success 200 {object} ports.BusinessMasterOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /businesses/{id} [get]
func (controller *BusinessController) Get(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	idUUID, err := uuid.Parse(id) // idをUUIDに変換
	if err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	output, err := controller.Interactor.Get(ctx, idUUID)
	if err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, output)
}

// List は職種マスターの情報をリストで取得する
// @Tags 職種マスター
// Businesses godoc
// @Summary 職種マスターの情報をリストで取得する
// @Description 職種マスターの情報をリストで取得する
// @Accept  json
// @Produce  json
// @Success 200 {object} []ports.BusinessMasterOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /businesses [get]
func (controller *BusinessController) List(c echo.Context) error {
	ctx := c.Request().Context()

	outputs, err := controller.Interactor.List(ctx)
	if err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, outputs)
}

// Update は職種マスターの情報を1件修正する
// @Tags 職種マスター
// Businesses godoc
// @Summary 職種マスターの情報を1件修正する
// @Description 職種マスターの情報を1件修正する
// @Accept  json
// @Produce  json
// @Param id path string true "職種マスターID"
// @Param wallet body ports.BusinessMasterInput true "職種マスター"
// @Success 200 {object} ports.BusinessMasterOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /businesses/{id} [put]
func (controller *BusinessController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	idUUID, err := uuid.Parse(id) // idをUUIDに変換
	if err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	var input ports.BusinessMasterInput
	if err := c.Bind(&input); err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	output, err := controller.Interactor.Update(ctx, idUUID, &input)
	if err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, output)
}

// Delete は職種マスターの情報を1件削除する
// @Tags 職種マスター
// Businesses godoc
// @Summary 職種マスターの情報を1件削除する
// @Description 職種マスターの情報を1件削除する
// @Accept  json
// @Produce  json
// @Param id path string true "職種マスターID"
// @Success 200 {object} ports.BusinessMasterOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /businesses/{id} [delete]
func (controller *BusinessController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	idUUID, err := uuid.Parse(id) // idをUUIDに変換
	if err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	if err := controller.Interactor.Delete(ctx, idUUID); err != nil {
		return controller.Error.ErrorResponse(c, err)
	}
	return c.JSON(http.StatusOK, nil)
}
