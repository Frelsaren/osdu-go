package osdu

import "context"

type SearchService service

type QueryParams struct {
	Kind            []string      `json:"kind"`
	Query           string        `json:"query,omitempty"`
	Offset          int           `json:"offset,omitempty"`
	Limit           int           `json:"limit,omitempty"`
	Sort            string        `json:"sort,omitempty"`
	TrackTotalCount bool          `json:"trackTotalCount,omitempty"`
	AggregateBy     []string      `json:"aggregateBy,omitempty"`
	SpatialFilter   SpatialFilter `json:"spatialFilter,omitempty"`
}

type LatLon struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type SpatialFilter struct {
	Field      string `json:"field"`
	ByDistance struct {
		Point    LatLon  `json:"point"`
		Distance float64 `json:"distance"`
	} `json:"byDistance"`
	ByBoundingBox struct {
		TopLeft     LatLon `json:"topLeft"`
		BottomRight LatLon `json:"bottomRight"`
	} `json:"byBoundingBox"`
	ByGeoPolygon struct {
		Points []LatLon `json:"points"`
	} `json:"byGeoPolygon"`
	ByIntersection struct {
		Polygons []struct {
			Points []LatLon `json:"points"`
		} `json:"polygons"`
	} `json:"byIntersection"`
}

type SearchResponse struct {
	Results      []interface{} `json:"results"`
	TotalCount   int           `json:"totalCount"`
	Aggregations []struct {
		Key   string `json:"key"`
		Count int    `json:"count"`
	} `json:"aggregations"`
}

func (s *SearchService) Query(ctx context.Context, params QueryParams) (SearchResponse, error) {
	req, err := s.client.NewRequest("POST", s.endpoint, params)
	if err != nil {
		return SearchResponse{}, err
	}

	var searchResponse SearchResponse
	_, err = s.client.Do(ctx, req, &searchResponse)
	if err != nil {
		return SearchResponse{}, err
	}

	return searchResponse, nil
}
