// Package interactor は、ビジネスロジックを実装します。
package interactor

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvmInteractor_Signer(t *testing.T) {
	// This test requires a running Ganache instance, as defined in the docker-compose.yml file.

	// Create an EvmInteractor
	interactor := NewEvmInteractor(nil) // Logging is not used in the test

	// Call the Signer method
	ctx := context.Background()
	output, err := interactor.Signer(ctx)

	// Assert that no error occurred
	assert.NoError(t, err)

	// Assert that the returned values are not empty
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.AddressHex)
	assert.NotEmpty(t, output.TransactionHash)
}
