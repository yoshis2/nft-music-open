// Package controllersは、HTTPリクエストのハンドリングとレスポンス制御を実装します。
package controllers

import (
	"log"
	"net/http"

	"nft-music/adapters/presenters"
	"nft-music/usecases/interactor"
	"nft-music/usecases/logging"
	"nft-music/usecases/ports"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CollectionController コレクションのコントローラー
type CollectionController struct {
	Interactor *interactor.CollectionInteractor
	Error      *presenters.ErrorPresenter
	Validator  *validator.Validate
}

// NewCollectionController コレクションのコントローラーのコンストラクタ
func NewCollectionController(interactor *interactor.CollectionInteractor, logging logging.Logging, validator *validator.Validate) *CollectionController {
	return &CollectionController{
		Interactor: interactor,
		Error:      presenters.NewErrorPresenter(logging),
		Validator:  validator,
	}
}

// Get はコレクションを取得する
// @Tags コレクション
// @Summary コレクションの情報を1件取得する
// @Description コレクションの情報を1件取得する
// @Accept  json
// @Produce  json
// @Param id path string true "コレクションID"
// @Success 200 {object} ports.CollectionOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /collections/{id} [get]
func (controller *CollectionController) Get(c echo.Context) error {
	ctx := c.Request().Context()
	id := uuid.MustParse(c.Param("id"))

	output, err := controller.Interactor.Get(ctx, id)
	if err != nil {
		return controller.Error.ErrorResponse(c, err)
	}
	return c.JSON(http.StatusOK, output)
}

// List はコレクションをリストで取得する
// @Tags コレクション
// @Summary コレクションの情報をリストで取得する
// @Description コレクションの情報をリストで取得する
// @Accept  json
// @Produce  json
// @Success 200 {object} []ports.CollectionOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /collections [get]
func (controller *CollectionController) List(c echo.Context) error {
	ctx := c.Request().Context()

	output, err := controller.Interactor.List(ctx)
	if err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, output)
}

// Create はコレクションを作成する
// @Tags コレクション
// @Summary コレクションの情報を1件作成する
// @Description コレクションの情報を1件作成する
// @Accept  json
// @Produce  json
// @Param collection body ports.CollectionInput true "コレクション"
// @Success 200 {object} ports.CollectionOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /collections [post]
func (controller *CollectionController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var input ports.CollectionInput
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

	log.Println("outputのIDの中身")
	log.Println(output.ID)

	return c.JSON(http.StatusOK, output)
}

// Update はコレクションを更新する
// @Tags コレクション
// @Summary コレクションの情報を1件修正する
// @Description コレクションの情報を1件修正する
// @Accept  json
// @Produce  json
// @Param id path string true "コレクションID"
// @Param collection body ports.CollectionInput true "コレクション"
// @Success 200 {object} ports.CollectionOutput
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /collections/{id} [put]
func (controller *CollectionController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	idUUID, err := uuid.Parse(id) // idをUUIDに変換
	if err != nil {
		return err
	}

	var input ports.CollectionInput
	if err := c.Bind(&input); err != nil {
		return controller.Error.ErrorResponse(c, err)
	}

	output, err := controller.Interactor.Update(ctx, idUUID, &input)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, output)
}

// Delete はコレクションを削除する
// @Tags コレクション
// @Summary コレクションの情報を1件削除する
// @Description コレクションの情報を1件削除する
// @Accept  json
// @Produce  json
// @Param id path string true "コレクションID"
// @Success 200
// @Failure 400 {object} ports.ErrorResponseObject
// @Failure 404 {object} ports.ErrorResponseObject
// @Failure 500 {object} ports.ErrorResponseObject
// @Router /collections/{id} [delete]
func (controller *CollectionController) Delete(c echo.Context) error {
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
