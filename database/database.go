package database

import (
	"github.com/desmos-labs/juno/db/postgresql"
	"github.com/jmoiron/sqlx"
)

type OsmosisDb struct {
	*postgresql.Database
	Sqlx *sqlx.DB
}
