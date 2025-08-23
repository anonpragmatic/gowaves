package state

import (
	"testing"

	"github.com/anonpragmatic/gowaves/pkg/crypto"
	"github.com/stretchr/testify/assert"
)

func TestFeeDistributionSerialization(t *testing.T) {
	asset0, err := crypto.NewDigestFromBase58(assetStr)
	assert.NoError(t, err, "NewDigestFromBase58() failed")
	asset1, err := crypto.NewDigestFromBase58(assetStr1)
	assert.NoError(t, err, "NewDigestFromBase58() failed")
	distr := feeDistribution{
		wavesFeeDistribution{100, 500},
		assetsFeeDistribution{assetFeeMap{asset0: 3, asset1: 42424242}, assetFeeMap{asset0: 2, asset1: 42}},
	}
	distrBytes := distr.marshalBinary()
	distr2 := newFeeDistribution()
	err = distr2.unmarshalBinary(distrBytes)
	assert.NoError(t, err, "unmarshalBinary() failed")
	assert.Equal(t, distr, distr2)
}
