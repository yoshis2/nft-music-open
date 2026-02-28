// Package interactor は、ビジネスロジックを実装します。
package interactor

import (
	"context"
	"testing"

	"nft-music/domain"
	"nft-music/usecases/gateways/mock"
	"nft-music/usecases/ports"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestIpfsInteractor_MetaJSON(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockIpfsGateway := mock.NewMockIpfsGateway(ctrl)
	mockUserGateway := mock.NewMockUserGateway(ctrl)
	interactor := NewIpfsInteractor(mockIpfsGateway, mockUserGateway)

	t.Run("正常系: メタデータをアップロードできる", func(t *testing.T) {
		input := ports.IpfsMetaInput{
			Name: "NFT Name",
		}

		mockIpfsGateway.EXPECT().
			Add(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(&domain.IpfsAdd{Hash: "QmMetaHash"}, nil)

		mockIpfsGateway.EXPECT().
			Publish(gomock.Any(), "QmMetaHash").
			Return(&domain.IpfsPublish{Name: "QmMetaName"}, nil)

		mockIpfsGateway.EXPECT().
			Localpin(gomock.Any(), "QmMetaHash").
			Return(&domain.IpfsPins{Pins: []string{"QmMetaHash"}}, nil)

		mockIpfsGateway.EXPECT().
			Resolve(gomock.Any(), "QmMetaName").
			Return(&domain.IpfsResolve{Path: "/ipfs/QmMetaHash"}, nil)

		output, err := interactor.MetaJSON(context.Background(), input)

		assert.NoError(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, "QmMetaHash", output.Cid)
	})
}
