package golamap

import (
	"errors"
	"strings"
)

var (
	MockDirectionsURL     = "routing/v1/directions"
	MockDistanceMatrixURL = "routing/v1/distanceMatrix"
	MockAutoCompleteURL   = "places/v1/autocomplete"
	MockGecodeURL         = "places/v1/geocode"
	MockReverseGecodeURL  = "places/v1/reverse-geocode"
	MockArrayOfDataURL    = "tiles/vector/v1/data"
	MockPlaceDetailURL    = "places/v1/details"
	MockNearBySearchURL   = "places/v1/nearbysearch"
	MockTextSearchURL     = "places/v1/textsearch"
	MockSnapToRoadURL     = "routing/v1/snapToRoad"
	MockNearestRoadURL    = "routing/v1/nearestRoads"
	MockStaticMapImageURL = "static/auto"
	MockMapStyleURL       = "tiles/vector/v1/styles.json"
	MockStyleDetailsURL   = "style.json"
)

type MockInterface interface {
	SendOlaMapRequest(method, url, requestID, oauthToken string, responseObj interface{}) error
}
type MockStruct struct {
	MockBody   string
	StatusCode int
}

func (mock *MockStruct) SendOlaMapRequest(method, url, requestID, oauthToken string, responseObj interface{}) error {
	switch {
	case strings.Contains(url, MockDirectionsURL):
		mock.StatusCode = 200
		mock.MockBody = DirectionResponse
	case strings.Contains(url, MockDistanceMatrixURL):
		mock.StatusCode = 200
		mock.MockBody = DistanceMatrixResponse
	case strings.Contains(url, MockGecodeURL):
		mock.StatusCode = 200
		mock.MockBody = GeoCodeResponse
	case strings.Contains(url, MockReverseGecodeURL):
		mock.StatusCode = 200
		mock.MockBody = ReverseGeocodeResponse
	case strings.Contains(url, MockSnapToRoadURL):
		mock.StatusCode = 200
		mock.MockBody = SnapToRoadResponse
	case strings.Contains(url, MockNearestRoadURL):
		mock.StatusCode = 200
		mock.MockBody = NearestRoadResponse
	case strings.Contains(url, MockArrayOfDataURL):
		mock.StatusCode = 200
		mock.MockBody = ArrayOfDataResponse
	case strings.Contains(url, MockAutoCompleteURL):
		mock.StatusCode = 200
		mock.MockBody = AutoCompleteResponse
	case strings.Contains(url, MockPlaceDetailURL):
		mock.StatusCode = 200
		mock.MockBody = PlaceDetailResponse
	case strings.Contains(url, MockNearBySearchURL):
		mock.StatusCode = 200
		mock.MockBody = NearBySearchResponses
	case strings.Contains(url, MockTextSearchURL):
		mock.StatusCode = 200
		mock.MockBody = TextSearchResponse
	case strings.Contains(url, MockMapStyleURL):
		mock.StatusCode = 200
		mock.MockBody = MapStyleResponse
	case strings.Contains(url, MockStyleDetailsURL):
		mock.StatusCode = 200
		mock.MockBody = StyleDetailResponse
	default:
		mock.StatusCode = 400
		mock.MockBody = nil
		return errors.New("Invalid request")
	}
	return nil
}
