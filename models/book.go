package models

import (
	"github.com/astaxie/beego/orm"
)

type Book struct {
	Id              int
	Name            string
	Author          string
	Member_id       int
	Type            string
	Path            string
	Cover           string
	CreatedTime     int
	LastUpdatedTime int
}

func (b *Book) Read(fields ...string) error {
	if err := orm.NewOrm().Read(b, fields...); err != nil {
		return err
	}
	return nil
}

func (b *Book) Insert() error {
	if _, err := orm.NewOrm().Insert(b); err != nil {
		return err
	}
	return nil
}

