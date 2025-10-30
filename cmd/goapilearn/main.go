package main

import "goApiLearn/config"

func main() {

	appConfig := config.LoadConfig()
	config.InitMongo(appConfig.Mongo.Connection)
	config.InitApiConfig()
}
