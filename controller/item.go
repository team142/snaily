package controller

import (
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"github.com/team142/snaily/model"
)

func InsertItem(conn *pgx.Conn, item *model.Item) (err error) {
	_, err = conn.Exec("insert into madast.items values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)",
		item.ID,
		item.Parent,
		item.Title,
		item.Body,
		item.CreateDate,
		item.CreatedBy,
		item.WaitingFor,
		item.OrgID,
		item.WaitingForDone,
		item.WaitingForDoneDate,
		item.CreatedByDone,
		item.CreatedByDoneDate)
	if err != nil {
		logrus.Errorln(err)
	}
	return
}

func GetItem(conn *pgx.Conn, ID string) (item *model.Item, err error) {
	rows, _ := conn.Query("select * from madast.items where id=$1", ID)
	item = &model.Item{}
	err = rows.Scan(
		&item.ID,
		&item.Parent,
		&item.Title,
		&item.Body,
		&item.CreateDate,
		&item.CreatedBy,
		&item.WaitingFor,
		&item.OrgID,
		&item.WaitingForDone,
		&item.WaitingForDoneDate,
		&item.CreatedByDone,
		&item.CreatedByDoneDate)
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
			err = rows.Scan(
				&item.ID,
				&item.Parent,
				&item.Title,
				&item.Body,
				&item.CreateDate,
				&item.CreatedBy,
				&item.WaitingFor,
				&item.OrgID,
				&item.WaitingForDone,
				&item.WaitingForDoneDate,
				&item.CreatedByDone,
				&item.CreatedByDoneDate)
			if err != nil {
				logrus.Errorln(err)
			}
			out <- &item
		}
		close(out)
	}()
	return
}
