package main

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"model"

	_ "github.com/go-sql-driver/mysql"
	"utils/log"
)

var engine *xorm.Engine

func init(){
	var err error
	engine, err = xorm.NewEngine("mysql","root:rock1204@tcp(localhost:3306)/compute?charset=utf8")
	if err != nil {
		log.Fatalf("Fail to create engne: %v", err)
	}
	if err = engine.Sync(new(model.Account)); err != nil {
		log.Fatalf("fail to sync database: %v", err)
	}
}

func newAccount(name string, balance float64) error {
	_, err := engine.Insert(&model.Account{Name:name, Balance:balance})
	return err
}

func getAccount(id int64) (*model.Account, error) {
	a := &model.Account{}
	has, err := engine.Id(id).Get(a)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("Account not found")
	}

	return  a, nil
}

func makeDeposit(id int64, depostit float64) (*model.Account, error) {
	a, err := getAccount(id)
	if err != nil {
		return  nil, err
	}
	a.Balance += depostit
	_, err = getAccount(id)
	if err != nil {
		return nil, err
	}
	a.Balance += depostit
	_, err = engine.Update(a)
	return a, err
}

func makeWithdraw(id int64, withdraw float64)(*model.Account, error){
	a, err := getAccount(id)
	fmt.Printf("%#v\n", a)
	if err != nil {
		return nil, err
	}
	if a.Balance <= withdraw {
		return nil, errors.New("Not enough balance")
	}
	a.Balance -= withdraw
	_,err = engine.Update(a)
	return a, err
}

func main() {
	//fmt.Println("Please enter <name> <balance>:")
	//var name string
	//var balance float64
	//fmt.Scanf("%s %f\n", &name, &balance)
	//if err := newAccount(name, balance); err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("please enter <id>")
	//var id int64
	///Sscanf()canf("%d\n", &id)
	//a, err := getAccount(id)
	//if err != nil{
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%#v\n", a)
	//}
	fmt.Println("please enter <id> <ba>")
	_, err := makeWithdraw(4,4)
	if err != nil {
		log.Error(err)
	}




}
