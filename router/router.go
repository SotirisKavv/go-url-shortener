package router

import (
  "net/http"

  "github.com/gorilla/mux"
)

func NewServer(urlHandler http.Handler) *mux.Router {
  router := mux.NewRouter()

  router.Handle("/{short}", urlHandler).Methods("GET")
  router.Handle("/list", urlHandler).Methods("GET")
  router.Handle("/shorten", urlHandler).Methods("POST")

  return router
}
