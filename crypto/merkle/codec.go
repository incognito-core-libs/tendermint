package merkle

import (
	amino "github.com/incognito-core-libs/go-amino"
)

var cdc *amino.Codec

func init() {
	cdc = amino.NewCodec()
	cdc.Seal()
}
