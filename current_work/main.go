package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	type goodsStr struct {
		Name int
		Type string
	}
	var goods1 []goodsStr
	var db *sql.DB
	db, err := sql.Open("mysql", "root:pass@tcp(62.152.59.9:3306)/shop")
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	fmt.Println(err)
	z := 1

	getFromGoods, err := db.Query("select id, Name from users where id>?;", z)
	for getFromGoods.Next() {
		var pn int
		var pt string
		err := getFromGoods.Scan(&pn, &pt)

		if err != nil {
			log.Fatal(err)
		}

		s := goodsStr{pn, pt}
		goods1 = append(goods1, s)
		//fmt.Println(pn)
	}
	fmt.Println(goods1)

}
