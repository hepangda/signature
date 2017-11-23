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

type queryRank struct {
}

type rankOne struct {
	Rank int    `json:"rank"`
	Name string `json:"name"`
	Time string `json:"time"`
}

type resultRank struct {
	Numbers int       `json:"num"`
	Res     []rankOne `json:"res"`
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
		case queryRank:
			res, err := dbConn.Query(`select user_name,sign_time from sign where date(sign_time)=curdate();`)
			defer res.Close()
			if err != nil {
				log.Println(err.Error())
				chDatabase <- resultRank{
					Numbers: 0,
				}
				break
			}

			var ret resultRank
			ret.Numbers = 0
			log.Print(1)
			for res.Next() {
				log.Print(1)
				var name, time string
				res.Scan(&name, &time)
				ret.Numbers++
				ret.Res = append(ret.Res, rankOne{
					Rank: ret.Numbers,
					Name: name,
					Time: time,
				})
			}

			chDatabase <- ret
		default:
			log.Fatalf("Unexcepted type(%T) through channel.", thisCase)
			// chDatabase <- fmt.Errorf("MODULE DB PANIC: A UNEXCEPTED TYPE FOUND")
		}
	}
}
