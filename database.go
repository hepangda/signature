package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var chDatabase chan interface{}

func (p actRecord) Do(conn *sql.DB) error {
	_, err := conn.Exec(`insert into sign(user_id,user_name)
						values(?,?)`, p.Dat.ID, p.Dat.Username)
	return err
}

func dbDistributor() {
	dbConn, err := sql.Open("mysql", databaseString)
	if err != nil {
		chDatabase <- err
		return
	}
	defer dbConn.Close()

	for {
		thisCase := <-chDatabase

		switch act := thisCase.(type) {
		case action:
			err := act.Do(dbConn)
			if err != nil {
				err.Error()
			} //prevent error
			// chDatabase <- err
		default:
			log.Fatalf("Unexcepted type(%T) through channel.", thisCase)
			// chDatabase <- fmt.Errorf("MODULE DB PANIC: A UNEXCEPTED TYPE FOUND")
		}
	}
}
