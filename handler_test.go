package golamap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDirections(t *testing.T) {
	t.Run("Invalid origin & destination", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		_, err := olaMap.GetDirections("", "")
		expectedErr := fmt.Errorf("Missing required query parameters: 'origin' and/or 'destination'")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		_, err := olaMap.GetDirections("mock-origin", "mock-destination")
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)
	})
	t.Run("success", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.Token = "mockToken"
		mocking := &MockStruct{}
		olaMap.HttpService = mocking
		_, err := olaMap.GetDirections("mock-origin", "mock-destination")
		assert.Nil(t, nil, err)
		if mocking.StatusCode != 200 {
			t.Error("expected : ", 200, "got : ", mocking.StatusCode)
		}
		assert.Equal(t, DirectionResponse, mocking.MockBody)

	})
}

func TestPlaceAutoComplete(t *testing.T) {
	t.Run("Invalid input", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		_, err := olaMap.PlaceAutoComplete("")
		expectedErr := fmt.Errorf("Missing required query parameters: 'input'")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		_, err := olaMap.PlaceAutoComplete("mock-input")
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("success", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.Token = "mockToken"
		mocking := &MockStruct{}
		olaMap.HttpService = mocking
		_, err := olaMap.PlaceAutoComplete("mock-input")
		assert.Nil(t, nil, err)
		if mocking.StatusCode != 200 {
			t.Error("expected : ", 200, "got : ", mocking.StatusCode)
		}
		assert.Equal(t, AutoCompleteResponse, mocking.MockBody)
	})
}

func TestGecode(t *testing.T) {
	t.Run("Invalid address", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		_, err := olaMap.GeoCode("", "", "")
		expectedErr := fmt.Errorf("Missing required query parameters: 'address'")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		_, err := olaMap.GeoCode("mock-address", "mock-bounds", "mock-language")
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("success", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.Token = "mockToken"
		mocking := &MockStruct{}
		olaMap.HttpService = mocking
		_, err := olaMap.GeoCode("mock-address", "mock-bounds", "mock-language")
		assert.Nil(t, nil, err)
		if mocking.StatusCode != 200 {
			t.Error("expected : ", 200, "got : ", mocking.StatusCode)
		}
		assert.Equal(t, GeoCodeResponse, mocking.MockBody)

	})
}

func TestReverseGeocode(t *testing.T) {
	t.Run("Invalid latlng", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		_, err := olaMap.ReverseGeocode("")
		expectedErr := fmt.Errorf("Missing required query parameters: 'latlng'")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		_, err := olaMap.ReverseGeocode("mock-latlng")
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("success", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.Token = "mockToken"
		mocking := &MockStruct{}
		olaMap.HttpService = mocking
		_, err := olaMap.ReverseGeocode("mock-latlng")
		assert.Nil(t, nil, err)
		if mocking.StatusCode != 200 {
			t.Error("expected : ", 200, "got : ", mocking.StatusCode)
		}
		assert.Equal(t, ReverseGeocodeResponse, mocking.MockBody)

	})
}

func TestGetDistanceMatrix(t *testing.T) {
	t.Run("Invalid origins & destinations", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		_, err := olaMap.GetDistanceMatrix("", "")
		expectedErr := fmt.Errorf("Missing required query parameters: 'origin' and/or 'destination'")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		_, err := olaMap.GetDistanceMatrix("mock-origins", "mock-destinations")
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("success", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.Token = "mockToken"
		mocking := &MockStruct{}
		olaMap.HttpService = mocking
		_, err := olaMap.GetDistanceMatrix("mock-origins", "mock-destinations")
		assert.Nil(t, nil, err)
		if mocking.StatusCode != 200 {
			t.Error("expected : ", 200, "got : ", mocking.StatusCode)
		}
		assert.Equal(t, DistanceMatrixResponse, mocking.MockBody)
	})
}

func TestArrayOfData(t *testing.T) {
	t.Run("Invalid datasetname", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		_, err := olaMap.ArrayOfData("")
		expectedErr := fmt.Errorf("Missing required query parameters: 'datasetname'")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		_, err := olaMap.ArrayOfData("mock-datasetname")
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("success", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.Token = "mockToken"
		mocking := &MockStruct{}
		olaMap.HttpService = mocking
		_, err := olaMap.ArrayOfData("mock-datasetname")
		assert.Nil(t, nil, err)
		if mocking.StatusCode != 200 {
			t.Error("expected : ", 200, "got : ", mocking.StatusCode)
		}
		assert.Equal(t, ArrayOfDataResponse, mocking.MockBody)

	})
}

func TestGetStyleDetails(t *testing.T) {
	t.Run("Invalid stylename", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		_, err := olaMap.GetStyleDetails("")
		expectedErr := fmt.Errorf("Missing required query parameters: 'stylename'")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		_, err := olaMap.GetStyleDetails("mock-stylename")
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("success", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.Token = "mockToken"
		mocking := &MockStruct{}
		olaMap.HttpService = mocking
		_, err := olaMap.GetStyleDetails("mock-stylename")
		assert.Nil(t, nil, err)
		if mocking.StatusCode != 200 {
			t.Error("expected : ", 200, "got : ", mocking.StatusCode)
		}
		assert.Equal(t, StyleDetailResponse, mocking.MockBody)

	})
}

func TestGetMapStyle(t *testing.T) {
	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		_, err := olaMap.GetMapStyle()
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)

	})

	t.Run("success", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.Token = "mockToken"
		mocking := &MockStruct{}
		olaMap.HttpService = mocking
		_, err := olaMap.GetMapStyle()
		assert.Nil(t, nil, err)
		if mocking.StatusCode != 200 {
			t.Error("expected : ", 200, "got : ", mocking.StatusCode)
		}
		assert.Equal(t, MapStyleResponse, mocking.MockBody)
	})
}

func TestGetPlaceDetail(t *testing.T) {
	t.Run("Invalid 'placeid'", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		_, err := olaMap.GetPlaceDetail("")
		expectedErr := fmt.Errorf("Missing required query parameters: 'placeid'")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		_, err := olaMap.GetPlaceDetail("mock-placeid")
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("success", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.Token = "mockToken"
		mocking := &MockStruct{}
		olaMap.HttpService = mocking
		_, err := olaMap.GetPlaceDetail("mock-placeid")
		assert.Nil(t, nil, err)
		if mocking.StatusCode != 200 {
			t.Error("expected : ", 200, "got : ", mocking.StatusCode)
		}
		assert.Equal(t, PlaceDetailResponse, mocking.MockBody)

	})
}

func TestGetNearBySearch(t *testing.T) {
	t.Run("Invalid 'layers' and 'location'", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		_, err := olaMap.GetNearBySearch(NearBySearch{Layers: "", Location: ""})
		expectedErr := fmt.Errorf("Missing required query parameters: 'layers' and/or 'location'")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		fmt.Printf("Mocking = %+v", olaMap)
		_, err := olaMap.GetNearBySearch(NearBySearch{Layers: "Mock-Layers", Location: "Mock-Location"})
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("success", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.Token = "mockToken"
		mocking := &MockStruct{}
		olaMap.HttpService = mocking
		_, err := olaMap.GetNearBySearch(NearBySearch{Layers: "Mock-Layers", Location: "Mock-Location"})
		assert.Nil(t, nil, err)
		if mocking.StatusCode != 200 {
			t.Error("expected : ", 200, "got : ", mocking.StatusCode)
		}
		assert.Equal(t, NearBySearchResponses, mocking.MockBody)
	})
}

func TestGetTextSearch(t *testing.T) {
	t.Run("Invalid input", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		_, err := olaMap.GetTextSearch(TextSearch{Input: ""})
		expectedErr := fmt.Errorf("Missing required query parameters: 'input'")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		_, err := olaMap.GetTextSearch(TextSearch{Input: "mock-input"})
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("success", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.Token = "mockToken"
		mocking := &MockStruct{}
		olaMap.HttpService = mocking
		_, err := olaMap.GetTextSearch(TextSearch{Input: "mock-input"})
		assert.Nil(t, nil, err)
		if mocking.StatusCode != 200 {
			t.Error("expected : ", 200, "got : ", mocking.StatusCode)
		}
		assert.Equal(t, TextSearchResponse, mocking.MockBody)
	})
}

func TestSnapToRoad(t *testing.T) {
	t.Run("Invalid points", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		_, err := olaMap.GetSnapToRoad("", "")
		expectedErr := fmt.Errorf("Missing required query parameters: 'points'")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		_, err := olaMap.GetSnapToRoad("mock-points", "mock-enhancepath")
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("success", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.Token = "mockToken"
		mocking := &MockStruct{}
		olaMap.HttpService = mocking
		_, err := olaMap.GetSnapToRoad("mock-points", "mock-enhancepath")
		assert.Nil(t, nil, err)
		if mocking.StatusCode != 200 {
			t.Error("expected : ", 200, "got : ", mocking.StatusCode)
		}
		assert.Equal(t, SnapToRoadResponse, mocking.MockBody)

	})
}

func TestNearestRoad(t *testing.T) {
	t.Run("Invalid points or radius", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		_, err := olaMap.GetNearestRoads("", "")
		expectedErr := fmt.Errorf("Missing required query parameters: 'points' and/or 'radius'")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		_, err := olaMap.GetNearestRoads("mock-points", "mock-radius")
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("success", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.Token = "mockToken"
		mocking := &MockStruct{}
		olaMap.HttpService = mocking
		_, err := olaMap.GetNearestRoads("mock-points", "mock-radius")
		assert.Nil(t, nil, err)
		if mocking.StatusCode != 200 {
			t.Error("expected : ", 200, "got : ", mocking.StatusCode)
		}
		assert.Equal(t, NearestRoadResponse, mocking.MockBody)

	})
}

func TestGetImageCenter(t *testing.T) {
	t.Run("Invalid lattitude & longtitude", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		mapImageCenter := MapImageCenter{
			Stylename:   "mock-style",
			Zoomlevel:   "22.0",
			Imagewidth:  "78",
			Imageformat: ".png",
			Imageheight: "57.8",
		}

		_, err := olaMap.GetStaticMapImageCenter(mapImageCenter)
		expectedErr := fmt.Errorf("Missing required query parameters: 'stylename' or 'longitude' or 'latitude' or 'zoomlevel' or 'width' or 'height' or 'format'")
		assert.Exactly(t, err, expectedErr)

	})

	t.Run("Error in zoomlevel conversion", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		mapImageCenter := MapImageCenter{
			Stylename:   "mock-style",
			Zoomlevel:   "mock-zoom",
			Imagewidth:  "78",
			Imageformat: ".png",
			Imageheight: "57.8",
			Latitude:    "78.0",
			Longitude:   "67.8",
		}

		_, err := olaMap.GetStaticMapImageCenter(mapImageCenter)
		expectedErr := fmt.Errorf("Invalid zoom level value")
		assert.Exactly(t, err, expectedErr)

	})

	t.Run("Error in Image width conversion", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		mapImageCenter := MapImageCenter{
			Stylename:   "mock-style",
			Zoomlevel:   "69",
			Imagewidth:  "mock-image",
			Imageformat: ".png",
			Imageheight: "80",
			Latitude:    "78.0",
			Longitude:   "67.8",
		}

		_, err := olaMap.GetStaticMapImageCenter(mapImageCenter)
		expectedErr := fmt.Errorf("Invalid image width value")
		assert.Exactly(t, err, expectedErr)

	})

	t.Run("Error in Image height conversion", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		mapImageCenter := MapImageCenter{
			Stylename:   "mock-style",
			Zoomlevel:   "69",
			Imagewidth:  "90",
			Imageformat: ".png",
			Imageheight: "mockImage",
			Latitude:    "78.0",
			Longitude:   "67.8",
		}

		_, err := olaMap.GetStaticMapImageCenter(mapImageCenter)
		expectedErr := fmt.Errorf("Invalid image height value")
		assert.Exactly(t, err, expectedErr)

	})

	t.Run("InValid Token", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		mapImageCenter := MapImageCenter{
			Stylename:   "mock-style",
			Zoomlevel:   "69",
			Imagewidth:  "90",
			Imageformat: ".png",
			Imageheight: "100",
			Latitude:    "78.0",
			Longitude:   "67.8",
		}

		_, err := olaMap.GetStaticMapImageCenter(mapImageCenter)
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)

	})
}

func TestGetImageBounded(t *testing.T) {
	t.Run("Invalid stylename", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		mapImageBounded := MapImageBounded{
			Stylename:   "",
			Minxstr:     "",
			Minystr:     "",
			Maxxstr:     "",
			Maxystr:     "",
			Imagewidth:  "",
			Imageformat: ".png",
			Imageheight: "",
		}

		_, err := olaMap.GetStaticMapImageBounded(mapImageBounded)
		expectedErr := fmt.Errorf("Missing required query parameters: 'styleName' or 'minXStr' or 'minYStr' or 'maxXStr' or 'maxYStr' or 'imageWidthStr' or 'imageHeightStr' or 'imageFormat'")
		assert.Exactly(t, err, expectedErr)

	})

	t.Run("Invalid min_x value", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		mapImageBounded := MapImageBounded{
			Stylename:   "mock-style",
			Minxstr:     "mock-minx",
			Minystr:     "12.8",
			Maxxstr:     "21.7",
			Maxystr:     "22.4",
			Imagewidth:  "78",
			Imageformat: ".png",
			Imageheight: "57",
		}

		_, err := olaMap.GetStaticMapImageBounded(mapImageBounded)
		expectedErr := fmt.Errorf("Invalid min_x value")
		assert.Exactly(t, err, expectedErr)

	})

	t.Run("Invalid min_y value", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		mapImageBounded := MapImageBounded{
			Stylename:   "mock-style",
			Minxstr:     "12.1",
			Minystr:     "mock-miny",
			Maxxstr:     "21.7",
			Maxystr:     "22.4",
			Imagewidth:  "78",
			Imageformat: ".png",
			Imageheight: "57",
		}

		_, err := olaMap.GetStaticMapImageBounded(mapImageBounded)
		expectedErr := fmt.Errorf("Invalid min_y value")
		assert.Exactly(t, err, expectedErr)

	})

	t.Run("Invalid max_x value", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		mapImageBounded := MapImageBounded{
			Stylename:   "mock-style",
			Minxstr:     "12.1",
			Minystr:     "33.2",
			Maxxstr:     "mock-maxx",
			Maxystr:     "22.4",
			Imagewidth:  "78",
			Imageformat: ".png",
			Imageheight: "57",
		}

		_, err := olaMap.GetStaticMapImageBounded(mapImageBounded)
		expectedErr := fmt.Errorf("Invalid max_x value")
		assert.Exactly(t, err, expectedErr)

	})

	t.Run("Invalid max_y value", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		mapImageBounded := MapImageBounded{
			Stylename:   "mock-style",
			Minxstr:     "12.1",
			Minystr:     "33.2",
			Maxxstr:     "8.5",
			Maxystr:     "mock-maxy",
			Imagewidth:  "78",
			Imageformat: ".png",
			Imageheight: "57",
		}

		_, err := olaMap.GetStaticMapImageBounded(mapImageBounded)
		expectedErr := fmt.Errorf("Invalid max_y value")
		assert.Exactly(t, err, expectedErr)

	})

	t.Run("Invalid image height value", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		mapImageBounded := MapImageBounded{
			Stylename:   "mock-style",
			Minxstr:     "12.1",
			Minystr:     "33.2",
			Maxxstr:     "8.5",
			Maxystr:     "33.2",
			Imagewidth:  "23",
			Imageformat: ".png",
			Imageheight: "mock-imageheight",
		}

		_, err := olaMap.GetStaticMapImageBounded(mapImageBounded)
		expectedErr := fmt.Errorf("Invalid image height value")
		assert.Exactly(t, err, expectedErr)

	})

	t.Run("Invalid image width value", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		mapImageBounded := MapImageBounded{
			Stylename:   "mock-style",
			Minxstr:     "12.1",
			Minystr:     "33.2",
			Maxxstr:     "8.5",
			Maxystr:     "33.2",
			Imagewidth:  "mock-imagewidth",
			Imageformat: ".png",
			Imageheight: "57",
		}

		_, err := olaMap.GetStaticMapImageBounded(mapImageBounded)
		expectedErr := fmt.Errorf("Invalid image width value")
		assert.Exactly(t, err, expectedErr)

	})

	t.Run("Invalid OAuth token", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		mapImageBounded := MapImageBounded{
			Stylename:   "mock-style",
			Minxstr:     "12.1",
			Minystr:     "33.2",
			Maxxstr:     "8.5",
			Maxystr:     "33.2",
			Imagewidth:  "23",
			Imageformat: ".png",
			Imageheight: "71",
		}

		_, err := olaMap.GetStaticMapImageBounded(mapImageBounded)
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)
	})
}

func TestStaticMapImage(t *testing.T) {
	t.Run("Invalid style", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		mapImage := MapImage{
			Stylename:   "",
			Imagewidth:  "78",
			Imageformat: ".png",
			Imageheight: "57",
		}

		_, err := olaMap.StaticMapImage(mapImage)
		expectedErr := fmt.Errorf("Missing required query parameters: 'stylename' or 'imagewidth' or 'imageheight' or 'imageformat' or path")
		assert.Exactly(t, err, expectedErr)

	})

	t.Run("Image-width conversion error", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		mapImage := MapImage{
			Stylename:   "mock-style",
			Imagewidth:  "mock-width",
			Imageformat: ".png",
			Imageheight: "57",
			Path:        "MockPath",
			Markers:     []string{"test-1", "test-2"},
		}

		_, err := olaMap.StaticMapImage(mapImage)
		expectedErr := fmt.Errorf("Invalid image width value")
		assert.Exactly(t, err, expectedErr)

	})

	t.Run("Image-height conversion error", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		mapImage := MapImage{
			Stylename:   "mock-style",
			Imagewidth:  "79",
			Imageformat: ".png",
			Imageheight: "mock-height",
			Path:        "MockPath",
			Markers:     []string{"test-1", "test-2"},
		}

		_, err := olaMap.StaticMapImage(mapImage)
		expectedErr := fmt.Errorf("Invalid image height value")
		assert.Exactly(t, err, expectedErr)

	})
	t.Run("InValid Token", func(t *testing.T) {
		olaMap := &OLAMap{}
		olaMap.HttpService = &MockStruct{}
		mapImage := MapImage{
			Stylename:   "Mock-Style",
			Imagewidth:  "90",
			Imageformat: ".png",
			Imageheight: "100",
			Path:        "MockPath",
			Markers:     []string{"test-1", "test-2"},
		}

		_, err := olaMap.StaticMapImage(mapImage)
		expectedErr := fmt.Errorf("Invalid OAuth token")
		assert.Exactly(t, err, expectedErr)

	})
}
