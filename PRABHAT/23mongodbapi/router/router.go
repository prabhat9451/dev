package router

import (
	"github.com/gorilla/mux"
	"github.com/hiteshchoudhary/mongodbapi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MArkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAllMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovies).Methods("DELETE")
	return router
}
