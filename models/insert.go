package models

import "currency-API/db"

func Insert(currency Currency) (id int64, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO currency (id,name, amount) VALUES ($1, S3, $2)RETURNING id`

	err = conn.QueryRow(sql, currency.Id, currency.Name, currency.Amount).Scan(&id)

	return

}
