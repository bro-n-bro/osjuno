package modules

import (
	"github.com/desmos-labs/juno/modules"
	"github.com/desmos-labs/juno/modules/messages"
	junoRegistrar "github.com/desmos-labs/juno/modules/registrar"
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
	return r.bdjunoRegistrar.BuildModules(ctx)
}