package resources

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

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
	url := fmt.Sprintf(app.DirectionsURL, origin, destination)

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}
	requestID := uuid.New().String()
	// Make external request
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
	url := fmt.Sprintf(app.PlaceAutoCompleteURL, input)

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

	if address == "" && bounds == "" && language == "" {
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
		app.GeoCodeURL,
		address,
		bounds,
		language,
	)

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
	urlWithParams := fmt.Sprintf(app.ReverseGeocodeURL, url.QueryEscape(latlng))

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

	if datasetName == "" && z == "" && x == "" && y == "" {
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
	urlWithParams := fmt.Sprintf(app.PbfFileURL, url.PathEscape(datasetName), url.PathEscape(z), url.PathEscape(x), url.PathEscape(y))

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

	oauthToken := r.Header.Get("Authorization")

	// Construct the URL for the Olamaps API request
	url := fmt.Sprintf(app.DistanceMatrixURL,
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
	apiURL := fmt.Sprintf(app.ArrayOfDataURL, url.QueryEscape(datasetName))

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
	if oauthToken == "" {
		http.Error(w, "Missing OAuth token", http.StatusUnauthorized)
		return
	}

	// Construct the URL for the Olamaps API request
	apiURL := fmt.Sprintf(app.StyleDetailsURL, url.QueryEscape(styleName))

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

func GetMapStyleHandler(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	oauthToken := r.Header.Get("Authorization")
	if oauthToken == "" {
		http.Error(w, "Missing OAuth token", http.StatusUnauthorized)
		return
	}

	apiURL := app.MapStyleURL

	// Define a variable to hold the API response
	var apiResponse []map[string]interface{}
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

func GetPlaceDetailHandler(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	placeID := r.URL.Query().Get("place_id")
	if placeID == "" {
		http.Error(w, "Missing place_id query parameter", http.StatusBadRequest)
		return
	}

	oauthToken := r.Header.Get("Authorization")
	if oauthToken == "" {
		http.Error(w, "Missing OAuth token", http.StatusUnauthorized)
		return
	}

	apiURL := fmt.Sprintf(app.PlaceDetailURL, placeID) // Replace with your actual endpoint

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}
	requestID := uuid.New().String()

	// Make the external request
	err := internal.MakeExternalRequest("GET", apiURL, requestID, oauthToken, &apiResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the API response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(apiResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func GetNearBySearchHandler(w http.ResponseWriter, r *http.Request) {
	layers := r.URL.Query().Get("layers")
	location := r.URL.Query().Get("location")
	if layers == "" || location == "" {
		http.Error(w, "Missing required query parameters", http.StatusBadRequest)
		return
	}

	types := r.URL.Query().Get("types")
	radius := r.URL.Query().Get("radius")
	strictbounds := r.URL.Query().Get("strictbounds")
	withCentroid := r.URL.Query().Get("withCentroid")
	limit := r.URL.Query().Get("limit")

	oauthToken := r.Header.Get("Authorization")
	if oauthToken == "" {
		http.Error(w, "Missing OAuth token", http.StatusUnauthorized)
		return
	}

	apiURL := fmt.Sprintf(app.NearBySearchURL, layers, location, types, radius, strictbounds, withCentroid, limit)
	requestID := uuid.New().String()

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}

	// Make the external request
	err := internal.MakeExternalRequest("GET", apiURL, requestID, oauthToken, &apiResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the API response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(apiResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func GetTextSearchHandler(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	input := r.URL.Query().Get("input")
	if input == "" {
		http.Error(w, "Missing input query parameter", http.StatusBadRequest)
		return
	}

	// Extract optional query parameters
	location := r.URL.Query().Get("location")
	radius := r.URL.Query().Get("radius")
	types := r.URL.Query().Get("types")
	size := r.URL.Query().Get("size")

	oauthToken := r.Header.Get("Authorization")
	if oauthToken == "" {
		http.Error(w, "Missing OAuth token", http.StatusUnauthorized)
		return
	}

	// Construct the API URL
	apiURL := fmt.Sprintf(app.TextSearchURL, url.QueryEscape(input), location, radius, types, size)

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}
	requestID := uuid.New().String()

	// Make the external request
	err := internal.MakeExternalRequest("GET", apiURL, requestID, oauthToken, &apiResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the API response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(apiResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func GetSnapToRoadHandler(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	points := r.URL.Query().Get("points")
	if points == "" {
		http.Error(w, "Missing points query parameter", http.StatusBadRequest)
		return
	}

	// Optional parameters
	enhancePath := r.URL.Query().Get("enhancePath")

	// Headers
	oauthToken := r.Header.Get("Authorization")
	if oauthToken == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	// Build URL
	apiURL := fmt.Sprintf(app.SnapToRoadURL, url.Values{
		"points":      {points},
		"enhancePath": {enhancePath},
	}.Encode())

	// Make the external request
	var apiResponse map[string]interface{}
	requestID := uuid.New().String()
	err := internal.MakeExternalRequest("GET", apiURL, requestID, oauthToken, &apiResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the API response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(apiResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func GetNearestRoadsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	points := r.URL.Query().Get("points")
	if points == "" {
		http.Error(w, "Missing points query parameter", http.StatusBadRequest)
		return
	}

	radius := r.URL.Query().Get("radius")
	if radius == "" {
		radius = "500" // Default radius
	}

	// Extract headers
	oauthToken := r.Header.Get("Authorization")
	if oauthToken == "" {
		http.Error(w, "Missing OAuth token", http.StatusUnauthorized)
		return
	}

	// Generate request and correlation IDs
	requestID := uuid.New().String()

	// Build the API URL
	apiURL := fmt.Sprintf(app.NearestRoadsURL, points, radius)

	// Define a variable to hold the API response
	var apiResponse map[string]interface{}

	// Make the external request
	err := internal.MakeExternalRequest("GET", apiURL, requestID, oauthToken, &apiResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the API response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(apiResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func GetStaticMapImageCenterHandler(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	styleName := r.URL.Query().Get("styleName")
	longitudeStr := r.URL.Query().Get("longitude")
	latitudeStr := r.URL.Query().Get("latitude")
	zoomLevelStr := r.URL.Query().Get("zoom")
	imageWidthStr := r.URL.Query().Get("width")
	imageHeightStr := r.URL.Query().Get("height")
	imageFormat := r.URL.Query().Get("format")

	// Validate required parameters
	if styleName == "" || longitudeStr == "" || latitudeStr == "" || zoomLevelStr == "" || imageWidthStr == "" || imageHeightStr == "" || imageFormat == "" {
		http.Error(w, "Missing required query parameters", http.StatusBadRequest)
		return
	}

	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		http.Error(w, "Invalid longitude value", http.StatusBadRequest)
		return
	}

	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		http.Error(w, "Invalid latitude value", http.StatusBadRequest)
		return
	}

	zoomLevel, err := strconv.Atoi(zoomLevelStr)
	if err != nil {
		http.Error(w, "Invalid zoom level value", http.StatusBadRequest)
		return
	}

	imageWidth, err := strconv.Atoi(imageWidthStr)
	if err != nil {
		http.Error(w, "Invalid image width value", http.StatusBadRequest)
		return
	}

	imageHeight, err := strconv.Atoi(imageHeightStr)
	if err != nil {
		http.Error(w, "Invalid image height value", http.StatusBadRequest)
		return
	}

	// Extract headers
	oauthToken := r.Header.Get("Authorization")
	if oauthToken == "" {
		http.Error(w, "Missing OAuth token", http.StatusUnauthorized)
		return
	}

	// Optional parameters
	markers := r.URL.Query()["marker"]
	path := r.URL.Query().Get("path")

	// Construct the API URL
	apiURL := fmt.Sprintf(app.StaticMapImageCenterURL,
		url.QueryEscape(styleName), longitude, latitude, zoomLevel, imageWidth, imageHeight, imageFormat)

	// Construct query parameters
	queryParams := url.Values{}
	if len(markers) > 0 {
		queryParams.Add("marker", strings.Join(markers, ","))
	}

	if path != "" {
		queryParams.Add("path", path)
	}

	if len(queryParams) > 0 {
		apiURL += "?" + queryParams.Encode()
	}

	// Generate request ID if not provided
	xRequestID := r.Header.Get("X-Request-Id")
	if xRequestID == "" {
		xRequestID = uuid.New().String()
	}

	// Create and send the external request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		http.Error(w, "Failed to create new request", http.StatusInternalServerError)
		return
	}

	req.Header.Add("X-Request-Id", xRequestID)
	if oauthToken != "" {
		req.Header.Add("Authorization", oauthToken)
	}
	if xCorrelationID := r.Header.Get("X-Correlation-Id"); xCorrelationID != "" {
		req.Header.Add("X-Correlation-Id", xCorrelationID)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to make external request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to fetch image", resp.StatusCode)
		return
	}

	// Set the Content-Type header based on image format
	contentType := "image/" + imageFormat
	w.Header().Set("Content-Type", contentType)

	// Write the image data to the response writer
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Failed to write image data", http.StatusInternalServerError)
		return
	}
}

func GetStaticMapImageBoundedHandler(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	styleName := r.URL.Query().Get("styleName")
	minXStr := r.URL.Query().Get("min_x")
	minYStr := r.URL.Query().Get("min_y")
	maxXStr := r.URL.Query().Get("max_x")
	maxYStr := r.URL.Query().Get("max_y")
	imageWidthStr := r.URL.Query().Get("width")
	imageHeightStr := r.URL.Query().Get("height")
	imageFormat := r.URL.Query().Get("format")

	// Validate required parameters
	if styleName == "" || minXStr == "" || minYStr == "" || maxXStr == "" || maxYStr == "" || imageWidthStr == "" || imageHeightStr == "" || imageFormat == "" {
		http.Error(w, "Missing required query parameters", http.StatusBadRequest)
		return
	}

	minX, err := strconv.ParseFloat(minXStr, 64)
	if err != nil {
		http.Error(w, "Invalid min_x value", http.StatusBadRequest)
		return
	}

	minY, err := strconv.ParseFloat(minYStr, 64)
	if err != nil {
		http.Error(w, "Invalid min_y value", http.StatusBadRequest)
		return
	}

	maxX, err := strconv.ParseFloat(maxXStr, 64)
	if err != nil {
		http.Error(w, "Invalid max_x value", http.StatusBadRequest)
		return
	}

	maxY, err := strconv.ParseFloat(maxYStr, 64)
	if err != nil {
		http.Error(w, "Invalid max_y value", http.StatusBadRequest)
		return
	}

	imageWidth, err := strconv.Atoi(imageWidthStr)
	if err != nil {
		http.Error(w, "Invalid image width value", http.StatusBadRequest)
		return
	}

	imageHeight, err := strconv.Atoi(imageHeightStr)
	if err != nil {
		http.Error(w, "Invalid image height value", http.StatusBadRequest)
		return
	}

	// Extract headers
	oauthToken := r.Header.Get("Authorization")
	if oauthToken == "" {
		http.Error(w, "Missing OAuth token", http.StatusUnauthorized)
		return
	}

	// Optional parameters
	markers := r.URL.Query()["marker"]
	path := r.URL.Query().Get("path")
	xRequestID := r.Header.Get("X-Request-Id")
	xCorrelationID := r.Header.Get("X-Correlation-Id")

	// Construct the API URL
	apiURL := fmt.Sprintf(app.StaticMapImageBoundedURL,
		url.QueryEscape(styleName), minX, minY, maxX, maxY, imageWidth, imageHeight, imageFormat)

	// Construct query parameters
	queryParams := url.Values{}
	if len(markers) > 0 {
		queryParams.Add("marker", strings.Join(markers, ","))
	}

	if path != "" {
		queryParams.Add("path", path)
	}

	if len(queryParams) > 0 {
		apiURL += "?" + queryParams.Encode()
	}

	// Generate request ID if not provided
	if xRequestID == "" {
		xRequestID = uuid.New().String()
	}

	// Make the external request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		http.Error(w, "Failed to create new request", http.StatusInternalServerError)
		return
	}

	req.Header.Add("X-Request-Id", xRequestID)
	if oauthToken != "" {
		req.Header.Add("Authorization", oauthToken)
	}
	if xCorrelationID != "" {
		req.Header.Add("X-Correlation-Id", xCorrelationID)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to make external request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to fetch image", resp.StatusCode)
		return
	}

	// Set the Content-Type header based on image format
	contentType := "image/" + imageFormat
	w.Header().Set("Content-Type", contentType)

	// Write the image data to the response writer
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Failed to write image data", http.StatusInternalServerError)
		return
	}
}

func StaticMapImageHandler(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	styleName := r.URL.Query().Get("styleName")
	imageWidthStr := r.URL.Query().Get("width")
	imageHeightStr := r.URL.Query().Get("height")
	imageFormat := r.URL.Query().Get("format")
	path := r.URL.Query().Get("path")
	markers := r.URL.Query()["marker"]

	// Validate required parameters
	if styleName == "" || imageWidthStr == "" || imageHeightStr == "" || imageFormat == "" {
		http.Error(w, "Missing required query parameters", http.StatusBadRequest)
		return
	}

	imageWidth, err := strconv.Atoi(imageWidthStr)
	if err != nil {
		http.Error(w, "Invalid image width value", http.StatusBadRequest)
		return
	}

	imageHeight, err := strconv.Atoi(imageHeightStr)
	if err != nil {
		http.Error(w, "Invalid image height value", http.StatusBadRequest)
		return
	}

	// Construct the API URL
	apiURL := fmt.Sprintf(app.StaticMapImageURL, url.QueryEscape(styleName), imageWidth, imageHeight, imageFormat)

	// Construct query parameters
	queryParams := url.Values{}
	if len(markers) > 0 {
		queryParams.Add("marker", strings.Join(markers, ","))
	}

	if path != "" {
		queryParams.Add("path", path)
	}

	if len(queryParams) > 0 {
		apiURL += "?" + queryParams.Encode()
	}
	// Create a new request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to send request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Handle response
	switch resp.StatusCode {
	case http.StatusOK:
		// Write the image data to the response writer
		w.Header().Set("Content-Type", "image/"+imageFormat)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response body", http.StatusInternalServerError)
			return
		}
		_, err = w.Write(body)
		if err != nil {
			http.Error(w, "Failed to write response body", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, fmt.Sprintf("Error: %s", resp.Status), resp.StatusCode)
	}
}
