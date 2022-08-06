package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Age      uint16 `json:"age"`
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Создание пользователя
	//insert, err := db.Query("INSERT INTO `users` (`name`, `lastname`, `age`) VALUES('Dmitry', 'Nekto', 18)")
	//if err != nil {
	//	panic(err)
	//}
	//defer insert.Close()

	// Обновляет все имена таблицы
	//update, err := db.Query("UPDATE `users` SET name = 'Aleksey'")
	//if err != nil {
	//	panic(err)
	//}
	//defer update.Close()

	// Выводит все данные из MySQL
	res, err := db.Query("SELECT `name`, `lastname`, `age` FROM `users`")
	if err != nil {
		panic(err)
	}

	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.LastName, &user.Age)
		if err != nil {
			panic(err)
		}

		fmt.Println(fmt.Sprintf("User: %s %s with age %d", user.Name, user.LastName, user.Age))
	}

	//defer res.Close()

}
