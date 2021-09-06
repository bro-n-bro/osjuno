package gamm

import (
	"errors"
	"github.com/bro-n-bro/osjuno/models"
	sdk "github.com/cosmos/cosmos-sdk/types"
	junotypes "github.com/desmos-labs/juno/types"
	gammtypes "github.com/osmosis-labs/osmosis/x/gamm/types"
)

func (m *Module) handleSwapExactAmountIn(index int, msg *gammtypes.MsgSwapExactAmountIn, tx *junotypes.Tx) error {
	if tx.Code != 0 {
		return nil
	}
	_, err := getChainedSwapFromTxLogs(msg, tx.Logs)
	return err
}

func getChainedSwapFromTxLogs(
	msg *gammtypes.MsgSwapExactAmountIn, txLogs sdk.ABCIMessageLogs,
) (*models.ChainedSwap, error) {

	transferEvent := findTransferEvent(txLogs)
	if transferEvent == nil {
		return nil, errors.New("could not extract swap results")
	}
	poolsSwaps, err := transferEventToPoolsSwaps(msg, transferEvent)
	if err != nil {
		return nil, err
	}

	swap := models.ChainedSwap{
		Sender:     msg.Sender,
		PoolsSwaps: poolsSwaps,
	}
	return &swap, nil
}

func findTransferEvent(txLogs sdk.ABCIMessageLogs) *sdk.StringEvent {
	for _, log := range txLogs {
		for _, event := range log.Events {
			if event.Type == "transfer" {
				return &event
			}
		}
	}
	return nil
}

// each route swap represents by 6 event attributes
//  0 pool address
//  1 sender address
//  2 in coin-amount
//  3 sender address
//  4 pool address
//  5 out coin-amount
func transferEventToPoolsSwaps(
	msg *gammtypes.MsgSwapExactAmountIn, event *sdk.StringEvent,
) ([]models.PoolSwap, error) {

	// defence check
	if len(msg.Routes)*6 != len(event.Attributes) {
		return nil, errors.New("event do not contain all attributes to parse routes")
	}

	var swaps []models.PoolSwap
	for routeIndex, route := range msg.Routes {
		inCoinAsStr := event.Attributes[routeIndex*6+2].Value
		outCoinAsStr := event.Attributes[routeIndex*6+5].Value

		inCoin, err := models.NewCoinFromStr(inCoinAsStr)
		if err != nil {
			return nil, err
		}
		outCoin, err := models.NewCoinFromStr(outCoinAsStr)
		if err != nil {
			return nil, err
		}
		swap := models.NewPoolSwap(route.PoolId, *inCoin, *outCoin)
		swaps = append(swaps, swap)
	}
	return swaps, nil
}
