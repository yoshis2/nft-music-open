// Package interactor は、ビジネスロジックを実装します。
package interactor

import (
	"context"
	"testing"

	"nft-music/domain"
	"nft-music/usecases/gateways/mock"
	"nft-music/usecases/logging"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type NullLogging struct{}

func (n *NullLogging) Close()                              {}
func (n *NullLogging) Error(s string)                      {}
func (n *NullLogging) Warning(s string)                    {}
func (n *NullLogging) Info(s string)                       {}
func (n *NullLogging) Debug(s string)                      {}
func (n *NullLogging) AccessLog(e *logging.AccessLogEntry) {}
func (n *NullLogging) SQLLog(a, b, c string)               {}

func TestNftInteractor_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserGateway := mock.NewMockUserGateway(ctrl)
	mockTransactionGateway := mock.NewMockTransactionGateway(ctrl)
	mockIpfsGateway := mock.NewMockIpfsGateway(ctrl)
	mockLogging := &NullLogging{}

	interactor := &NftInteractor{
		UserGateway:        mockUserGateway,
		TransactionGateway: mockTransactionGateway,
		IpfsGateway:        mockIpfsGateway,
		Logging:            mockLogging,
	}

	t.Run("正常系: NFT一覧を取得できる", func(t *testing.T) {
		transactionID := "0x123"
		userID := uuid.New()

		expectedTransactions := []*domain.Transaction{
			{ID: transactionID, UserID: userID, TokenURL: "QmToken"},
		}

		mockTransactionGateway.EXPECT().
			List(gomock.Any(), 10).
			Return(expectedTransactions, nil)

		mockIpfsGateway.EXPECT().
			Get(gomock.Any(), "QmToken").
			Return(&domain.IpfsJSON{Name: "NFT Name", Description: "Desc"}, nil)

		outputs, err := interactor.List(context.Background(), 10)

		assert.NoError(t, err)
		assert.Equal(t, 1, len(outputs))
		assert.Equal(t, "NFT Name", outputs[0].Name)
	})
}
