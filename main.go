package main

import (
	"fmt"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/database"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/router"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("environment.env")
	if err != nil {
		fmt.Println(err)
		panic("Error while trying to load env")
	}

}
func main() {
	// daabse connection
	db := database.DBConnection()

	// router
	router.Routes(db)

}
