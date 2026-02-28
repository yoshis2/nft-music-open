// Package interactor は、ビジネスロジックを実装します。
package interactor

import (
	"context"
	"testing"

	"nft-music/adapters/gateways"
	"nft-music/infrastructure/mysql"
	"nft-music/usecases/ports"

	"github.com/stretchr/testify/assert"
)

func TestGenreInteractor_CreateAndGet(t *testing.T) {
	// Setup test database
	newMysql := mysql.NewTMysql()
	client := newMysql.TestOpen()
	genreGateway := gateways.NewGenreGateway(client)
	interactor := NewGenreInteractor(genreGateway)

	// Input data for creating a genre
	input := &ports.GenreMasterInput{
		Name: "Test Genre",
	}

	// Call Create method
	ctx := context.Background()
	createdGenre, err := interactor.Create(ctx, input)
	assert.NoError(t, err)
	assert.NotNil(t, createdGenre)

	// Assert the created genre's data
	assert.Equal(t, input.Name, createdGenre.Name)

	// Call Get method with the created genre's ID
	retrievedGenre, err := interactor.Get(ctx, createdGenre.ID)
	assert.NoError(t, err)
	assert.NotNil(t, retrievedGenre)
	assert.Equal(t, createdGenre.ID, retrievedGenre.ID)
	assert.Equal(t, createdGenre.Name, retrievedGenre.Name)

	// Clean up
	err = interactor.Delete(ctx, createdGenre.ID)
	assert.NoError(t, err)
}
