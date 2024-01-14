package main

import (
	"fmt"
	"golang_udemy/to-do-app/app/controllers"
	"golang_udemy/to-do-app/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

	/*
		sessionチェック実装テスト用
	*/
	// user, _ := models.GetUserByEmail("test@example.com")
	// fmt.Println(user)

	// session, err := user.CreateSession()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(session)

	// valid, _ := session.CheckSession()
	// fmt.Println(valid)
}
