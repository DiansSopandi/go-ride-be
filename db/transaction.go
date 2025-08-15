package db

import (
	"database/sql"
)

func RollbackOnError(tx *sql.Tx, err error) {
	if err != nil {
		_ = tx.Rollback()
	}
}
