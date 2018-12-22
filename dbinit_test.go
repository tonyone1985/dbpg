package dbpg

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"testing"
)

func Test_testf(t *testing.T) {
	for i := 0; i < 7000; i++ {
		u := Queryusr()
		u.NickName = "nname" + strconv.FormatInt(int64(i), 10)
		u.Pwd = "pwd" + strconv.FormatInt(int64(i), 10)
		u.RealName = "realname" + strconv.FormatInt(int64(i), 10)
		Update(u)
		log.Println(u.NickName)
	}

}

func testf() {
	//db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=123456 dbname=dbtest sslmode=disable")
	db, err := sql.Open("postgres", "postgres://postgres:123456@localhost/dbtest?sslmode=disable")
	log.Println("oepn")
	checkErr(err)
	log.Println("oepnend")

	stmt, err := db.Prepare("update role set Remark=$1 where RoleName=$2;")
	checkErr(err)

	log.Println("updateend")

	res, err := stmt.Exec("321", "管理员")
	checkErr(err)
	log.Println("ficow")

	affect, err := res.RowsAffected()
	checkErr(err)

	log.Println("RowsAffected")
	fmt.Println("rows affect:", affect)

	log.Println("rows affect")
}
