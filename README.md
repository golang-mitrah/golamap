# Go OLA Map Wrapper

This Go package provides a wrapper for the OLA Map API, allowing you to interact with various map-related functionalities such as getting directions, geocoding, and more. 

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Available Methods](#available-methods)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Features
- Easy-to-use Golang wrapper for OLA Map.
- Provides methods to interact with OLA Map and its data structures.
- Supports searching and querying geographical data using OLA codes.
- Lightweight and performant implementation suitable for production use.
- Supports searching and querying geographical data using OLA codes.
- Customizable: Add additional utilities to enhance functionality.

## Installation

1. To install the OLA Map wrapper, use the following command:
   ```bash
   go get github.com/golang-mitrah/golamap
   ```

## Usage

Below is an example of how to use the OLA Map Wrapper in your project:

```go
package main

import (
    "fmt"
    "github.com/golang-mitrah/golamap"
)

func main() {
    // Initialize the OLA Map with a request ID
    olaMap := golamap.Initialize("your-unique-request-id")

    // Configure access token
    err := olaMap.ConfigureAccessToken("your-client-id", "your-client-secret")
    if err != nil {
        fmt.Println("Error configuring access token:", err)
        return
    }

    // Get directions
    directions, err := olaMap.GetDirections("origin", "destination")
    if err != nil {
        fmt.Println("Error getting directions:", err)
        return
    }
    fmt.Println("Directions:", directions)
}
```

## Available Methods

The following methods are available for use with the `OLAMap` struct:

- **`Initialize(requestID string) *OLAMap`**: Initializes a new OLA Map instance with a unique request ID.
- **`ConfigureAccessToken(clientID, clientSecret string) error`**: Configures the OLA access token using client credentials.
- **`GetDirections(origin, destination string) (Directions, error)`**: Retrieves directions from the origin to the destination.
- **`PlaceAutoComplete(input string) (Places, error)`**: Provides place suggestions based on the input.
- **`GeoCode(address, bounds, language string) (GeoData, error)`**: Converts an address into geographic coordinates.
- **`ReverseGeocode(latlng string) (Address, error)`**: Converts geographic coordinates back into an address.
- **`GetDistanceMatrix(origins, destinations string) (DistanceMatrix, error)`**: Calculates distances between multiple origins and destinations.
- **`ArrayOfData(datasetName string)`**: Retrieves an array of data associated with the specified dataset name.
- **`GetStyleDetails(styleName string)`**: Fetches details about a specific style using the provided style name.
- **`GetMapStyle()`**: Retrieves the current map style being used.
- **`GetPlaceDetail(placeID string)`**: Fetches detailed information about a specific place using its unique identifier.
- **`GetNearBySearch(nearBySearch NearBySearch)`**: Conducts a nearby search based on the provided parameters in NearBySearch.
- **`GetTextSearch(textSearch TextSearch)`**: Executes a text-based search using the specified criteria in TextSearch.
- **`GetSnapToRoad(points, enhancePath string)`**: Snaps the provided GPS points to the nearest roads, enhancing the path as specified.
- **`GetNearestRoads(points string, radius string)`**: Retrieves the nearest roads to the specified GPS points within the given radius.
- **`GetStaticMapImageCenter(mapImageCenter MapImageCenter)`**: Generates a static map image centered around the specified coordinates.
- **`GetStaticMapImageBounded(mapImageBounded MapImageBounded)`**: Generates a static map image within specified bounding coordinates.
- **`StaticMapImage(mapImage MapImage)`**: Fetches a static map image based on the provided map image parameters.

## Testing

Test cases have been written for all files to ensure functionality and reliability. To run the tests, use the following command:

```bash
go test ./...
```

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/new-feature`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/new-feature`).
5. Open a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE]