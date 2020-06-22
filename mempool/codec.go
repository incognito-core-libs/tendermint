package mempool

import (
	amino "github.com/incognito-core-libs/go-amino"
)

var cdc = amino.NewCodec()

func init() {
	RegisterMempoolMessages(cdc)
}
