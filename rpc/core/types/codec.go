package core_types

import (
	amino "github.com/incognito-core-libs/go-amino"
	"github.com/incognito-core-libs/tendermint/types"
)

func RegisterAmino(cdc *amino.Codec) {
	types.RegisterEventDatas(cdc)
	types.RegisterBlockAmino(cdc)
}
