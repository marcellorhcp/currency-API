package models

import "currency-API/db"

func Update(id int64, currency Currency) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE currency SET name=$1, amount=$2, WHERE id=$3`, currency.Name, currency.Amount, currency.Id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
