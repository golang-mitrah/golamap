package resources

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

func BadRequestTestCase(handlerName string) (responseCode int, err error) {
	switch handlerName {
	case "GetDirectionsHandler":
		server := httptest.NewServer(http.HandlerFunc(GetDirectionsHandler))
		defer server.Close()

		urlPath := `/routing/v1/direction?&destination=8`
		req, err := http.NewRequest(http.MethodGet, server.URL+urlPath, nil)
		if err != nil {
			return 0, err
		}
		req.Header.Set("Authorization", "mock-auth")
		rr := httptest.ResponseRecorder{}
		if req.Header.Get("Authorization") != "" && (req.URL.Query().Get("origin") == "" || req.URL.Query().Get("destination") == "") {
			rr.Code = 400
		}
		return rr.Code, nil
	case "GeoCodeHandler":
		server := httptest.NewServer(http.HandlerFunc(GeoCodeHandler))
		defer server.Close()
		url := fmt.Sprintf("/places/v1/geocode?address=%s&language=en", "mockaddress")
		req, err := http.NewRequest(http.MethodGet, url, nil)
		req.Header.Set("Authorization", "mock-auth")
		if err != nil {
			return 0, err
		}
		rr := httptest.ResponseRecorder{}
		if req.Header.Get("Authorization") != "" && (req.URL.Query().Get("address") != "" &&
			req.URL.Query().Get("bounds") != "" ||
			req.URL.Query().Get("language") != "") {
			rr.Code = 400
		}
		return rr.Code, nil
	case "PlaceAutoCompleteHandler":
		server := httptest.NewServer(http.HandlerFunc(PlaceAutoCompleteHandler))
		defer server.Close()

		req, err := http.NewRequest(http.MethodGet, "/places/v1/autocomplete", nil)
		req.Header.Set("Authorization", "mock-auth")
		if err != nil {
			return 0, err
		}
		rr := httptest.ResponseRecorder{}
		if req.Header.Get("Authorization") != "" && req.URL.Query().Get("input") == "" {
			rr.Code = 400
		}
		return rr.Code, nil
	case "GetPbfFileHandler":
		server := httptest.NewServer(http.HandlerFunc(GetPbfFileHandler))
		defer server.Close()
		url := server.URL + `/tiles/vector/v1/data/test_dataset/1/2/`
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return 0, nil
		}
		req.Header.Set("Authorization", "mock-auth")
		rr := httptest.ResponseRecorder{}
		segments := strings.Split(url, "/")

		// Ensure there are at least 4 segments for the last 4 nodes
		if len(segments) < 5 {
			fmt.Println("URL does not contain enough path segments")
			return 0, fmt.Errorf("url does not contain enough path segments")
		}
		// Extract the last 4 segments
		datasetName := segments[len(segments)-4]
		z := segments[len(segments)-3]
		x := segments[len(segments)-2]
		y := segments[len(segments)-1]
		if req.Header.Get("Authorization") != "" && (datasetName == "" ||
			z == "" ||
			y == "" ||
			x == "") {
			rr.Code = 400
		}
		return rr.Code, nil
	case "ArrayOfDataHandler":
		server := httptest.NewServer(http.HandlerFunc(ArrayOfDataHandler))
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, server.URL+`//api.olamaps.io/tiles/vector/v1/data`, nil)
		if err != nil {
			return 0, err
		}
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		req.Header.Set("Authorization", "mock-Auth")
		defer resp.Body.Close()
		rr := httptest.ResponseRecorder{}
		if req.URL.Query().Get("dataset_name") == "" && req.Header.Get("Authorization") != "" {
			rr.Code = 400
		}
		return rr.Code, nil
	case "ReverseGeocodeHandler":
		server := httptest.NewServer(http.HandlerFunc(ReverseGeocodeHandler))
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, server.URL+"/reverse-geocode", nil)
		if err != nil {
			return 0, err
		}
		client := http.Client{}
		resp, err := client.Do(req)
		rr := httptest.ResponseRecorder{}
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		req.Header.Add("Authorization", "mock-Auth")
		if req.URL.Query().Get("latlng") == "" && req.Header.Get("Authorization") != "" {
			rr.Code = 400
		}
		return rr.Code, nil
	case "GetStyleDetailsHandler":
		server := httptest.NewServer(http.HandlerFunc(GetStyleDetailsHandler))
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, server.URL+"//api.olamaps.io/tiles/vector/v1/styles/mock/style.json", nil)
		if err != nil {
			return 0, nil
		}
		req.Header.Set("Authorization", "mock-auth")
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		rr := httptest.ResponseRecorder{}
		if req.URL.Query().Get("style_name") == "" && req.Header.Get("Authorization") != "" {
			rr.Code = 400
		}
		return rr.Code, nil
	case "GetMapStyleHandler":
		server := httptest.NewServer(http.HandlerFunc(GetMapStyleHandler))
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, server.URL+"/tiles/vector/v1/styles.json", nil)
		if err != nil {
			return 0, err
		}
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		req.Header.Set("Authorization", "mock-auth")
		rr := httptest.ResponseRecorder{}
		rr.Body = nil
		if req.Header.Get("Authorization") != "" && rr.Body == nil {
			rr.Code = 500
		}
		return rr.Code, nil
	case "GetPlaceDetailHandler":
		server := httptest.NewServer(http.HandlerFunc(GetPlaceDetailHandler))
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, server.URL+"/places/detail", nil)
		if err != nil {
			return 0, err
		}
		req.Header.Set("Authorization", "mock-auth")
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		rr := httptest.ResponseRecorder{}
		if req.Header.Get("Authorization") != "" && req.URL.Query().Get("place_id") == "" {
			rr.Code = 400
		}
		return rr.Code, nil
	case "GetNearBySearchHandler":
		server := httptest.NewServer(http.HandlerFunc(GetNearBySearchHandler))
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, server.URL+`/api.olamaps.io/places/v1/nearbysearch?input=6&location=loc&radius=7&types=mocktype&size=3`, nil)
		if err != nil {
			return 0, err
		}
		req.Header.Set("Authorization", "mock-auth")
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		rr := httptest.ResponseRecorder{}
		if req.Header.Get("Authorization") != "" && req.URL.Query().Get("layers") == "" {
			rr.Code = 400
		}

		return rr.Code, nil
	case "GetSnapToRoadHandler":
		server := httptest.NewServer(http.HandlerFunc(GetSnapToRoadHandler))
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, server.URL+`/api.olamaps.io/places/v1/textsearch?`, nil)
		if err != nil {
			return 0, err
		}
		req.Header.Set("Authorization", "mock-auth")
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		rr := httptest.ResponseRecorder{}
		if req.Header.Get("Authorization") != "" && req.URL.Query().Get("points") == "" {
			rr.Code = 400
		}
		return rr.Code, nil
	case "GetNearestRoadsHandler":
		server := httptest.NewServer(http.HandlerFunc(GetNearestRoadsHandler))
		url := fmt.Sprintf("/api.olamaps.io/routing/v1/nearestRoads")
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, server.URL+url, nil)
		if err != nil {
			return 0, err
		}

		req.Header.Set("Authorization", "mock-auth")
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		rr := httptest.ResponseRecorder{}
		if req.Header.Get("Authorization") != "" && req.URL.Query().Get("points") == "" {
			rr.Code = 400
		}
		return rr.Code, nil
	case "GetStaticMapImageCenterHandler":
		server := httptest.NewServer(http.HandlerFunc(GetStaticMapImageCenterHandler))
		defer server.Close()
		url := fmt.Sprintf("/api.olamaps.io/tiles/v1/styles/%s/static/%f,%f,%d/%dx%d.%s", "mock", 4.4, 4.5, 4, 5, 6, "jpeg")
		req, err := http.NewRequest(http.MethodGet, server.URL+url, nil)

		if err != nil {
			return 0, err
		}
		req.Header.Set("Authorization", "mock-auth")
		req.URL.Query().Set("longitude", "7.7")
		req.URL.Query().Set("latitude", "7.8")
		req.URL.Query().Set("zoom", "8.8")
		req.URL.Query().Set("width", "8.85")
		req.URL.Query().Set("height", "8.81")
		req.URL.Query().Set("format", ".jpeg")
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		rr := httptest.ResponseRecorder{}
		if req.Header.Get("Authorization") != "" && req.URL.Query().Get("styleName") == "" {
			rr.Code = 400
		}
		return rr.Code, nil
	case "GetStaticMapImageBoundedHandler":
		server := httptest.NewServer(http.HandlerFunc(GetStaticMapImageBoundedHandler))
		url := fmt.Sprintf("/api.olamaps.io/tiles/v1/styles/%s/static/%f,%f,%f,%f/%dx%d.%s",
			"", 77.611182859373, 12.93219851203095, 77.61513567417848, 12.935739723360513, 800, 600, "png")

		paramValue := staticMapImageBoundedHandlerRegex(url)
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, server.URL+url, nil)
		if err != nil {
			return 0, nil
		}
		req.Header.Set("Authorization", "mock-auth")
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, nil
		}
		defer resp.Body.Close()
		rr := httptest.ResponseRecorder{}
		if req.Header.Get("Authorization") != "" && paramValue["style"] == "" {
			rr.Code = 400
		}
		return rr.Code, nil
	case "GetDistanceMatrixHandler":
		server := httptest.NewServer(http.HandlerFunc(GetDistanceMatrixHandler))
		// Construct the URL for the Olamaps API request
		url := fmt.Sprint("/api.olamaps.io/routing/v1/distanceMatrix?origins=&destinations=")
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, server.URL+url, nil)
		if err != nil {
			return 0, err
		}

		req.Header.Set("Authorization", "mock-auth")
		rr := httptest.ResponseRecorder{}
		query := req.URL.Query()
		origins := query.Get("origins")
		destinations := query.Get("destinations")
		if req.Header.Get("Authorization") != "" && (origins == "" || destinations == "") {
			rr.Code = 400
		}
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		return rr.Code, err
	case "StaticMapImageHandler":
		server := httptest.NewServer(http.HandlerFunc(StaticMapImageHandler))
		url := "/tiles/v1/styles/static/auto/87x90.png"
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, server.URL+url, nil)
		if err != nil {
			return 0, err
		}
		pathValues := urlPathParamRegexStaticTiles(url)
		req.Header.Set("Authorization", "mock-auth")
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		rr := httptest.ResponseRecorder{}
		if req.Header.Get("Authorization") != "" && pathValues["styleName"] == "" {
			rr.Code = 400
		}
		return rr.Code, nil
	case "GetTextSearchHandler":
		server := httptest.NewServer(http.HandlerFunc(GetTextSearchHandler))
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, server.URL+`/api.olamaps.io/places/v1/textsearch?location=loc&radius=7&types=mocktype&size=3`, nil)
		if err != nil {
			return 0, err
		}
		req.Header.Set("Authorization", "mock-auth")
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		rr := httptest.ResponseRecorder{}
		if req.Header.Get("Authorization") != "" && req.URL.Query().Get("input") == "" {
			rr.Code = 400
		}
		return rr.Code, nil
	}
	return 0, fmt.Errorf("Invalid Handler")
}

func successAndUnauthorizedCase(handlerName, testCondition string) (responseCode int, err error) {
	switch handlerName {
	case "GetDirectionsHandler":
		server := httptest.NewServer(http.HandlerFunc(GetDirectionsHandler))
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, server.URL+"/routing/v1/direction?origin=7&destination=8", nil)
		rr := httptest.ResponseRecorder{}
		if err != nil {
			return 0, err
		}
		if testCondition == "success" {
			req.Header.Set("Authorization", "mock-auth")
			if req.Header.Get("Authorization") != "" && req.URL.Query().Get("origin") != "" && req.URL.Query().Get("destination") != "" {
				rr.Code = 200
			}
		} else {
			if req.Header.Get("Authorization") == "" && req.URL.Query().Get("origin") != "" && req.URL.Query().Get("destination") != "" {
				rr.Code = 401
			}
		}
		return rr.Code, nil
	case "GeoCodeHandler":
		server := httptest.NewServer(http.HandlerFunc(GeoCodeHandler))
		defer server.Close()
		url := fmt.Sprintf("/places/v1/geocode?address=%s&bounds=%s&language=en", "mockaddress", "mock-bounds")
		req, err := http.NewRequest(http.MethodGet, url, nil)

		if err != nil {
			return 0, err
		}
		rr := httptest.ResponseRecorder{}
		if testCondition == "success" {
			req.Header.Set("Authorization", "mock-auth")
			if req.Header.Get("Authorization") != "" && (req.URL.Query().Get("address") != "" &&
				req.URL.Query().Get("bounds") != "" &&
				req.URL.Query().Get("language") != "") {
				rr.Code = 200
			}
		} else {
			if req.Header.Get("Authorization") == "" && (req.URL.Query().Get("address") != "" &&
				req.URL.Query().Get("bounds") != "" &&
				req.URL.Query().Get("language") != "") {
				rr.Code = 401
			}
		}
		return rr.Code, nil

	case "PlaceAutoCompleteHandler":
		server := httptest.NewServer(http.HandlerFunc(PlaceAutoCompleteHandler))
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, "/places/v1/autocomplete?input=7", nil)
		if err != nil {
			return 0, err
		}
		rr := httptest.ResponseRecorder{}
		if testCondition == "success" {
			req.Header.Set("Authorization", "mock-auth")
			if req.Header.Get("Authorization") != "" && req.URL.Query().Get("input") != "" {
				rr.Code = 200
			}
		} else {
			if req.Header.Get("Authorization") == "" && req.URL.Query().Get("input") != "" {
				rr.Code = 401
			}
		}
		return rr.Code, nil

	case "GetPbfFileHandler":
		server := httptest.NewServer(http.HandlerFunc(GetPbfFileHandler))
		defer server.Close()
		url := `/tiles/vector/v1/data/test_dataset/1/2/3`
		req, err := http.NewRequest(http.MethodGet, server.URL+url, nil)
		if err != nil {
			return 0, err
		}

		rr := httptest.ResponseRecorder{}
		segments := strings.Split(url, "/")
		// Ensure there are at least 4 segments for the last 4 nodes
		if len(segments) < 5 {
			return 400, fmt.Errorf("url does not contain enough path segments")
		}
		// Extract the last 4 segments
		datasetName := segments[len(segments)-4]
		z := segments[len(segments)-3]
		x := segments[len(segments)-2]
		y := segments[len(segments)-1]
		if testCondition == "success" {
			req.Header.Set("Authorization", "mock-auth")
			if req.Header.Get("Authorization") != "" && (datasetName != "" &&
				z != "" &&
				y != "" &&
				x != "") {
				rr.Code = 200
			}

		} else {
			if req.Header.Get("Authorization") == "" && (datasetName != "" &&
				z != "" &&
				y != "" &&
				x != "") {
				rr.Code = 401
			}
		}
		return rr.Code, nil
	case "ArrayOfDataHandler":
		server := httptest.NewServer(http.HandlerFunc(ArrayOfDataHandler))
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, server.URL+`//api.olamaps.io/tiles/vector/v1/data?dataset_name=mockdata`, nil)
		if err != nil {
			return 0, err
		}
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		rr := httptest.ResponseRecorder{}
		if testCondition == "success" {
			req.Header.Set("Authorization", "mock-Auth")
			if req.URL.Query().Get("dataset_name") != "" && req.Header.Get("Authorization") == "mock-Auth" {
				rr.Code = 200
			}
		} else {
			if req.URL.Query().Get("dataset_name") != "" && req.Header.Get("Authorization") == "" {
				rr.Code = 401
			}
		}
		return rr.Code, nil
	case "ReverseGeocodeHandler":
		server := httptest.NewServer(http.HandlerFunc(ReverseGeocodeHandler))
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, server.URL+"/reverse-geocode?latlng=mockValue", nil)
		if err != nil {
			return 0, err
		}
		client := http.Client{}
		resp, err := client.Do(req)
		rr := httptest.ResponseRecorder{}
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		if testCondition == "success" {
			req.Header.Add("Authorization", "mock-Auth")
			if req.URL.Query().Get("latlng") != "" && req.Header.Get("Authorization") != "" {
				rr.Code = 200
			}
		} else {
			if req.URL.Query().Get("latlng") != "" && req.Header.Get("Authorization") == "" {
				rr.Code = 401
			}
		}
		return rr.Code, nil
	case "GetStyleDetailsHandler":
		server := httptest.NewServer(http.HandlerFunc(GetStyleDetailsHandler))
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, server.URL+"//api.olamaps.io/tiles/vector/v1/styles/mock/style.json?style_name=mockstyle", nil)
		if err != nil {
			return 0, err
		}
		if testCondition == "success" {
			req.Header.Set("Authorization", "mock-auth")
		}
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		rr := httptest.ResponseRecorder{}
		if req.URL.Query().Get("style_name") != "" && req.Header.Get("Authorization") != "" {
			rr.Code = 200
		} else if req.URL.Query().Get("style_name") != "" && req.Header.Get("Authorization") == "" {
			rr.Code = 401
		}
		return rr.Code, nil
	case "GetMapStyleHandler":
		server := httptest.NewServer(http.HandlerFunc(GetMapStyleHandler))
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, server.URL+"/tiles/vector/v1/styles.json", nil)
		if err != nil {
			return 0, err
		}
		if testCondition == "success" {
			req.Header.Set("Authorization", "mock-auth")
		}

		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		rr := httptest.ResponseRecorder{}
		if req.Header.Get("Authorization") != "" {
			rr.Code = 200
		} else if req.Header.Get("Authorization") == "" {
			rr.Code = 401
		}
		return rr.Code, nil
	case "GetPlaceDetailHandler":
		server := httptest.NewServer(http.HandlerFunc(GetPlaceDetailHandler))
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, server.URL+"/places/detail?place_id=3", nil)
		if err != nil {
			return 0, err
		}
		if testCondition == "success" {
			req.Header.Set("Authorization", "mock-auth")
		}

		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		rr := httptest.ResponseRecorder{}
		if req.Header.Get("Authorization") != "" && req.URL.Query().Get("place_id") != "" {
			rr.Code = 200
		} else if req.Header.Get("Authorization") == "" && req.URL.Query().Get("place_id") != "" {
			rr.Code = 401
		}
		return rr.Code, nil
	case "GetNearBySearchHandler":
		server := httptest.NewServer(http.HandlerFunc(GetNearBySearchHandler))
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, server.URL+`/api.olamaps.io/places/v1/textsearch?layers=6&location=loc&radius=7&types=mocktype&size=3`, nil)
		if err != nil {
			return 0, err
		}
		if testCondition == "success" {
			req.Header.Set("Authorization", "mock-auth")
		}
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		rr := httptest.ResponseRecorder{}
		if req.Header.Get("Authorization") != "" && req.URL.Query().Get("location") != "" && req.URL.Query().Get("layers") != "" {
			rr.Code = 200
		} else if req.Header.Get("Authorization") == "" && req.URL.Query().Get("location") != "" && req.URL.Query().Get("layers") != "" {
			rr.Code = 401

		}
		return rr.Code, nil
	case "GetSnapToRoadHandler":
		server := httptest.NewServer(http.HandlerFunc(GetSnapToRoadHandler))
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, server.URL+`/api.olamaps.io/places/v1/textsearch?points=6&enhancepath=mock-path`, nil)
		if err != nil {
			return 0, err
		}
		if testCondition == "success" {
			req.Header.Set("Authorization", "mock-auth")
		}
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		rr := httptest.ResponseRecorder{}
		if req.Header.Get("Authorization") != "" && req.URL.Query().Get("points") != "" {
			rr.Code = 200
		} else if req.Header.Get("Authorization") == "" && req.URL.Query().Get("points") != "" {
			rr.Code = 401
		}
		return rr.Code, nil
	case "GetNearestRoadsHandler":
		server := httptest.NewServer(http.HandlerFunc(GetNearestRoadsHandler))
		url := fmt.Sprintf("/api.olamaps.io/routing/v1/nearestRoads?points=%s&radius=%s",
			"37.7749", "66.8")
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, server.URL+url, nil)
		if err != nil {
			return 0, err
		}
		if testCondition == "success" {
			req.Header.Set("Authorization", "mock-auth")
		}
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		rr := httptest.ResponseRecorder{}
		if req.Header.Get("Authorization") != "" && req.URL.Query().Get("points") != "" {
			rr.Code = 200
		} else if req.Header.Get("Authorization") == "" && req.URL.Query().Get("points") != "" {
			rr.Code = 401
		}
		return rr.Code, nil
	case "GetStaticMapImageCenterHandler":
		server := httptest.NewServer(http.HandlerFunc(GetStaticMapImageCenterHandler))
		defer server.Close()
		url := "/tiles/center?styleName=default&min_x=8&&min_y=9&max_x&max_y&height=9.8&width=9.0&format=jpeg"
		req, err := http.NewRequest(http.MethodGet, server.URL+url, nil)

		if err != nil {
			return 0, err
		}
		if testCondition == "success" {
			req.Header.Set("Authorization", "mock-auth")
		}

		req.URL.Query().Set("styleName", "mock-styleName")
		req.URL.Query().Set("longitude", "7.7")
		req.URL.Query().Set("latitude", "7.8")
		req.URL.Query().Set("zoom", "8.8")
		req.URL.Query().Set("width", "8.85")
		req.URL.Query().Set("height", "8.81")
		req.URL.Query().Set("format", ".jpeg")
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		rr := httptest.ResponseRecorder{}
		if req.Header.Get("Authorization") != "" && req.URL.Query().Get("styleName") != "" {
			rr.Code = 200
		} else if req.Header.Get("Authorization") == "" && req.URL.Query().Get("styleName") != "" {
			rr.Code = 401
		}
		return rr.Code, nil
	case "GetStaticMapImageBoundedHandler":
		server := httptest.NewServer(http.HandlerFunc(GetStaticMapImageBoundedHandler))
		url := fmt.Sprintf("/api.olamaps.io/tiles/v1/styles/%s/static/%f,%f,%f,%f/%dx%d.%s",
			"default-light-standard", 77.611182859373, 12.93219851203095, 77.61513567417848, 12.935739723360513, 800, 600, "png")
		paramValue := staticMapImageBoundedHandlerRegex(url)
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, server.URL+url, nil)
		if err != nil {
			return 0, err
		}
		if testCondition == "success" {
			req.Header.Set("Authorization", "mock-auth")
		}
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		rr := httptest.ResponseRecorder{}
		if req.Header.Get("Authorization") != "" && paramValue["style"] == "default-light-standard" {
			rr.Code = 200
		} else if req.Header.Get("Authorization") == "" && paramValue["style"] == "default-light-standard" {
			rr.Code = 401
		}
		return rr.Code, nil
	case "GetDistanceMatrixHandler":
		server := httptest.NewServer(http.HandlerFunc(GetDistanceMatrixHandler))
		// Construct the URL for the Olamaps API request
		url := fmt.Sprintf("/api.olamaps.io/routing/v1/distanceMatrix?origins=%s&destinations=%s",
			"mock-origin", "mock-destination")
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, server.URL+url, nil)
		if err != nil {
			return 0, nil
		}
		if testCondition == "success" {
			req.Header.Set("Authorization", "mock-auth")
		}
		rr := httptest.ResponseRecorder{}
		query := req.URL.Query()
		origins := query.Get("origins")
		destinations := query.Get("destinations")
		if req.Header.Get("Authorization") != "" && origins != "" && destinations != "" {
			rr.Code = 200
		} else if req.Header.Get("Authorization") == "" && origins != "" && destinations != "" {
			rr.Code = 401
		}
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, nil
		}
		defer resp.Body.Close()
		return rr.Code, nil
	case "StaticMapImageHandler":
		server := httptest.NewServer(http.HandlerFunc(StaticMapImageHandler))
		url := "/tiles/v1/styles/default-light-standard/static/auto/87x90.png"
		defer server.Close()
		pathVAlues := urlPathParamRegexStaticTiles(url)
		req, err := http.NewRequest(http.MethodGet, server.URL+url, nil)
		if err != nil {
			return 0, err
		}
		if testCondition == "success" {
			req.Header.Set("Authorization", "mock-auth")
		}
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		rr := httptest.ResponseRecorder{}
		if req.Header.Get("Authorization") != "" && pathVAlues["styleName"] == "default-light-standard" {
			rr.Code = 200
		} else if req.Header.Get("Authorization") == "" && pathVAlues["styleName"] == "default-light-standard" {
			rr.Code = 401

		}
		return rr.Code, nil
	case "GetTextSearchHandler":
		server := httptest.NewServer(http.HandlerFunc(GetTextSearchHandler))
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, server.URL+`/api.olamaps.io/places/v1/textsearch?input=6&location=loc&radius=7&types=mocktype&size=3`, nil)
		if err != nil {
			return 0, err
		}
		if testCondition == "success" {
			req.Header.Set("Authorization", "mock-auth")
		}
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		rr := httptest.ResponseRecorder{}
		if req.Header.Get("Authorization") != "" && req.URL.Query().Get("location") != "" && req.URL.Query().Get("input") != "" {
			rr.Code = 200
		} else if req.Header.Get("Authorization") == "" && req.URL.Query().Get("location") != "" && req.URL.Query().Get("input") != "" {
			rr.Code = 401
		}
		return rr.Code, nil

	}

	return 0, fmt.Errorf("invalid handler")
}

type MockConfigUrl struct {
	HandlerName string
	HandlerUrl  string
}
type MockConfig interface {
	LoadConfig()
}

func (config *MockConfigUrl) LoadConfig() {
	log.Println("Loading environment variables from .env file")
	err := godotenv.Load("mockurl/mockurls.txt")
	if err != nil {
		log.Fatalf("Error loading .env file = %v ", err)
	}

	config.HandlerUrl = os.Getenv(config.HandlerName)
}
func TestGetDirectionsHandler(t *testing.T) {
	t.Run("Success Response", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GetDirectionsHandler", "success")
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		if statusCode != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

	})
	t.Run("Bad Request", func(t *testing.T) {
		statusCode, err := BadRequestTestCase("GetDirectionsHandler")
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		if statusCode != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusBadRequest)
		}

	})
	t.Run("unauthorized error", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GetDirectionsHandler", "unauthorized")
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		if statusCode != http.StatusUnauthorized {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnauthorized)
		}

	})
}

func TestGeoCodeHandler(t *testing.T) {
	t.Run("success Response", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GeoCodeHandler", "success")
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		if statusCode != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

	})

	t.Run("Bad Request", func(t *testing.T) {
		statusCode, err := BadRequestTestCase("GeoCodeHandler")
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		if statusCode != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusBadRequest)
		}

	})
	t.Run("unauthorized error", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GeoCodeHandler", "unauthorized")
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		if statusCode != http.StatusUnauthorized {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnauthorized)
		}

	})

}
func TestPlaceAutoCompleteHandler(t *testing.T) {
	t.Run("success Response", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("PlaceAutoCompleteHandler", "success")
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		if statusCode != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

	})

	t.Run("Bad Request", func(t *testing.T) {
		statusCode, err := BadRequestTestCase("PlaceAutoCompleteHandler")
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}

		if statusCode != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusBadRequest)
		}

	})
	t.Run("unaAuthorized error", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("PlaceAutoCompleteHandler", "unauthorized")
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		if statusCode != http.StatusUnauthorized {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnauthorized)
		}

	})

}

func TestGetPbfFileHandler(t *testing.T) {
	t.Run("success Response", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GetPbfFileHandler", "unauthorized")
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		if statusCode != http.StatusUnauthorized {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnauthorized)
		}

	})
	t.Run("Bad Request Error", func(t *testing.T) {
		statusCode, err := BadRequestTestCase("GetPbfFileHandler")
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		if statusCode != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusBadRequest)
		}

	})
	t.Run("unAuthorized Error", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GetPbfFileHandler", "unauthorized")
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		if statusCode != http.StatusUnauthorized {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnauthorized)
		}

	})

}

func TestArrayOfDataHandler(t *testing.T) {
	t.Run("success Response", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("ArrayOfDataHandler", "success")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

	})

	t.Run("Bad Request", func(t *testing.T) {
		statusCode, err := BadRequestTestCase("ArrayOfDataHandler")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}

		if statusCode != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusBadRequest)
		}

	})
	t.Run("unauthorized error", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("ArrayOfDataHandler", "unauthorized")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusUnauthorized {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnauthorized)
		}
	})
}

func TestReverseGeocodeHandler(t *testing.T) {
	t.Run("success Response", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("ReverseGeocodeHandler", "success")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

	})

	t.Run("Bad Request", func(t *testing.T) {
		statusCode, err := BadRequestTestCase("ReverseGeocodeHandler")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusBadRequest)
		}

	})
	t.Run("unauthorized error", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("ReverseGeocodeHandler", "unauthorized")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusUnauthorized {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnauthorized)
		}
	})
}

func TestGetStyleDetailsHandler(t *testing.T) {
	t.Run("success Response", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GetStyleDetailsHandler", "success")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

	})
	t.Run("Bad Request", func(t *testing.T) {
		statusCode, err := BadRequestTestCase("GetStyleDetailsHandler")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusBadRequest)
		}

	})
	t.Run("unAuthorized error", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GetStyleDetailsHandler", "unauthorized")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusUnauthorized {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnauthorized)
		}

	})

}
func TestGetMapStyleHandler(t *testing.T) {

	t.Run("success Response", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GetMapStyleHandler", "success")

		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

	})
	t.Run("unauthorized error", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GetMapStyleHandler", "unauthorized")

		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusUnauthorized {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnauthorized)
		}

	})

	t.Run("internal server error", func(t *testing.T) {
		statusCode, err := BadRequestTestCase("GetMapStyleHandler")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusInternalServerError {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusInternalServerError)
		}

	})
}

func TestGetPlaceDetailHandler(t *testing.T) {
	t.Run("success Response", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GetPlaceDetailHandler", "success")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

	})

	t.Run("Bad Request", func(t *testing.T) {
		statusCode, err := BadRequestTestCase("GetPlaceDetailHandler")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusBadRequest)
		}

	})
	t.Run("unauthorized error", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GetPlaceDetailHandler", "unauthorized")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusUnauthorized {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnauthorized)
		}
	})

}
func TestGetTextSearchHandler(t *testing.T) {
	t.Run("success Response", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GetTextSearchHandler", "success")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

	})

	t.Run("Bad Request", func(t *testing.T) {
		statusCode, err := BadRequestTestCase("GetTextSearchHandler")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusBadRequest)
		}

	})
	t.Run("unauthorized error", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GetTextSearchHandler", "unauthorized")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusUnauthorized {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}
	})

}

func TestGetNearBySearchHandler(t *testing.T) {

	t.Run("success Response", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GetNearBySearchHandler", "success")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

	})

	t.Run("unauthorized error", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GetNearBySearchHandler", "unauthorized")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}

		if statusCode != http.StatusUnauthorized {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnauthorized)
		}

	})
	t.Run("BadRequest error", func(t *testing.T) {
		statusCode, err := BadRequestTestCase("GetNearBySearchHandler")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusBadRequest)
		}

	})
}

func TestGetSnapToRoadHandler(t *testing.T) {
	// Create a new HTTP GET request for the /snapToRoad endpoint without the 'points' query parameter.

	t.Run("success Response", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GetSnapToRoadHandler", "success")

		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

	})

	t.Run("BadRequest Error", func(t *testing.T) {
		statusCode, err := BadRequestTestCase("GetSnapToRoadHandler")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusBadRequest)
		}

	})

	t.Run("unauthorized error", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GetSnapToRoadHandler", "unauthorized")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusUnauthorized {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnauthorized)
		}

	})
}
func TestGetNearestRoadsHandler(t *testing.T) {
	t.Run("success Response", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GetNearestRoadsHandler", "success")

		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

	})
	t.Run("Bad Request", func(t *testing.T) {
		statusCode, err := BadRequestTestCase("GetNearestRoadsHandler")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusBadRequest)
		}

	})
	t.Run("unauthorized error", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GetNearestRoadsHandler", "unauthorized")

		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusUnauthorized {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnauthorized)
		}

	})
}

func TestGetStaticMapImageCenterHandler(t *testing.T) {

	t.Run("success Response", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GetStaticMapImageCenterHandler", "success")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

	})
	t.Run("Bad Request", func(t *testing.T) {
		statusCode, err := BadRequestTestCase("GetStaticMapImageCenterHandler")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusBadRequest)
		}

	})

	t.Run("unauthorized error", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GetStaticMapImageCenterHandler", "unauthorized")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusUnauthorized {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnauthorized)
		}

	})
}

func staticMapImageBoundedHandlerRegex(url string) map[string]string {
	pattern := `^/api\.olamaps\.io/tiles/v1/styles/([^/]+)/static/([^,]+),([^,]+),([^,]+),([^/]+)/(\d+)x(\d+)\.(png|jpg|jpeg)$`

	// Compile the regex
	re := regexp.MustCompile(pattern)
	style := ""
	minLon, minLat := "", ""
	maxLon, maxLat := "", ""
	width, height := "", ""
	format := ""
	// Find matches
	matches := re.FindStringSubmatch(url)
	if matches != nil && len(matches) == 9 {
		//log.Fatal("No matches found")
		style = matches[1]
		minLon = matches[2]
		minLat = matches[3]
		maxLon = matches[4]
		maxLat = matches[5]
		width = matches[6]
		height = matches[7]
		format = matches[8]
	} else {
		return map[string]string{}
	}

	return map[string]string{
		"style":  style,
		"minLon": minLon,
		"minLat": minLat,
		"maxLon": maxLon,
		"maxLat": maxLat,
		"width":  width,
		"height": height,
		"format": format,
	}
}
func urlPathParamRegexStaticTiles(url string) map[string]string {
	pattern := `/tiles/v1/styles/([^/]+)/static/auto/(\d+)x(\d+)\.(png|jpg|jpeg)$`

	// Compile the regex
	re := regexp.MustCompile(pattern)

	styleName := ""
	width, height := "", ""
	format := ""
	// Find matches
	matches := re.FindStringSubmatch(url)

	// Extract values from matches
	if len(matches) == 5 && matches != nil {
		styleName = matches[1]
		width = matches[2]
		height = matches[3]
		format = matches[4]
	} else {
		return map[string]string{}
	}

	return map[string]string{
		"styleName": styleName,
		"width":     width,
		"height":    height,
		"format":    format,
	}
}

func TestGetStaticMapImageBoundedHandler(t *testing.T) {

	t.Run("success Response", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GetStaticMapImageBoundedHandler", "success")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

	})
	t.Run("Bad Request Error", func(t *testing.T) {
		statusCode, err := BadRequestTestCase("GetStaticMapImageBoundedHandler")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusBadRequest)
		}

	})
	t.Run("unauthorized error", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GetStaticMapImageBoundedHandler", "unauthorized")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusUnauthorized {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnauthorized)
		}

	})
}

func TestGetDistanceMatrixHandler(t *testing.T) {
	t.Run("success Response", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GetDistanceMatrixHandler", "success")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

	})

	t.Run("Bad Request Error", func(t *testing.T) {
		statusCode, err := BadRequestTestCase("GetDistanceMatrixHandler")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusBadRequest)
		}

	})
	t.Run("unauthorized error", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("GetDistanceMatrixHandler", "unauthorized")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusUnauthorized {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnauthorized)
		}

	})
}

func TestGetStaticMapImageBoundedHandler_MissingOAuthToken(t *testing.T) {
	// Create a new HTTP GET request for the /tiles/center endpoint with all required parameters but no Authorization header.
	req, err := http.NewRequest(http.MethodGet, "/tiles/center?styleName=default&min_x=-122.4194&min_y=37.7749&max_x=-122.4194&max_y=37.7749&width=800&height=600&format=png", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new ResponseRecorder to capture the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetStaticMapImageBoundedHandler)

	// Serve the HTTP request using the handler.
	handler.ServeHTTP(rr, req)

	// Check if the status code is 401 Unauthorized.
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}

	// Check if the response body matches the expected error message.
	expected := "Missing OAuth token\n"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestStaticMapImageHandler(t *testing.T) {
	t.Run("success Response", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("StaticMapImageHandler", "success")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

	})

	t.Run("Bad Request error", func(t *testing.T) {
		StatusCode, err := BadRequestTestCase("StaticMapImageHandler")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if StatusCode != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				StatusCode, http.StatusBadRequest)
		}

	})
	t.Run("unauthorized error", func(t *testing.T) {
		statusCode, err := successAndUnauthorizedCase("StaticMapImageHandler", "unauthorized")
		if err != nil {
			t.Fatalf("Failed to reach server: %v", err)
		}
		if statusCode != http.StatusUnauthorized {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnauthorized)
		}
	})
}
