package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var chDatabase chan interface{}

func (p actRecord) Do(conn *sql.DB) error {
	_, err := conn.Exec(`insert into sign(user_id,user_name,state)
						values(?,?,?)`, p.Dat.ID, p.Dat.Username, p.Dat.State)
	return err
}

type queryState struct {
	State string
}

type queryResult struct {
	Ok   bool   `json:"ok"`
	Name string `json:"name"`
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
			_ = act.Do(dbConn)
		case queryState:
			res, err := dbConn.Query(`select user_name from sign where date(sign_time)=curdate() and state=?;`, act.State)
			defer res.Close()
			if err != nil {
				chDatabase <- queryResult{
					Ok:   false,
					Name: "#a",
				}
				break
			}
			hasResult := false
			var name string
			if res.Next() {
				res.Scan(&name)
				hasResult = true
			}

			if hasResult {
				chDatabase <- queryResult{
					Ok:   true,
					Name: name,
				}

			} else {
				chDatabase <- queryResult{
					Ok:   false,
					Name: "#b",
				}
			}

		default:
			log.Fatalf("Unexcepted type(%T) through channel.", thisCase)
			// chDatabase <- fmt.Errorf("MODULE DB PANIC: A UNEXCEPTED TYPE FOUND")
		}
	}
}
