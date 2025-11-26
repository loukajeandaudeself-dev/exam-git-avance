package main

import (
	"log"
	"main/server"
	"net/http"
)

func main() {
	log.Println("Serveur Go en démarrage...")

	http.HandleFunc("/book", server.BookHandler)

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
