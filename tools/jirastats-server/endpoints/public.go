package endpoints

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewPublic(router *mux.Router) {
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("public"))))
}
