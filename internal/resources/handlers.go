package resources

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/google/uuid"
	"github.com/ola-maps/internal"
	"github.com/ola-maps/internal/app"
)

func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	// get access token
	token, err := GetAccessToken(app.ClientID, app.ClientSecret, app.TokenURL)
	if err != nil {
		log.Printf("Error getting access token: %v", err)
		http.Error(w, "Unable to get access token", http.StatusInternalServerError)
		return
	}

	// Return the token as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"access_token": token})
}

func GetDirectionsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	origin := query.Get("origin")
	destination := query.Get("destination")

	if origin == "" || destination == "" {
		http.Error(w, "Missing required query parameters: 'origin' and/or 'destination'", http.StatusBadRequest)
		return
	}

	oauthToken := r.Header.Get("Authorization")
	if oauthToken == "" {
		http.Error(w, "Missing OAuth token", http.StatusUnauthorized)
		return
	}

	// Construct the URL for the Olamaps API request
	url := fmt.Sprintf("https://api.olamaps.io/routing/v1/directions?origin=%s&destination=%s",
		origin, destination)

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}
	requestID := uuid.New().String()
	// Make the external request
	err := internal.MakeExternalRequest("POST", url, requestID, oauthToken, &apiResponse)
	if err != nil {
		http.Error(w, "Failed to send request to Olamaps API", http.StatusInternalServerError)
		return
	}

	// Respond with the API response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(apiResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func PlaceAutoCompleteHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	input := query.Get("input")

	if input == "" {
		http.Error(w, "Missing 'input' query parameter", http.StatusBadRequest)
		return
	}

	// Retrieve the OAuth token and request ID from the headers
	oauthToken := r.Header.Get("Authorization")
	if oauthToken == "" {
		http.Error(w, "Missing OAuth token", http.StatusUnauthorized)
		return
	}

	requestID := uuid.New().String()

	// Construct the URL for the Olamaps API request
	url := fmt.Sprintf("https://api.olamaps.io/places/v1/autocomplete?input=%s", input)

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}

	// Make the external request
	err := internal.MakeExternalRequest("GET", url, requestID, oauthToken, &apiResponse)
	if err != nil {
		http.Error(w, "Failed to send request to Olamaps API", http.StatusInternalServerError)
		return
	}

	// Respond with the API response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(apiResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func GeoCodeHandler(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters from the URL
	query := r.URL.Query()
	address := query.Get("address")
	bounds := query.Get("bounds")
	language := query.Get("language")

	if address == "" || bounds == "" || language == "" {
		http.Error(w, "Missing required query parameters", http.StatusBadRequest)
		return
	}

	// Retrieve the OAuth token and request ID from the headers
	oauthToken := r.Header.Get("Authorization")
	if oauthToken == "" {
		http.Error(w, "Missing OAuth token", http.StatusUnauthorized)
		return
	}

	// Retrieve the X-Request-Id from the headers
	requestID := uuid.New().String()

	// Construct the URL for the Olamaps API request
	url := fmt.Sprintf(
		"https://api.olamaps.io/places/v1/geocode?address=%s&bounds=%s&language=%s",
		address,
		bounds,
		language,
	)
	fmt.Println("url----------->", url)

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}

	// Make the external request
	err := internal.MakeExternalRequest("GET", url, requestID, oauthToken, &apiResponse) // Empty token if not needed
	if err != nil {
		http.Error(w, "Failed to send request to Olamaps API", http.StatusInternalServerError)
		return
	}

	// Respond with the API response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(apiResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func ReverseGeocodeHandler(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters from the URL
	query := r.URL.Query()
	latlng := query.Get("latlng")

	if latlng == "" {
		http.Error(w, "Missing required query parameter: latlng", http.StatusBadRequest)
		return
	}

	// Retrieve the OAuth token and request ID from the headers
	oauthToken := r.Header.Get("Authorization")
	if oauthToken == "" {
		http.Error(w, "Missing OAuth token", http.StatusUnauthorized)
		return
	}

	// Retrieve the X-Request-Id from the headers
	requestID := uuid.New().String()

	// Construct the URL for the API request
	urlWithParams := fmt.Sprintf("https://api.olamaps.io/places/v1/reverse-geocode?latlng=%s", url.QueryEscape(latlng))

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}

	// Make the external request
	err := internal.MakeExternalRequest("GET", urlWithParams, requestID, oauthToken, &apiResponse)
	if err != nil {
		http.Error(w, "Failed to send request to API", http.StatusInternalServerError)
		return
	}

	// Respond with the API response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(apiResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func GetTileJSONHandler(w http.ResponseWriter, r *http.Request) {
	// Extract dataset name from URL path
	datasetName := r.URL.Query().Get("dataset_name")
	if datasetName == "" {
		http.Error(w, "Missing required query parameter: dataset_name", http.StatusBadRequest)
		return
	}

	// Retrieve the OAuth token and request ID from the headers
	oauthToken := r.Header.Get("Authorization")
	if oauthToken == "" {
		http.Error(w, "Missing OAuth token", http.StatusUnauthorized)
		return
	}

	// Retrieve the X-Request-Id from the headers
	requestID := uuid.New().String()

	// Construct the URL for the API request
	urlWithParams := fmt.Sprintf("https://api.olamaps.io/tiles/vector/v1/data/%s.json", url.PathEscape(datasetName))

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}

	// Make the external request
	err := internal.MakeExternalRequest("GET", urlWithParams, requestID, oauthToken, &apiResponse)
	if err != nil {
		http.Error(w, "Failed to send request to API", http.StatusInternalServerError)
		return
	}

	// Respond with the API response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(apiResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func GetPbfFileHandler(w http.ResponseWriter, r *http.Request) {
	// Extract dataset name from URL path
	datasetName := r.URL.Query().Get("dataset_name")
	z := r.URL.Query().Get("z")
	x := r.URL.Query().Get("x")
	y := r.URL.Query().Get("y")

	if datasetName == "" || z == "" || x == "" || y == "" {
		http.Error(w, "Missing required query parameters: dataset_name, z, x, y", http.StatusBadRequest)
		return
	}

	// Retrieve the OAuth token and request ID from the headers
	oauthToken := r.Header.Get("Authorization")
	if oauthToken == "" {
		http.Error(w, "Missing OAuth token", http.StatusUnauthorized)
		return
	}

	// Retrieve the X-Request-Id from the headers
	requestID := uuid.New().String()

	// Construct the URL for the API request
	urlWithParams := fmt.Sprintf("https://api.olamaps.io/tiles/vector/v1/data/%s/%s/%s/%s.pbf", url.PathEscape(datasetName), url.PathEscape(z), url.PathEscape(x), url.PathEscape(y))

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}

	// Make the external request
	err := internal.MakeExternalRequest("GET", urlWithParams, requestID, oauthToken, &apiResponse)
	if err != nil {
		http.Error(w, "Failed to send request to API", http.StatusInternalServerError)
		return
	}

	// Respond with the API response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(apiResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func GetDistanceMatrixHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	origins := query.Get("origins")
	destinations := query.Get("destinations")

	if origins == "" || destinations == "" {
		http.Error(w, "Missing required query parameters: 'origins' and/or 'destinations'", http.StatusBadRequest)
		return
	}

	// Extract optional headers
	// xRequestID := r.Header.Get("x_request_id")
	// xCorrelationID := r.Header.Get("x_correlation_id")
	oauthToken := r.Header.Get("Authorization")

	// Construct the URL for the Olamaps API request
	url := fmt.Sprintf("https://api.olamaps.io/routing/v1/distanceMatrix?origins=%s&destinations=%s",
		url.QueryEscape(origins), url.QueryEscape(destinations))

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}
	requestID := uuid.New().String()

	// Make the external request
	err := internal.MakeExternalRequest("GET", url, requestID, oauthToken, &apiResponse)
	if err != nil {
		http.Error(w, "Failed to send request to Olamaps API", http.StatusInternalServerError)
		return
	}

	// Respond with the API response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(apiResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func ArrayOfDataHandler(w http.ResponseWriter, r *http.Request) {
	datasetName := r.URL.Query().Get("dataset_name")
	if datasetName == "" {
		http.Error(w, "Missing required query parameter: 'dataset_name'", http.StatusBadRequest)
		return
	}

	// Extract optional headers
	oauthToken := r.Header.Get("Authorization")

	// Construct the URL for the Olamaps API request
	apiURL := fmt.Sprintf("https://api.olamaps.io/tiles/vector/v1/data/%s.json", url.QueryEscape(datasetName))

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}
	requestID := uuid.New().String()

	// Make the external request
	err := internal.MakeExternalRequest("GET", apiURL, requestID, oauthToken, &apiResponse)
	if err != nil {
		http.Error(w, "Failed to send request to Olamaps API", http.StatusInternalServerError)
		return
	}

	// Respond with the API response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(apiResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func GetStyleDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the style_name query parameter
	styleName := r.URL.Query().Get("style_name")
	if styleName == "" {
		http.Error(w, "Missing required query parameter: 'style_name'", http.StatusBadRequest)
		return
	}

	// Extract optional headers
	oauthToken := r.Header.Get("Authorization")

	// Construct the URL for the Olamaps API request
	apiURL := fmt.Sprintf("https://api.olamaps.io/tiles/vector/v1/styles/%s/style.json", url.QueryEscape(styleName))

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}
	requestID := uuid.New().String()

	fmt.Println()
	// Make the external request
	err := internal.MakeExternalRequest("GET", apiURL, requestID, oauthToken, &apiResponse)
	if err != nil {
		http.Error(w, "Failed to send request to Olamaps API", http.StatusInternalServerError)
		return
	}

	// Respond with the API response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(apiResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
