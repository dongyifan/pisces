package models

import (
	"github.com/astaxie/beego/orm"
)

// User :this is user models
type User struct {
	Id         int64
	CounterNum string `orm:"unique"`
	Name       string
	Property   float64
	Debt       float64
}

func init() {
	orm.RegisterModel(new(User))
}
