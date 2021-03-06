package ethereum

import (
	"testing"

	"github.com/0xProject/0x-mesh/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetContractAddressesForChainID(t *testing.T) {
	// a valid chainId returns no error
	_, err := GetContractAddressesForChainID(constants.TestChainID)
	require.NoError(t, err)

	// an invalid chainId returns an error stating the desired chain id
	_, err = GetContractAddressesForChainID(-1)
	assert.EqualError(t, err, "invalid chain: -1")
}
