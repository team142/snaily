package controller

import (
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"github.com/team142/snaily/model"
)

func InsertUser(conn *pgx.Conn, item *model.User) (err error) {
	_, err = conn.Exec("insert into madast.users values($1, $2, $3, $4)", item.ID, item.Email, item.FirstName, item.LastName)
	if err != nil {
		logrus.Errorln(err)
	}
	return
}

func GetUser(conn *pgx.Conn, ID string) (item *model.User, err error) {
	rows, _ := conn.Query("select * from madast.users where id=$1", ID)
	item = &model.User{}
	err = rows.Scan(&item.ID, &item.Email, &item.FirstName, &item.LastName)
	if err != nil {
		logrus.Errorln(err)
	}
	return
}

func UserExists(conn *pgx.Conn, email string) bool {
	item, _ := GetUserByEmail(conn, email)
	if item == nil || item.ID == "" {
		return false
	}
	return true
}

func GetUserByEmail(conn *pgx.Conn, email string) (item *model.User, err error) {
	rows, _ := conn.Query("select * from madast.users where email=$1", email)
	item = &model.User{}
	err = rows.Scan(&item.ID, &item.Email, &item.FirstName, &item.LastName)
	if err != nil {
		logrus.Errorln(err)
	}
	return
}
