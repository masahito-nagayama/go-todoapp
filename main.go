package main 

import (
	"fmt"
	// "log" 
	"go_todoapp/app/models"
	"go_todoapp/app/controllers"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}