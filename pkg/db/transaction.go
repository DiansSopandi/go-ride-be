package db

import (
	"database/sql"
	"fmt"
)

func RollbackOnError(tx *sql.Tx, err error) {
	if err != nil {
		fmt.Println("error create user db rollback", err)
		_ = tx.Rollback()
	}
}
