package controller

import (
	"github.com/jackc/pgx"
	"github.com/just1689/hotpotato/model"
	"github.com/sirupsen/logrus"
)

func InsertItem(conn *pgx.Conn, item *model.Item) (err error) {
	_, err = conn.Exec("insert into madast.items values($1, $2, $3, $4, $5, $6, $7)", item.ID, item.Parent, item.Title, item.Body, item.CreateDate, item.Thrower, item.Holder)
	if err != nil {
		logrus.Errorln(err)
	}
	return
}

func GetItem(conn *pgx.Conn, ID string) (item *model.Item, err error) {
	rows, _ := conn.Query("select * from madast.items where id=$1", ID)
	item = &model.Item{}
	err = rows.Scan(&item.ID, &item.Parent, &item.Title, &item.Body, &item.CreateDate, &item.Thrower, &item.Holder)
	if err != nil {
		logrus.Errorln(err)
	}
	return
}

func GetChildrenByParent(conn *pgx.Conn, ID string) (out chan *model.Item, err error) {
	rows, err := conn.Query("select * from madast.items where parent_id=$1", ID)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	out = make(chan *model.Item)
	go func() {
		for rows.Next() {
			item := model.Item{}
			err = rows.Scan(&item.ID, &item.Parent, &item.Title, &item.Body, &item.CreateDate, &item.Thrower, &item.Holder)
			if err != nil {
				logrus.Errorln(err)
			}
			out <- &item
		}
		close(out)
	}()
	return
}
