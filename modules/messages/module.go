package messages

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	junomessages "github.com/desmos-labs/juno/modules/messages"
	gammtypes "github.com/osmosis-labs/osmosis/x/gamm/types"
)

func OsmosisMessageAddressesParser(cdc codec.Marshaler, osmoMsg sdk.Msg) ([]string, error) {
	switch msg := osmoMsg.(type) {
	case *gammtypes.MsgSwapExactAmountIn:
		return []string{msg.Sender}, nil
	}
	return nil, junomessages.MessageNotSupported(osmoMsg)
}
