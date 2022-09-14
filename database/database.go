package database

import "database/sql"

func Query(db *sql.DB) (err error) {
	rows, err := db.Query("SELECT 1 FROM PRODUCT WHERE USER_ID = ?", 0)
	if err == nil {
		defer rows.Close()
	}
	return
}

func Execute(db *sql.DB, userID, productID int64) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()
	if _, err = tx.Exec("UPDATE PRODUCT SET COUNT = COUNT + 1"); err != nil {
		return
	}
	if _, err = tx.Exec("INSERT INTO PRODUCT (USER_ID, PRODUCT_ID) VALUES (?, ?)", userID, productID); err != nil {
		return
	}
	return
}
