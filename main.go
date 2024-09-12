package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ola-maps/internal/app"
	"github.com/ola-maps/internal/resources"
)

func main() {
	envLoading := app.Config(&app.EnvConfig{})
	envLoading.LoadConfig()
	//app.LoadConfig()
	router := mux.NewRouter()

	api := router.PathPrefix("/api/v1").Subrouter()

	// Obtain the access token
	api.HandleFunc("/token", resources.GetTokenHandler).Methods("GET")

	//Routing API
	routingRouter := api.PathPrefix("/routing").Subrouter()

	routingRouter.HandleFunc("/directions", resources.GetDirectionsHandler).Methods("GET")
	routingRouter.HandleFunc("/distanceMatrix", resources.GetDistanceMatrixHandler).Methods("GET")

	//Roads API
	routingRouter.HandleFunc("/snapToRoad", resources.GetSnapToRoadHandler).Methods("GET")
	routingRouter.HandleFunc("/nearestRoads", resources.GetSnapToRoadHandler).Methods("GET")

	//Geocode API
	placesRouter := api.PathPrefix("/places").Subrouter()
	placesRouter.HandleFunc("/geocoding", resources.GeoCodeHandler).Methods("GET")
	placesRouter.HandleFunc("/reverseGeocoding", resources.ReverseGeocodeHandler).Methods("GET")

	//Places API
	placesRouter.HandleFunc("/autocomplete", resources.PlaceAutoCompleteHandler).Methods("GET")
	placesRouter.HandleFunc("/details", resources.GetPlaceDetailHandler).Methods("GET")
	placesRouter.HandleFunc("/nearbysearch", resources.GetNearBySearchHandler).Methods("GET")
	placesRouter.HandleFunc("/textsearch", resources.GetTextSearchHandler).Methods("GET")

	//Maptiles API
	tilesRouter := api.PathPrefix("/tiles").Subrouter()
	tilesRouter.HandleFunc("/data", resources.ArrayOfDataHandler).Methods("GET")
	tilesRouter.HandleFunc("/styles", resources.GetMapStyleHandler).Methods("GET")
	tilesRouter.HandleFunc("/stylesByName", resources.GetStyleDetailsHandler).Methods("GET")
	tilesRouter.HandleFunc("/pbfFile", resources.GetPbfFileHandler).Methods("GET")

	//Static Tiles API
	tilesRouter.HandleFunc("/styleMapImageCenterpoint", resources.GetStaticMapImageCenterHandler).Methods("GET")
	tilesRouter.HandleFunc("/styleMapImageBoundingbox", resources.GetStaticMapImageBoundedHandler).Methods("GET")
	tilesRouter.HandleFunc("/styleMapImage", resources.StaticMapImageHandler).Methods("GET")

	log.Println("Starting Ola Maps server on :8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
