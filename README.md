# Ola Maps Go API Wrapper library

This Go package provides handlers for interacting with the Olamaps API. The handlers are designed to support various API endpoints related to token retrieval, directions, place searches, geocoding,road search and Maptiles.

## Table of Contents
- [Installation](#installation)
- [Setup](#setup)
- [Api Endpoints](#api-endpoints)
- [License](#license)
- [Handling Errors](#handling-errors)

## 1.Installation
- To use this package, you'll need Go installed on your machine.

## 2.Setup
### Clone the repository
 ```git clone https://github.com/your-repo/ola-maps.git```

### Navigate to the project repository
  ```cd repository```

### Install dependencies
- To run the dependencies
   ```go mod tidy```

### Set up API credentials
- To configure the `app.clientID` , `app.clientSecret` and `app.tokenURL` in the internal package.

### Run
- ```go run main.go```

## 3.API endpoints:
### Get Token `/api/token`
   - Retrieves the access token required for making API requests.
   - It gets a parameter of `client id` , `client secret` , `token url` and it returns access token.

### Get PBF File `/pbfFile`
   - Retrieves protocol buffer (PBF) file data.

### Get Directions `/direction`
   - Provides directions between an origin and a destination.

### Get Distance `/routing/distanceMatrix`
   - Provides the distance between multiple origins and destinations.

### Get Geocode `/geocoding`
   - Converts an address into latitude and longitude.

### Get Reverse Geocode `/reverseGeocoding`
   - Converts latitude and longitude into a human-readable address.

### Get Array Of Data `/mapTiles/ArrayData`
   - Retrieves array data for a specific dataset.

### Get Style Details `/mapTiles/styleDetails`
   - Retrieves details for a specific map style.

### Get Map Style `/map/style`
   - Lists available map styles.

### Place Auto Complete `/autocomplete`
   - Provides autocomplete suggestions for place names based on user input.

### Get Place detail `/place/details`
   - Retrieves details for a specific place.

### Get Nearby Search `/nearby/search`
   - Searches for places near a specified location.

### Get Text Search `/text/search`
   - Searches for places based on a text query.

### Get Snap to Road `/snap/toroad`
   - Snaps a point to the nearest road.

### Get Nearest to Road `/nearest/road`
   - Finds the nearest road to specified points.

### Get a map image center `/tiles/center`
   - Center Based where latitude, longitude and zoom are provided.

### Get a map image bounded `/tiles/bounded`
   - Area based where bounding box coordinates are provided.

### Get a map image `/tiles/mapimage`
   - Auto Based where the the coordinates are calculated based on the path/polyline which is provided.

## License

- This project is licensed under the MIT License - see the [click here](https://maps.olakrutrim.com/legal-docs/terms-and-conditions.pdf)


## Handling Errors
Error codes are followed
| Status Code | Error Type |
| ----------- | ---------- |
|    400      |  Bad Request  |
|  401 | Unauthorized |
|  403 | Forbidden    |
|  404 | Not Found    |
|  409 | Conflict     |
|  422 | Unprocessable entity |
|  429 | too many request |
|  500 | Internal server error |
