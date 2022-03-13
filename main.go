package main

import (
	"golang-twitter-clone/api"
	"golang-twitter-clone/database"
	"golang-twitter-clone/migrations"
)

func main() {
	//router.InitializeRouter()
	//migrations.Migrate()
	//migrations.Migrate()
	database.InitDatabase()
	migrations.Migrate()
	api.InitializeRouter()

}












