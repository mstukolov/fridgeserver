package psql

import (
	"time"
	"fmt"
	"mstukolov/fridgeserver/database/connect"
)

type Retailstore struct {
	Id        int
	Name      string
	RetailchainId int
	Retailchain *Retailchain
	Createdat time.Time
	Updatedat time.Time
}

var model Retailstore

func GetAll_Retailstores() []Retailstore {
	var all []Retailstore
	err := connect.GetDB().Model(&all).Select()
	if err != nil {
		panic(err)
	}
	return all
}
func Get_ChainRetailStores(id int) []Retailstore{
	var all []Retailstore
	err := connect.GetDB().Model(&all).Where("retailchainid = ?", id).Select()
	if err != nil {
		panic(err)
	}
	return all
}

func Get_RetailStore(id int) Retailstore{
	err := connect.GetDB().Model(&model).Where("id = ?", id).Select()
	if err != nil {
		panic(err)
	}
	return model
}

func Create_RetailStore(model Retailstore) Retailstore{
	err := connect.GetDB().Insert(&model)
	if err != nil {
		panic(err)
	}
	return model
}
func Update_RetailStore(model Retailstore) Retailstore{
	err := connect.GetDB().Update(&model)
	if err != nil {
		panic(err)
	}
	return model
}

func Delete_RetailStore(model Retailstore) Retailstore{
	_, err := connect.GetDB().Model(&model).Delete(&model)
	if err != nil {
		panic(err)
	}
	return model
}

func DeleteAll_RetailStore(){
	res, err := connect.GetDB().Exec("DELETE FROM retailstores")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deleted success: %s", res)
}
