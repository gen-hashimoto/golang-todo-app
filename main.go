package main

import (
	"fmt"
	"golang_udemy/to-do-app/app/models"
	"log"
)

func main() {
	fmt.Println(models.Db)

	// controllers.StartMainServer()
	user, _ := models.GetUserByEmail("test@example.com")
	fmt.Println(user)

	session, err := user.CreateSession()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(session)
}
