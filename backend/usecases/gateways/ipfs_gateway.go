// Package gateways は、データベースや外部サービスへのアクセスを実装します。
package gateways

import (
	"bytes"
	"context"

	"nft-music/domain"
)

// IpfsGateway はIPFSノードにアクセスする処理です
//
//go:generate mockgen -package mock -source $GOFILE -destination mock/$GOFILE
type IpfsGateway interface {
	Get(ctx context.Context, cid string) (*domain.IpfsJSON, error)
	Add(ctx context.Context, body *bytes.Buffer, contentType string) (*domain.IpfsAdd, error)
	Publish(ctx context.Context, cid string) (*domain.IpfsPublish, error)
	Localpin(ctx context.Context, cid string) (*domain.IpfsPins, error)
	Resolve(ctx context.Context, name string) (*domain.IpfsResolve, error)
}
