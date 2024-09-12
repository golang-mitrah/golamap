package app

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ClientID     string
	ClientSecret string
	TokenURL     string
	// PlacesURL                string
	DirectionsURL            string
	PlaceAutoCompleteURL     string
	GeoCodeURL               string
	ReverseGeocodeURL        string
	PbfFileURL               string
	DistanceMatrixURL        string
	ArrayOfDataURL           string
	StyleDetailsURL          string
	MapStyleURL              string
	PlaceDetailURL           string
	NearBySearchURL          string
	TextSearchURL            string
	SnapToRoadURL            string
	NearestRoadsURL          string
	StaticMapImageCenterURL  string
	StaticMapImageBoundedURL string
	StaticMapImageURL        string
)

type Configuration interface {
	LoadConfig()
}

func Config(config Configuration) Configuration {
	return config
}

type EnvConfig struct{}

func (congig *EnvConfig) LoadConfig() {
	log.Println("Loading environment variables from .env file")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file = %v ", err)
	}

	ClientID = os.Getenv("CLIENT_ID")
	ClientSecret = os.Getenv("CLIENT_SECRET")
	TokenURL = os.Getenv("TOKEN_URL")
	// PlacesURL = os.Getenv("PLACES_URL")
	DirectionsURL = os.Getenv("DIRECTIONS_URL")
	PlaceAutoCompleteURL = os.Getenv("PLACE_AUTO_COMPLETE_URL")
	GeoCodeURL = os.Getenv("GEOCODE_URL")
	ReverseGeocodeURL = os.Getenv("REVERSE_GEOCODE_URL")
	PbfFileURL = os.Getenv("PBFFILE_URL")
	DistanceMatrixURL = os.Getenv("DISTANCE_MATRIX_URL")
	StyleDetailsURL = os.Getenv("STYLE_DETAILS_URL")
	ArrayOfDataURL = os.Getenv("ARRAY_OF_DATA_URL")
	MapStyleURL = os.Getenv("MAP_STYLE_URL")
	PlaceDetailURL = os.Getenv("PLACE_DETAIL_URL")
	NearBySearchURL = os.Getenv("NEARBY_SEARCH_URL")
	TextSearchURL = os.Getenv("TEXT_SEARCH_URL")
	SnapToRoadURL = os.Getenv("SNAP_TO_ROAD_URL")
	NearestRoadsURL = os.Getenv("NEAREST_ROADS_URL")
	StaticMapImageCenterURL = os.Getenv("STATIC_MAP_IMAGE_CENTER_URL")
	StaticMapImageBoundedURL = os.Getenv("STATIC_MAP_IMAGE_BOUNDED_URL")
	StaticMapImageURL = os.Getenv("STATIC_MAP_IMAGE_URL")

}
