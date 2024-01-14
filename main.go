package main

import (
	"fmt"
	"golang_udemy/to-do-app/app/controllers"
	"golang_udemy/to-do-app/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
