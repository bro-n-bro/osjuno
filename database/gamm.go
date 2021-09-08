package database

import (
	"github.com/bro-n-bro/osjuno/database/types"
	"github.com/bro-n-bro/osjuno/models"
	junotypes "github.com/desmos-labs/juno/types"
	"github.com/lib/pq"
)

var queryInsertSwap = `
		INSERT INTO swap (transaction_hash, sender, routes)
		VALUES ($1, $2, $3::POOL_SWAP[])
	`

func (db *OsmosisDb) SaveChainedSwap(tx *junotypes.Tx, swap models.ChainedSwap) error {
	dbPoolsSwaps := types.NewDbPoolSwaps(swap.PoolsSwaps)
	_, err := db.Sql.Exec(queryInsertSwap, tx.TxHash, swap.Sender, pq.Array(dbPoolsSwaps))
	return err
}
