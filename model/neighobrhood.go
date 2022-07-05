package model

type Neighborhood struct {
	Name     string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`
	Area     string `json:"area,omitempty"`
}

type NeighborhoodSearchHit struct {
	Name    string `json:"name,omitempty"`
	Content string `json:"content,omitempty"`
}

type NeighborhoodCounterHit struct {
	Neighborhood
	Count int64 `json:"count"`
}
