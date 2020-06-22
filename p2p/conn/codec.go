package conn

import (
	amino "github.com/incognito-core-libs/go-amino"
	cryptoAmino "github.com/incognito-core-libs/tendermint/crypto/encoding/amino"
)

var cdc *amino.Codec = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(cdc)
	RegisterPacket(cdc)
}
