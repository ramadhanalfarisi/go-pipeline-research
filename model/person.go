package model

import (
	"database/sql"
)

type PersonModel struct {
	DB *sql.DB
}

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Contact string `json:"contact"`
	Gender  string `json:"gender"`
}

func (p *PersonModel) InsertPerson(persons chan Person) error {
	tx, err := p.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // The rollback will be ignored if the tx has been committed later in the function.

	stmt, err := tx.Prepare("INSERT INTO persons(id, name, age, contact, gender) VALUES( ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	for person := range persons {
		if _, err := stmt.Exec(nil, person.Name, person.Age, person.Contact, person.Gender); err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
