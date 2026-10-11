package main

import (
	"github.com/ShashankRaoCoding/kumo/structs/server"
	"github.com/ShashankRaoCoding/kumo/structs/database" 
	"log"
)

func main() {
	s, err := server.New()
	if err != nil {
		log.Println(err.Error()) 
	}

	log.Println(s.Listener.Addr()) 
	s.Start() 

	err = testDB()
	if err != nil {
		log.Println(err.Error()) 
	}
	
}

func testDB() error {
	db, err := database.New("database1.db") 
	if err != nil {
		return err 
	}
	defer db.Close() 

	err = db.CreateTable("users", "username", "password") 
	if err != nil {
		return err 
	}

	err = db.Insert("users", database.Row{"username": "Shashank", "password": "1234567890"})
	if err != nil {
		return err 
	}

	rows, err := db.Query("select * from users where username = Shashank") 
	
}








































