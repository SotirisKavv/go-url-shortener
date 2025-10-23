package main

import (
	"fmt"
	"net/http"
	"url-shortener/handler"
	"url-shortener/router"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	urlHandler := handler.NewURLHandler()
	server := router.NewServer(urlHandler)

	fmt.Println("Serving and listening on :8080")
	http.ListenAndServe(":8080", server)
}
