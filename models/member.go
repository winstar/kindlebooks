package models

import "github.com/astaxie/beego/orm"

type Member struct {
	Id              int
	Name            string    `orm:"size(30)"`
	Email           string
	Mobile          string
	Password        string
	CreatedTime     int
	lastUpdatedTime int
}

func (m *Member) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Member) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}