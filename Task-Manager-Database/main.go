package main

import (
	"log"
	"task-manager-database/config"
	"task-manager-database/internal"
)

func main() {
	err := config.ConfigureLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer config.LogFile.Close()

	log.Println("LOGGER CONFIGURED")

	err = config.ConfigureDatabase()
	if err != nil {
		log.Fatal(err)
	}

	//err = config.ConfigureRedis()
	//if err != nil {
	//	log.Fatal(err)
	//}

	internal.CreateServer()
	err = internal.RunApp()
	if err != nil {
		log.Fatal(err)
	}
}
