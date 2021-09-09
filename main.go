package main 

import (
	"fmt"
	"go_todoapp/app/controllers"
	"go_todoapp/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}