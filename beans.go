package dbpg

import (
	_ "github.com/lib/pq"
)

type UserBean struct {
	Username string `sql:"username"`
	Pwd string `sql:"pwd"`
	NickName string `sql:"nickname"`
	RealName string  `sql:"realname"`
}

type RoleBean struct {

}



func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}