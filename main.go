package main

import (
	"log"
	"net/http"
	"tomovale_deleta_duplicado/src/router"
)

func main() {
	r := router.GetRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
