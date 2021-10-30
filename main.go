package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mkholjuraev/aha_engine/db/admin"
	"github.com/mkholjuraev/aha_engine/manager"
	"github.com/mkholjuraev/aha_engine/manager/auth"
)

func main() {
	conn := admin.NewDatabaseConncetion()

	fmt.Println(conn)

	http.HandleFunc("/login", auth.Login)
	http.HandleFunc("/profile", manager.Profile)

	fmt.Printf("Listening on port: 8082\n")
	log.Fatal((http.ListenAndServe(":8082", nil)))
}
