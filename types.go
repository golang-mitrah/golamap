package golamap

type NearBySearch struct {
	Layers       string
	Location     string
	Types        string
	Radius       string
	Strictbounds string
	WithCentroid string
	Limit        string
}

type MapImage struct {
	Stylename   string
	Imagewidth  string
	Imageheight string
	Imageformat string
	Path        string
	Markers     []string
}

type MapImageBounded struct {
	Stylename   string
	Minxstr     string
	Minystr     string
	Maxxstr     string
	Maxystr     string
	Imagewidth  string
	Imageheight string
	Imageformat string
	Markers     []string
	Path        string
}

type MapImageCenter struct {
	Stylename   string
	Longitude   string
	Latitude    string
	Zoomlevel   string
	Imagewidth  string
	Imageheight string
	Imageformat string
	Markers     []string
	Path        string
}

type TextSearch struct {
	Input    string
	Location string
	Radius   string
	Types    string
	Size     string
}

type Directions struct {
	GeocodedWaypoints []GeocodedWaypoint `json:"geocoded_waypoints"`
	Routes            []Route            `json:"routes"`
	Status            string             `json:"status"`
	SourceFrom        string             `json:"source_from"`
}

type GeocodedWaypoint struct {
	GeocoderStatus string   `json:"geocoder_status"`
	PlaceID        string   `json:"place_id"`
	Types          []string `json:"types"`
}

// Step represents a single step in the route.
type Step struct {
	Distance         int      `json:"distance"`
	ReadableDistance string   `json:"readable_distance"`
	Duration         int      `json:"duration"`
	ReadableDuration string   `json:"readable_duration"`
	StartLocation    Location `json:"start_location"`
	EndLocation      Location `json:"end_location"`
	Instructions     string   `json:"instructions"`
	Maneuver         string   `json:"maneuver"`
	BearingBefore    int      `json:"bearing_before"`
	BearingAfter     int      `json:"bearing_after"`
}

// Location represents geographical coordinates.
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Leg struct {
	Distance         int      `json:"distance"`
	ReadableDistance string   `json:"readable_distance"`
	Duration         int      `json:"duration"`
	ReadableDuration string   `json:"readable_duration"`
	EndAddress       string   `json:"end_address"`
	EndLocation      Location `json:"end_location"`
	StartAddress     string   `json:"start_address"`
	StartLocation    Location `json:"start_location"`
	Steps            []Step   `json:"steps"`
}

type Route struct {
	Bounds           Bounds   `json:"bounds"`
	Copyrights       string   `json:"copyrights"`
	Legs             []Leg    `json:"legs"`
	OverviewPolyline string   `json:"overview_polyline"`
	TravelAdvisory   string   `json:"travel_advisory"`
	Summary          string   `json:"summary"`
	Warnings         []string `json:"warnings"`
	WaypointOrder    []int    `json:"waypoint_order"`
}

// Bounds represents the geographical bounds of the route.
type Bounds struct {
	Southwest Location `json:"southwest"`
	Northeast Location `json:"northeast"`
}

type DistanceMatrix struct {
	Rows   []Row  `json:"rows"`
	Status string `json:"status"`
}

type Row struct {
	Elements []Element `json:"elements"`
}

type Element struct {
	Duration int    `json:"duration"`
	Distance int    `json:"distance"`
	Polyline string `json:"polyline"`
	Status   string `json:"status"`
}

type SnapToRoad struct {
	Status        string         `json:"status"`
	SnappedPoints []SnappedPoint `json:"snapped_points"`
}

type SnappedPoint struct {
	Location      Location `json:"location"`
	OriginalIndex int      `json:"original_index"`
	SnappedType   string   `json:"snapped_type"`
}

type NearestRoad struct {
	Status  string   `json:"status"`
	Results []Result `json:"results"`
}

type Result struct {
	Lat           float64 `json:"lat"`
	Lng           float64 `json:"lng"`
	Distance      float64 `json:"distance"`
	OriginalIndex int     `json:"originalIndex"`
	Status        string  `json:"status"`
}

type AutoComplete struct {
	Predictions  []Prediction `json:"predictions"`
	InfoMessages []string     `json:"info_messages"`
	ErrorMessage string       `json:"error_message"`
	Status       string       `json:"status"`
}

type Prediction struct {
	Reference            string                           `json:"reference"`
	Types                []string                         `json:"types"`
	MatchedSubstrings    []MatchedSubstring               `json:"matched_substrings"`
	DistanceMeters       int                              `json:"distance_meters"`
	Terms                []Term                           `json:"terms"`
	StructuredFormatting AutoCompleteStructuredFormatting `json:"structured_formatting"`
	Description          string                           `json:"description"`
	Geometry             Geometry                         `json:"geometry"`
	PlaceID              string                           `json:"place_id"`
}

type MatchedSubstring struct {
	Offset int `json:"offset"`
	Length int `json:"length"`
}

type Term struct {
	Offset int    `json:"offset"`
	Value  string `json:"value"`
}

type AutoCompleteStructuredFormatting struct {
	MainText                  string             `json:"main_text"`
	SecondaryText             string             `json:"secondary_text"`
	MainTextMatchedSubstrings []MatchedSubstring `json:"main_text_matched_substrings"`
}

type Geometry struct {
	Location Location `json:"location"`
}

type PlaceDetail struct {
	HtmlAttributions  []interface{}     `json:"html_attributions"`
	Resultplacedetail Resultplacedetail `json:"result"`
	InfoMessages      []string          `json:"info_messages"`
	ErrorMessage      string            `json:"error_message"`
	Status            string            `json:"status"`
}

type Resultplacedetail struct {
	AddressComponents        []AddressComponent  `json:"address_components"`
	FormattedAddress         string              `json:"formatted_address"`
	Geometryplacedetail      Geometryplacedetail `json:"geometry"`
	PlaceID                  string              `json:"place_id"`
	Reference                string              `json:"reference"`
	BusinessStatus           string              `json:"business_status"`
	FormattedPhoneNumber     string              `json:"formatted_phone_number"`
	Icon                     string              `json:"icon"`
	IconBackgroundColor      string              `json:"icon_background_color"`
	IconMaskBaseURI          string              `json:"icon_mask_base_uri"`
	InternationalPhoneNumber string              `json:"international_phone_number"`
	Name                     string              `json:"name"`
	OpeningHours             OpeningHours        `json:"opening_hours"`
	PlusCode                 PlusCode            `json:"plus_code"`
	Rating                   float64             `json:"rating"`
	Reviews                  []Review            `json:"reviews"`
	Types                    []string            `json:"types"`
	Layer                    []string            `json:"layer"`
	URL                      string              `json:"url"`
	UserRatingsTotal         int                 `json:"user_ratings_total"`
	UTCOffset                int                 `json:"utc_offset"`
	Vicinity                 string              `json:"vicinity"`
	Website                  string              `json:"website"`
	PriceLevel               string              `json:"price_level"`
	Photos                   []interface{}       `json:"photos"`
	AdrAddress               string              `json:"adr_address"`
}

type AddressComponent struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

type Geometryplacedetail struct {
	Location Location `json:"location"`
	Viewport Viewport `json:"viewport"`
}

type Viewport struct {
	Northeast Location `json:"northeast"`
	Southwest Location `json:"southwest"`
}

type OpeningHours struct {
	OpenNow     bool            `json:"open_now"`
	Periods     []OpeningPeriod `json:"periods"`
	WeekdayText []string        `json:"weekday_text"`
}

type OpeningPeriod struct {
	Close OpeningTime `json:"close"`
	Open  OpeningTime `json:"open"`
}

type OpeningTime struct {
	Day  int    `json:"day"`
	Time string `json:"time"`
}

type PlusCode struct {
	CompoundCode string `json:"compound_code"`
	GlobalCode   string `json:"global_code"`
}

type Review struct {
	AuthorName              string  `json:"author_name"`
	AuthorURL               string  `json:"author_url"`
	Language                string  `json:"language"`
	ProfilePhotoURL         string  `json:"profile_photo_url"`
	Rating                  float64 `json:"rating"`
	RelativeTimeDescription string  `json:"relative_time_description"`
	Text                    string  `json:"text"`
	Time                    int64   `json:"time"`
}

type NearBySearchResponse struct {
	Predictions  []PredictionNearbySearch `json:"predictions"`
	InfoMessages []string                 `json:"info_messages"`
	ErrorMessage string                   `json:"error_message"`
	Status       string                   `json:"status"`
}

type PredictionNearbySearch struct {
	Description          string               `json:"description"`
	MatchedSubstrings    []MatchedSubstring   `json:"matched_substrings"`
	PlaceID              string               `json:"place_id"`
	Reference            string               `json:"reference"`
	StructuredFormatting StructuredFormatting `json:"structured_formatting"`
	Terms                []Term               `json:"terms"`
	Types                []string             `json:"types"`
	Layer                []string             `json:"layer"`
	DistanceMeters       int                  `json:"distance_meters"`
}

type StructuredFormatting struct {
	MainText                       string             `json:"main_text"`
	MainTextMatchedSubstrings      []MatchedSubstring `json:"main_text_matched_substrings"`
	SecondaryText                  string             `json:"secondary_text"`
	SecondaryTextMatchedSubstrings []MatchedSubstring `json:"secondary_text_matched_substrings"`
}

type TextBySearch struct {
	Predictions  []PredictionTextBySearch `json:"predictions"`
	InfoMessages []string                 `json:"info_messages"`
	ErrorMessage string                   `json:"error_message"`
	Status       string                   `json:"status"`
}

type PredictionTextBySearch struct {
	FormattedAddress string   `json:"formatted_address"`
	Geometry         Geometry `json:"geometry"`
	PlaceID          string   `json:"place_id"`
	Name             string   `json:"name"`
	Types            []string `json:"types"`
}

type ForwardGecode struct {
	Status           string            `json:"status"`
	GeocodingResults []GeocodingResult `json:"geocodingResults"`
}

type GeocodingResult struct {
	FormattedAddress  string             `json:"formatted_address"`
	Types             []string           `json:"types"`
	Name              string             `json:"name"`
	Geometry          GecodeGeometry     `json:"geometry"`
	AddressComponents []AddressComponent `json:"address_components"`
	PlusCode          PlusCode           `json:"plus_code"`
	PlaceID           string             `json:"place_id"`
	Layer             []string           `json:"layer"`
}

type GecodeGeometry struct {
	Location     Location `json:"location"`
	Viewport     Viewport `json:"viewport"`
	LocationType string   `json:"location_type"`
}

type ReverseGecode struct {
	ErrorMessage string             `json:"error_message"`
	InfoMessages []string           `json:"info_messages"`
	Results      []ReverseGeoResult `json:"results"`
	PlusCode     PlusCode           `json:"plus_code"`
	Status       string             `json:"status"`
}

type ReverseGeoResult struct {
	FormattedAddress  string             `json:"formatted_address"`
	Types             []string           `json:"types"`
	Name              string             `json:"name"`
	Geometry          ReverseGeometry    `json:"geometry"`
	AddressComponents []AddressComponent `json:"address_components"`
	PlusCode          PlusCode           `json:"plus_code"`
	PlaceID           string             `json:"place_id"`
	Layer             []string           `json:"layer"`
}

type ReverseGeometry struct {
	Location     Location `json:"location"`
	Viewport     Viewport `json:"viewport"`
	LocationType string   `json:"location_type"`
}

type ArrayOfData struct {
	Attribution     string        `json:"attribution"`
	Basename        string        `json:"basename"`
	Bounds          []float64     `json:"bounds"`
	Center          []float64     `json:"center"`
	Compression     string        `json:"compression"`
	DataVersion     string        `json:"data_version"`
	Description     string        `json:"description"`
	Format          string        `json:"format"`
	ID              string        `json:"id"`
	MaxZoom         int           `json:"maxzoom"`
	MinZoom         int           `json:"minzoom"`
	Name            string        `json:"name"`
	Tiles           []string      `json:"tiles"`
	Type            string        `json:"type"`
	VectorLayers    []VectorLayer `json:"vector_layers"`
	TileJSONVersion string        `json:"tilejson"`
}

type VectorLayer struct {
	ID      string    `json:"id"`
	MaxZoom int       `json:"maxzoom"`
	MinZoom int       `json:"minzoom"`
	Fields  FieldInfo `json:"fields"`
}

type FieldInfo struct {
	Name      string `json:"name"`
	Status    string `json:"status"`
	Ref       string `json:"ref"`
	Class     string `json:"class"`
	SubClass  string `json:"subclass"`
	Elevation string `json:"ele"`
}

type VectorMapStyle struct {
	Version int    `json:"version"`
	Name    string `json:"name"`
	ID      string `json:"id"`
	URL     string `json:"url"`
}

type VectorStyleDetails struct {
	ID      string            `json:"id"`
	Version int               `json:"version"`
	Name    string            `json:"name"`
	Sources map[string]Source `json:"sources"`
	Glyphs  string            `json:"glyphs"`
	Layers  []Layer           `json:"layers"`
}

type Layer struct {
	ID          string      `json:"id"`
	Type        string      `json:"type"`
	Source      string      `json:"source"`
	SourceLayer string      `json:"source-layer"`
	Filter      interface{} `json:"filter"` // Using interface{} to handle nested structures
	Layout      Layout      `json:"layout"`
	Paint       Paint       `json:"paint"`
}

type Layout struct {
	Visibility string `json:"visibility"`
}

type Paint struct {
	FillAntialias bool   `json:"fill-antialias"`
	FillColor     string `json:"fill-color"`
}

type Source struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}
