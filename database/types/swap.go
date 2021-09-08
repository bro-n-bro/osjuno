package types

import (
	"database/sql/driver"
	"fmt"
	"github.com/bro-n-bro/osjuno/models"
	bdjunodbtypes "github.com/forbole/bdjuno/database/types"
)

type DbPoolSwap struct {
	PoolId  uint64
	InCoin  bdjunodbtypes.DbCoin
	OutCoin bdjunodbtypes.DbCoin
}

func (dbPoolSwap DbPoolSwap) Value() (driver.Value, error) {
	return fmt.Sprintf(
		`(%v,%s,%s,%s,%s)`,
		dbPoolSwap.PoolId, dbPoolSwap.InCoin.Denom, dbPoolSwap.InCoin.Amount,
		dbPoolSwap.OutCoin.Denom, dbPoolSwap.OutCoin.Amount,
	), nil
}

type DbPoolsSwaps []DbPoolSwap

func NewDbPoolSwaps(poolsSwaps []models.PoolSwap) DbPoolsSwaps {
	dbPoolsSwaps := make([]DbPoolSwap, 0)
	for _, poolSwap := range poolsSwaps {
		dbPoolsSwaps = append(dbPoolsSwaps, DbPoolSwap{
			PoolId:  poolSwap.PoolId,
			InCoin:  bdjunodbtypes.NewDbCoin(poolSwap.InCoin),
			OutCoin: bdjunodbtypes.NewDbCoin(poolSwap.OutCoin),
		})
	}
	return dbPoolsSwaps
}
