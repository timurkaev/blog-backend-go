package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", homeHandler).Methods("GET")

	return router
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the GO API"))
}
