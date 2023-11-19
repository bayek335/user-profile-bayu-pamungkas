package main

import (
	"fmt"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("environment.env")
	if err != nil {
		fmt.Println(err)
		panic("Error while trying to load env")
	}

	database.DBConnection()

}
