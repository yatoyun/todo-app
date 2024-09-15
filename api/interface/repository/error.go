package repository

import (
	"database/sql"
	"fmt"
)

func checkRowsAffected(res sql.Result) error {
    rowsAffected, err := res.RowsAffected()
    if err != nil {
        return err
    }
    if rowsAffected != 1 {
        return fmt.Errorf("weird  behavior. total affected: %d", rowsAffected)
    }
    return nil
}
