package main

import (
	_ "github.com/wishdream/pisces/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/wishdream/pisces/routers"
)

func init() {
	orm.RegisterDriver("sqlite", orm.DR_Sqlite)
	orm.RegisterDataBase("default", "sqlite3", "pisces.db")

	name := "default"
	force := false
	verbose := false

	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		beego.Error("orm runsyncdb err:", err.Error())
	}
}

func main() {
	beego.Run()
}
