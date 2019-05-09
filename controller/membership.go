package controller

import (
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"github.com/team142/snaily/model"
)

func InsertMembership(conn *pgx.Conn, item *model.Membership) (err error) {
	_, err = conn.Exec("insert into madast.memberships values($1, $2, $3, $4)", item.ID, item.OrganizationID, item.UserID, item.IsAdmin)
	if err != nil {
		logrus.Errorln(err)
	}
	return
}

func GetMembershipByOrganization(conn *pgx.Conn, ID string) (out chan *model.Membership, err error) {
	rows, err := conn.Query("select * from madast.memberships where organization_id=$1", ID)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	out = make(chan *model.Membership)
	go func() {
		for rows.Next() {
			item := model.Membership{}
			err = rows.Scan(&item.ID, &item.OrganizationID, &item.UserID, &item.IsAdmin)
			if err != nil {
				logrus.Errorln(err)
			}
			out <- &item
		}
		close(out)
	}()
	return
}

func GetMembershipByUser(conn *pgx.Conn, ID string) (out chan *model.Membership, err error) {
	rows, err := conn.Query("select * from madast.memberships where user_id=$1", ID)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	out = make(chan *model.Membership)
	go func() {
		for rows.Next() {
			item := model.Membership{}
			err = rows.Scan(&item.ID, &item.OrganizationID, &item.UserID, &item.IsAdmin)
			if err != nil {
				logrus.Errorln(err)
			}
			out <- &item
		}
		close(out)
	}()
	return
}
