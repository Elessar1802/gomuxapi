package main

import (
	"github.com/Elessar1802/api/src/v1"
	"log"
	"net/http"
)

const PORT = ":8000"

func main() {
	r := router.InitRouter()
	log.Fatal(http.ListenAndServe(PORT, r))
}
