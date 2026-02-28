// Package ethereum は、イーサリアムネットワークとの通信を管理します。
package ethereum

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"nft-music/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumVirtualMachine struct{}

func NewEthereumVirtualMachine() *EthereumVirtualMachine {
	return &EthereumVirtualMachine{}
}

func (evm *EthereumVirtualMachine) Auth() (*ethclient.Client, *bind.TransactOpts, *contracts.Contracts, error) {
	ctx := context.Background()
	// 環境変数からGanacheのURLを取得。なければデフォルト値を使用
	ganacheURL := os.Getenv("GANACHE_URL")
	if ganacheURL == "" {
		ganacheURL = "http://ganache:8545" // GanacheのデフォルトRPCサーバー
	}
	client, err := ethclient.Dial(ganacheURL)
	if err != nil {
		return nil, nil, nil, err
	}

	// TODO: シークレットキーをどのように取得するか
	privateKey, err := crypto.HexToECDSA("01a3c379a8b07fb8f6ae7735fd7de35c5156c1cb8c3ac76a9d5157e3eedd2c4a")
	if err != nil {
		return nil, nil, nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, nil, nil, err
	}

	gasPrice := big.NewInt(20000000000) // 20 Gwei

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, nil, nil, err
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, nil, nil, err
	}

	transactOpts.Nonce = big.NewInt(int64(nonce))
	transactOpts.Value = big.NewInt(0)      // in wei
	transactOpts.GasLimit = uint64(3000000) // in units
	transactOpts.GasPrice = gasPrice

	name := "test"
	symbol := "tst"
	initialListingPrice := big.NewInt(1000000000000000)
	initialRoyaltyFeeBps := big.NewInt(1000)

	address, tx, instance, err := contracts.DeployContracts(transactOpts, client, name, symbol, initialListingPrice, initialRoyaltyFeeBps)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to deploy contracts: %w", err)
	}

	fmt.Printf("tx: %v\n", tx)
	fmt.Printf("instance: %v\n", instance)
	contracts, err := contracts.NewContracts(address, client)
	if err != nil {
		return nil, nil, nil, err
	}
	return client, transactOpts, contracts, nil
}
