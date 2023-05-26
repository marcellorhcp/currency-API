package models

import "currency-API/db"

func GetAll() (currencies []Currency, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM currencies`)
	if err != nil {
		return
	}

	for rows.Next() {
		var currency Currency

		err = rows.Scan(&currency.Id, &currency.Name, &currency.Amount)
		if err != nil {
			continue
		}

		currencies = append(currencies, currency)
	}

	return
}
