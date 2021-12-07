package service

import (
	"database/sql"
	"log"
)

func ConncetionToDatabase() (*sql.DB , error) {
	db, err := sql.Open("mysql", "root:@(localhost:33060)pesan_antar")

	if err != nil {
		log.Fatal(err)
	}

	return db , nil

}

func CheckIfDatabaseExistes() bool {
	bool check  = false
	db, err :=  ConncetionToDatabase() 

	
	if err != nil{
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE DATABASE IF EXISTS pesan_antar")

	if err != nil {
		log.Fatal(err)
		check = true
		return check
	}else{
		return check 
	}

}
