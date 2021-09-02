module github.com/bro-n-bro/osjuno

go 1.16

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/cosmos/cosmos-sdk => github.com/osmosis-labs/cosmos-sdk v0.42.5-0.20210630232304-f792e47135c3

require (
	github.com/cosmos/cosmos-sdk v0.42.9
	github.com/desmos-labs/juno v0.0.0-20210820090829-4142e0029177
	github.com/forbole/bdjuno v0.0.0-20210902082946-0f488bf0eb30 // indirect
	github.com/osmosis-labs/osmosis v1.0.3
	github.com/tendermint/tendermint v0.34.12 // indirect
)
