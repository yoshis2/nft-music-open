// Package gatewaysは、データベースや外部サービスへのアクセスを実装します。
package gateways

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"nft-music/domain"

	"gorm.io/gorm"
)

type IpfsGateway struct {
	Database *gorm.DB
}

func NewIpfsGateway(db *gorm.DB) *IpfsGateway {
	return &IpfsGateway{Database: db}
}

func (gateway *IpfsGateway) Get(ctx context.Context, cid string) (*domain.IpfsJSON, error) {
	client := http.Client{}
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://ipfs:8080/"+cid, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var ipfsJSON domain.IpfsJSON
	if err := json.Unmarshal(respBody, &ipfsJSON); err != nil {
		return nil, err
	}

	return &ipfsJSON, nil
}

func (gateway *IpfsGateway) Add(ctx context.Context, body *bytes.Buffer, contentType string) (*domain.IpfsAdd, error) {
	path := os.Getenv("IPFS_HOST") + os.Getenv("IPFS_API_PORT") + "/api/v0/add"

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, path, body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", contentType)
	request.Header.Set("Content-Disposition", `form-data; name="file"; filename="file"`)

	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(response.Body)
		return nil, fmt.Errorf("API returned non-OK status: %s, body: %s", response.Status, string(respBody))
	}

	respBody, _ := io.ReadAll(response.Body)

	var ipfsAdd domain.IpfsAdd
	if err := json.Unmarshal(respBody, &ipfsAdd); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return &ipfsAdd, nil
}

func (gateway *IpfsGateway) Publish(ctx context.Context, cid string) (*domain.IpfsPublish, error) {
	path := os.Getenv("IPFS_HOST") + os.Getenv("IPFS_API_PORT") + "/api/v0/name/publish"

	data := url.Values{}
	data.Set("arg", cid)

	respBody, err := apiRequest(ctx, path, data)
	if err != nil {
		return nil, err
	}

	var ipfsPublish domain.IpfsPublish
	if err := json.Unmarshal(respBody, &ipfsPublish); err != nil {
		return nil, err
	}

	return &ipfsPublish, nil
}

func (gateway *IpfsGateway) Localpin(ctx context.Context, cid string) (*domain.IpfsPins, error) {
	path := os.Getenv("IPFS_HOST") + os.Getenv("IPFS_API_PORT") + "/api/v0/pin/add"
	data := url.Values{}
	data.Set("arg", cid)

	respBody, err := apiRequest(ctx, path, data)
	if err != nil {
		return nil, err
	}

	var ipfsPins domain.IpfsPins
	if err := json.Unmarshal(respBody, &ipfsPins); err != nil {
		return nil, err
	}
	return &ipfsPins, nil
}

func (gateway *IpfsGateway) Resolve(ctx context.Context, name string) (*domain.IpfsResolve, error) {
	path := os.Getenv("IPFS_HOST") + os.Getenv("IPFS_API_PORT") + "/api/v0/name/resolve"

	data := url.Values{}
	data.Set("arg", name)

	respBody, err := apiRequest(ctx, path, data)
	if err != nil {
		return nil, err
	}

	var ipfsResolve domain.IpfsResolve
	if err := json.Unmarshal(respBody, &ipfsResolve); err != nil {
		return nil, err
	}
	return &ipfsResolve, nil
}

func apiRequest(ctx context.Context, path string, data url.Values) ([]byte, error) {
	dataPath := fmt.Sprintf("%s?%s", path, data.Encode())

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, dataPath, bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	closeErr := response.Body.Close()
	if closeErr != nil {
		return nil, closeErr
	}

	return respBody, nil
}
