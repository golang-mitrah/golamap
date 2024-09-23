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
