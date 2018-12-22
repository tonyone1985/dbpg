package dbpg

import (
	"context"
	"database/sql"
	"log"

	"../sqlmapper"
)
import _ "github.com/lib/pq"

var db *sql.DB

func init() {

	var err error
	db, err = sql.Open("postgres", "postgres://postgres:123456@localhost/dbtest?sslmode=disable")
	sqlmapper.SetDriver(sqlmapper.POSTGRES)
	sqlstr := `SELECT * from "User"`

	ctx := context.Background()

	stmt, err := db.PrepareContext(ctx, sqlstr)
	//dd.
	if err != nil {
		log.Println("init db error")
		return
	}

	r := stmt.QueryRowContext(ctx)
	if r == nil {
		return
	}

	if err != nil {
		log.Println("init db error")
		return
	}

}

func Update(u *UserBean) bool {

	fm, err := sqlmapper.NewFieldsMap("User", u)
	ctx := context.Background()
	if err != nil {
		return false
	}
	err = fm.SQLUpdateByPriKey(ctx, nil, db)
	if err != nil {
		return false
	}
	return true
}

func Queryusr() *UserBean {
	//r,e:=db.Exec("select * from \"user\" where username=?","cxx")

	//sqlmapper.SetSqlTag("\"", "\"")

	var row UserBean
	row.Username = "cxx"
	fm, err := sqlmapper.NewFieldsMap("User", &row)
	if err != nil {
		return nil
	}
	ctx := context.Background()
	ub, err := fm.SQLSelectByPriKey(ctx, nil, db)
	if err != nil {
		return nil
	}
	return ub.(*UserBean)

}
