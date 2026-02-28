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

// GenreController ジャンルマスターの構造体
type GenreController struct {
	Interactor *interactor.GenreInteractor
	Error      *presenters.ErrorPresenter
	Validator  *validator.Validate
}

// NewGenreController ジャンルマスターのコンストラクタ
func NewGenreController(interactor *interactor.GenreInteractor, logging logging.Logging, validator *validator.Validate) *GenreController {
	return &GenreController{
		Interactor: interactor,
		Error:      presenters.NewErrorPresenter(logging),
		Validator:  validator,
	}
}

// Get はジャンルマスターの情報を1件取得する
// @Tags ジャンルマスター
// Genres godoc
// @Summary ジャンルマスターから特定の１レコードを抽出する
// @Description ジャンルマスターの情報を出力し
// @Accept  json
// @Produce  json
// @Param id path string true "ジャンルマスターID"
// @Success 200 {object} ports.GenreMasterOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /genres/{id} [get]
func (controller *GenreController) Get(c echo.Context) error {
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

// List はジャンルマスターの情報をリストで取得する
// @Tags ジャンルマスター
// Genres godoc
// @Summary ジャンルマスターの情報をリストで取得する
// @Description ジャンルマスターの情報をリストで取得する
// @Accept  json
// @Produce  json
// @Success 200 {object} []ports.GenreMasterOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /genres [get]
func (controller *GenreController) List(c echo.Context) error {
	ctx := c.Request().Context()

	outputs, err := controller.Interactor.List(ctx)
	if err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, outputs)
}

// Create はジャンルマスターの情報を1件作成する
// @Tags ジャンルマスター
// Genres godoc
// @Summary ジャンルマスターの情報を1件作成する
// @Description ジャンルマスターの情報を1件作成する
// @Accept  json
// @Produce  json
// @Param wallet body ports.GenreMasterInput true "ジャンルマスター"
// @Success 200 {object} ports.GenreMasterOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /genres [post]
func (controller *GenreController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var input ports.GenreMasterInput

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

// Update はジャンルマスターの情報を1件修正する
// @Tags ジャンルマスター
// Genres godoc
// @Summary ジャンルマスターの情報を1件修正する
// @Description ジャンルマスターの情報を1件修正する
// @Accept  json
// @Produce  json
// @Param id path string true "ジャンルマスターID"
// @Param wallet body ports.GenreMasterInput true "ジャンルマスター"
// @Success 200 {object} ports.GenreMasterOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /genres/{id} [put]
func (controller *GenreController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	idUUID, err := uuid.Parse(id) // idをUUIDに変換
	if err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	var input ports.GenreMasterInput
	if err := c.Bind(&input); err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	if err := controller.Validator.Struct(&input); err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	output, err := controller.Interactor.Update(ctx, idUUID, &input)
	if err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, output)
}

// Delete はジャンルマスターの情報を1件削除する
// @Tags ジャンルマスター
// Genres godoc
// @Summary ジャンルマスターの情報を1件削除する
// @Description ジャンルマスターの情報を1件削除する
// @Accept  json
// @Produce  json
// @Param id path string true "ジャンルマスターID"
// @Success 200 {string} string "OK"
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /genres/{id} [delete]
func (controller *GenreController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	idUUID, err := uuid.Parse(id) // idをUUIDに変換
	if err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	if err := controller.Interactor.Delete(ctx, idUUID); err != nil {
		return controller.Error.ErrorResponse(c, err)
	}
	return c.JSON(http.StatusOK, "OK")
}
