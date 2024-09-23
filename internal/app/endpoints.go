package app

var (
	TokenURL                 = "https://account.olamaps.io/realms/olamaps/protocol/openid-connect/token"
	DirectionsURL            = "https://api.olamaps.io/routing/v1/directions?origin=%s&destination=%s"
	PlaceAutoCompleteURL     = "https://api.olamaps.io/places/v1/autocomplete?input=%s"
	GeoCodeURL               = "https://api.olamaps.io/places/v1/geocode?address=%s&bounds=%s&language=%s"
	ReverseGeocodeURL        = "https://api.olamaps.io/places/v1/reverse-geocode?latlng=%s"
	DistanceMatrixURL        = "https://api.olamaps.io/routing/v1/distanceMatrix?origins=%s&destinations=%s"
	ArrayOfDataURL           = "https://api.olamaps.io/tiles/vector/v1/data/%s.json"
	StyleDetailsURL          = "https://api.olamaps.io/tiles/vector/v1/styles/%s/style.json"
	MapStyleURL              = "https://api.olamaps.io/tiles/vector/v1/styles.json"
	PlaceDetailURL           = "https://api.olamaps.io/places/v1/details?place_id=%v"
	NearBySearchURL          = "https://api.olamaps.io/places/v1/nearbysearch?layers=%s&location=%s&types=%s&radius=%s&strictbounds=%s&withCentroid=%s&limit=%s"
	TextSearchURL            = "https://api.olamaps.io/places/v1/textsearch?input=%s&location=%s&radius=%s&types=%s&size=%s"
	SnapToRoadURL            = "https://api.olamaps.io/routing/v1/snapToRoad?%s"
	NearestRoadsURL          = "https://api.olamaps.io/routing/v1/nearestRoads?points=%s&radius=%s"
	StaticMapImageCenterURL  = "https://api.olamaps.io/tiles/v1/styles/%s/static/%f,%f,%d/%dx%d.%s"
	StaticMapImageBoundedURL = "https://api.olamaps.io/tiles/v1/styles/%s/static/%f,%f,%f,%f/%dx%d.%s"
	StaticMapImageURL        = "https://api.olamaps.io/tiles/v1/styles/%s/static/auto/%dx%d.%s"
)
