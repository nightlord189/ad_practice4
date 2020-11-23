package main

import (
	"fmt"
	"github.com/nightLord189/ad_practice4/config"
	"github.com/nightLord189/ad_practice4/db"
	"github.com/nightLord189/ad_practice4/model"
)

func main() {
	fmt.Println("start")
	conf := config.Load("config.json")
	session := db.Launch(conf)
	defer session.Close()

	//CREATE
	fmt.Println("create")
	item := model.Almobglu{
		ID:       145,
		Dolobimp: "5.2%",
		Importk:  "Ресей",
		Importr:  "Россия",
		Namer:    "Станки металлообрабатывающие",
	}
	db.Create(session, &item)

	//READ
	fmt.Println("read")
	itemRead, err := db.Read(session, 145)
	if err != nil {
		fmt.Println("read err: " + err.Error())
	} else {
		fmt.Printf("read result: %v\n", itemRead)
	}

	//UPDATE
	fmt.Println("update")
	err = db.Update(session, 145, "Корма животноводческие")
	if err != nil {
		fmt.Println("update err: " + err.Error())
	}
	itemRead, _ = db.Read(session, 145)
	fmt.Printf("update result: %v\n", itemRead)

	//DELETE
	fmt.Println("delete")
	err = db.Delete(session, 145)
	if err != nil {
		fmt.Println("delete err: " + err.Error())
	}
}
