package main

import (
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/desmos-labs/juno/cmd"
	parsecmd "github.com/desmos-labs/juno/cmd/parse"
	"github.com/desmos-labs/juno/modules/messages"
	"github.com/forbole/bdjuno/database"
	"github.com/forbole/bdjuno/modules"
	"github.com/forbole/bdjuno/types/config"

	osmosismessages "github.com/bro-n-bro/osjuno/modules/messages"
	osmosisapp "github.com/osmosis-labs/osmosis/app"
)

func main() {

	parseCfg := parsecmd.NewConfig().
		WithDBBuilder(database.Builder).
		WithConfigParser(config.Parser).
		WithEncodingConfigBuilder(config.MakeEncodingConfig(getBasicManagers())).
		WithRegistrar(modules.NewRegistrar(getAddressesParser()))

	cfg := cmd.NewConfig("osjuno").WithParseConfig(parseCfg)

	executor := cmd.BuildDefaultExecutor(cfg)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

// getBasicManagers returns the various basic managers that are used to register the encoding to
// support custom messages.
// This should be edited by custom implementations if needed.
func getBasicManagers() []module.BasicManager {
	return []module.BasicManager{
		simapp.ModuleBasics,
		osmosisapp.ModuleBasics,
	}
}

func getAddressesParser() messages.MessageAddressesParser {
	return messages.JoinMessageParsers(
		osmosismessages.OsmosisMessageAddressesParser,
		messages.CosmosMessageAddressesParser,
	)
}
