package main

import "github.com/gorilla/mux"

//-----Setup route koneksi ke controller-----//
func setRoutes(r *mux.Router) {
	r.HandleFunc("/pesan", olahPesan).Methods("POST")
}
