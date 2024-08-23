package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ola-maps/internal/app"
	"github.com/ola-maps/internal/resources"
)

func main() {
	app.LoadConfig()
	router := mux.NewRouter()
	// Define the route for obtaining the access token
	router.HandleFunc("/api/token", resources.GetTokenHandler).Methods("GET")
	router.HandleFunc("/autocomplete", resources.PlaceAutoCompleteHandler).Methods("GET")
	router.HandleFunc("/direction", resources.GetDirectionsHandler).Methods("GET")
	router.HandleFunc("/geocoding", resources.GeoCodeHandler).Methods("GET")
	router.HandleFunc("/reverseGeocoding", resources.ReverseGeocodeHandler).Methods("GET")
	router.HandleFunc("/titleJson", resources.GetTileJSONHandler).Methods("GET")
	router.HandleFunc("/pbfFile", resources.GetPbfFileHandler).Methods("GET")
	router.HandleFunc("/routing/distanceMatrix", resources.GetDistanceMatrixHandler).Methods("GET")
	router.HandleFunc("/mapTiles/ArrayData", resources.ArrayOfDataHandler).Methods("GET")
	router.HandleFunc("/mapTiles/styleDetails", resources.GetStyleDetailsHandler).Methods("GET")
	router.HandleFunc("/map/style", resources.GetMapStyleHandler).Methods("GET")

	log.Println("Starting Ola Maps server on :8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
