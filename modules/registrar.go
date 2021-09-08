package modules

import (
	"github.com/bro-n-bro/osjuno/database"
	"github.com/bro-n-bro/osjuno/modules/gamm"
	"github.com/desmos-labs/juno/modules"
	"github.com/desmos-labs/juno/modules/messages"
	junoRegistrar "github.com/desmos-labs/juno/modules/registrar"
	bdjunodb "github.com/forbole/bdjuno/database"
	bdjunoModule "github.com/forbole/bdjuno/modules"
)

var (
	_ junoRegistrar.Registrar = &Registrar{}
)

func NewRegistrar(parser messages.MessageAddressesParser) *Registrar {
	return &Registrar{
		bdjunoRegistrar: bdjunoModule.NewRegistrar(parser),
	}
}

type Registrar struct {
	bdjunoRegistrar *bdjunoModule.Registrar
}

func (r *Registrar) BuildModules(ctx junoRegistrar.Context) modules.Modules {
	cosmosMoules := r.bdjunoRegistrar.BuildModules(ctx)
	osmosisModules := r.buildOsmosisMoules(ctx)
	return append(cosmosMoules, osmosisModules...)
}

func (r *Registrar) buildOsmosisMoules(ctx junoRegistrar.Context) modules.Modules {
	bigDipperBd := bdjunodb.Cast(ctx.Database)
	osmosisDb := &database.OsmosisDb{
		Database: bigDipperBd.Database,
		Sqlx:     bigDipperBd.Sqlx,
	}
	return []modules.Module{
		gamm.NewModule(osmosisDb),
	}
}
