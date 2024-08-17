package main

const DBName = "learning"
const DBUser = "root"
const DBPwd = "0123456"

import (
	"fmt"
	- "github.com/go-sql-driver/mysql"
)

func checkErreur(e error){
	if e != nil{
		log.fatalln(e)
	}

}

type Data struct {
	id int
	name string
}

func main()  {
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v",DBUser,DBPwd, DBName)
	db , err := sql.open("mysql",connectionString)
	checkErreur(err)
	defer db.close()
	result, err := db.Exec("insert into data values (3,"ali")")
	checkErreur(err)
	lastInsertedId, err := result.lastInsertedId()
	fmt.Println("lastInsertedId: ", lastInsertedId)
	rowsAffected, err := result.RowsAffected()
	fmt.Println("rowsAffected: ", rowsAffected)
	checkErreur(err)

	rows , err := db.Query("select * from data")
	checkErreur(err)

	for rows.Next(){
		var data Data
		err := rows.Scan(&data.id , &data.name)
		checkErreur(err)
		fmt.Println(data)
	}

}