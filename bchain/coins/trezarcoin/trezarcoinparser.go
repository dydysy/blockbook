package trezarcoin

import (
	"github.com/dydysy/blockbook/bchain"
	"github.com/dydysy/blockbook/bchain/coins/btc"
	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
)

// magic numbers
const (
	MainnetMagic wire.BitcoinNet = 0xfddbefe4
)

// chain parameters
var (
	MainNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{66}
	MainNetParams.ScriptHashAddrID = []byte{8}
}

// TrezarcoinParser handle
type TrezarcoinParser struct {
	*btc.BitcoinParser
	baseparser *bchain.BaseParser
}

// NewTrezarcoinParser returns new TrezarcoinParser instance
func NewTrezarcoinParser(params *chaincfg.Params, c *btc.Configuration) *TrezarcoinParser {
	return &TrezarcoinParser{
		BitcoinParser: btc.NewBitcoinParser(params, c),
		baseparser:    &bchain.BaseParser{},
	}
}

// GetChainParams returns network parameters
func GetChainParams(chain string) *chaincfg.Params {
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err != nil {
			panic(err)
		}
	}
	switch chain {
	default:
		return &MainNetParams
	}
}

// PackTx packs transaction to byte array using protobuf
func (p *TrezarcoinParser) PackTx(tx *bchain.Tx, height uint32, blockTime int64) ([]byte, error) {
	return p.baseparser.PackTx(tx, height, blockTime)
}

// UnpackTx unpacks transaction from protobuf byte array
func (p *TrezarcoinParser) UnpackTx(buf []byte) (*bchain.Tx, uint32, error) {
	return p.baseparser.UnpackTx(buf)
}
