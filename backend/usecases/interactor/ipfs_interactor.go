// Package interactor は、ビジネスロジックを実装します。
package interactor

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"

	"nft-music/domain"
	"nft-music/usecases/gateways"
	"nft-music/usecases/ports"
)

// IpfsInteractor はIPFSのユースケースです
type IpfsInteractor struct {
	IpfsGateway gateways.IpfsGateway
	UserGateway gateways.UserGateway
}

func NewIpfsInteractor(ipfsGateway gateways.IpfsGateway, userGateway gateways.UserGateway) *IpfsInteractor {
	return &IpfsInteractor{
		IpfsGateway: ipfsGateway,
		UserGateway: userGateway,
	}
}

// Upload はIpfsにデータをアップロードする
func (interactor *IpfsInteractor) Upload(ctx context.Context, header *multipart.FileHeader, form ports.IpfsInput) (ipfsOutput *ports.IpfsOutput, err error) {
	user := &domain.User{
		Wallet: form.Wallet,
	}
	user, err = interactor.UserGateway.GetByWallet(ctx, user)
	if err != nil {
		return nil, err
	}
	file, err := header.Open()
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", header.Filename)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(part, file); err != nil {
		return nil, err
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	ipfsAdd, err := interactor.IpfsGateway.Add(ctx, body, writer.FormDataContentType())
	if err != nil {
		return nil, err
	}

	// public
	ipfsOutput, err = public(ctx, interactor.IpfsGateway, ipfsAdd.Hash)
	if err != nil {
		return nil, err
	}

	ipfsOutput.UserID = user.ID

	return ipfsOutput, nil
}

func (interactor *IpfsInteractor) MetaJSON(ctx context.Context, input ports.IpfsMetaInput) (*ports.IpfsOutput, error) {
	metaJSON, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile("file", "meta.json") // "file" は必須、"data.json" は任意のファイル名
	if err != nil {
		return nil, err
	}

	_, err = part.Write(metaJSON)
	if err != nil {
		return nil, err
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	ipfsAdd, err := interactor.IpfsGateway.Add(ctx, &body, writer.FormDataContentType())
	if err != nil {
		return nil, err
	}

	return public(ctx, interactor.IpfsGateway, ipfsAdd.Hash)
}

func public(ctx context.Context, gateways gateways.IpfsGateway, hash string) (*ports.IpfsOutput, error) {
	ipfsPublish, err := gateways.Publish(ctx, hash)
	if err != nil {
		return nil, err
	}
	_, err = gateways.Localpin(ctx, hash)
	if err != nil {
		return nil, err
	}

	ipfsResolve, err := gateways.Resolve(ctx, ipfsPublish.Name)
	if err != nil {
		return nil, err
	}

	ipfsOutput := &ports.IpfsOutput{
		Cid:  hash,
		Path: ipfsResolve.Path,
	}

	return ipfsOutput, nil
}
