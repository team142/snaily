package controller

import (
	"github.com/jackc/pgx"
	"github.com/just1689/hotpotato/model"
	"github.com/sirupsen/logrus"
)

func InsertOrganization(conn *pgx.Conn, item *model.Organization) (err error) {
	_, err = conn.Exec("insert into madast.orgs values($1, $2)", item.ID, item.Name)
	if err != nil {
		logrus.Errorln(err)
	}
	return
}

func GetOrganization(conn *pgx.Conn, ID string) (item *model.Organization, err error) {
	rows, _ := conn.Query("select * from madast.orgs where id=$1", ID)
	item = &model.Organization{}
	err = rows.Scan(&item.ID, &item.Name)
	if err != nil {
		logrus.Errorln(err)
	}
	return
}
