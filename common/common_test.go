package common

import (
	"github.com/DSiSc/craft/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var MockHash = types.Hash{0x44, 0x49, 0xc9, 0xd9, 0xa3, 0x6a, 0x96, 0xeb, 0x28, 0xc9, 0xe1, 0x80, 0x99, 0x0, 0x5c, 0xcc, 0x65, 0x94, 0x2d, 0x5f, 0x88, 0xdd, 0x1a, 0x5a, 0x9c, 0xcf, 0xff, 0x1, 0xaa, 0x2, 0xf1, 0x76}

func MockBlock() *types.Block {
	return &types.Block{
		Header: &types.Header{
			ChainID:       1,
			PrevBlockHash: MockHash,
			StateRoot:     MockHash,
			TxRoot:        MockHash,
			ReceiptsRoot:  MockHash,
			Height:        1,
			Timestamp:     uint64(time.Date(2018, time.August, 28, 0, 0, 0, 0, time.UTC).Unix()),
		},
		Transactions: make([]*types.Transaction, 0),
	}
}

var MockBlockHash = types.Hash{0xf1, 0xc0, 0x6c, 0x80, 0xf2, 0xd2, 0xa7, 0xc6, 0x9c, 0xd6, 0x26, 0x70, 0xd5, 0xe8, 0x95, 0xcf, 0x85, 0xf8, 0xa6, 0x90, 0x1, 0x18, 0xab, 0x7c, 0xaf, 0xb0, 0x52, 0x4d, 0xc8, 0xd8, 0x9, 0x63}

func TestHeaderHash(t *testing.T) {
	block := MockBlock()
	hash := HeaderHash(block)
	assert.Equal(t, MockBlockHash, hash)
}

func TestHexToAddress(t *testing.T) {
	addHex := "333c3310824b7c685133f2bedb2ca4b8b4df633d"
	address := HexToAddress(addHex)
	b := types.Address{
		0x33, 0x3c, 0x33, 0x10, 0x82, 0x4b, 0x7c, 0x68, 0x51, 0x33,
		0xf2, 0xbe, 0xdb, 0x2c, 0xa4, 0xb8, 0xb4, 0xdf, 0x63, 0x3d,
	}
	assert.Equal(t, b, address)
}

func TestHex2Bytes(t *testing.T) {
	addHex := "333c3310824b7c685133f2bedb2ca4b8b4df633d"
	address := Hex2Bytes(addHex)
	b := []byte{
		0x33, 0x3c, 0x33, 0x10, 0x82, 0x4b, 0x7c, 0x68, 0x51, 0x33,
		0xf2, 0xbe, 0xdb, 0x2c, 0xa4, 0xb8, 0xb4, 0xdf, 0x63, 0x3d,
	}
	assert.Equal(t, b, address)
}

func TestFromHex(t *testing.T) {
	addHex := "333c3310824b7c685133f2bedb2ca4b8b4df633d"
	address := FromHex(addHex)
	b := []byte{
		0x33, 0x3c, 0x33, 0x10, 0x82, 0x4b, 0x7c, 0x68, 0x51, 0x33,
		0xf2, 0xbe, 0xdb, 0x2c, 0xa4, 0xb8, 0xb4, 0xdf, 0x63, 0x3d,
	}
	assert.Equal(t, b, address)
}
