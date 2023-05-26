package models

import "currency-API/db"

func Get(id int64) (currency Currency, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM currencies WHERE id=$1`, id)

	err = row.Scan(&currency.Id, &currency.Name, &currency.Amount)

	return

}
