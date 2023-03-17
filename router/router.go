package router

import (
	"github.com/gorilla/mux"
	"go-postgres-menu/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	// API Routes
	router.HandleFunc("/api/menu", controllers.CreateMenu).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/menu", controllers.GetAllMenu).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/menu/{id}", controllers.GetMenu).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/menu/{id}", controllers.UpdateMenu).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/menu/{id}", controllers.DeleteMenu).Methods("DELETE", "OPTIONS")
	return router
}
