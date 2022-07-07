package router

// import middleware and mux
import (
	"go-rice-api-postgres/middleware"

	"github.com/gorilla/mux"
)

// router is exported and in main.go
func Router() *mux.Router {
	// create a new router
	router := mux.NewRouter()
	router.HandleFunc("/api/rice", middleware.GetAllRice).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/rice/{id}", middleware.GetRice).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/rice", middleware.CreateRice).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/rice/{id}", middleware.UpdateRice).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/rice/{id}", middleware.DeleteRice).Methods("DELETE", "OPTIONS")
	return router
}