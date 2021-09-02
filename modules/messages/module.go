package messages

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	junomessages "github.com/desmos-labs/juno/modules/messages"
)

func OsmosisMessageAddressesParser(cdc codec.Marshaler, osmoMsg sdk.Msg) ([]string, error) {
	return nil, junomessages.MessageNotSupported(osmoMsg)
}
