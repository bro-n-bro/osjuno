package gamm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/juno/modules"
	junotypes "github.com/desmos-labs/juno/types"
	gammtypes "github.com/osmosis-labs/osmosis/x/gamm/types"
)

var (
	_ modules.Module        = &Module{}
	_ modules.MessageModule = &Module{}
)

type Module struct {
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) Name() string {
	return "osmosis/gamm"
}

func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *junotypes.Tx) error {

	switch osmosisMsg := msg.(type) {
	case *gammtypes.MsgSwapExactAmountIn:
		return m.handleSwapExactAmountIn(index, osmosisMsg, tx)
	}
	return nil
}
