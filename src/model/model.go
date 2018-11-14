package model

type Account struct {
	Id int64
	Name string `xorm:"unique"`
	Balance float64
	Version int `xorm:"version"`
}

