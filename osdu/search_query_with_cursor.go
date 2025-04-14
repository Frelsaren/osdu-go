package osdu

import "context"

type QueryWithCursorParams struct {
	Kind            []string       `json:"kind"`
	Query           string         `json:"query,omitempty"`
	Offset          int            `json:"offset,omitempty"`
	Limit           int            `json:"limit,omitempty"`
	Sort            string         `json:"sort,omitempty"`
	TrackTotalCount bool           `json:"trackTotalCount,omitempty"`
	AggregateBy     []string       `json:"aggregateBy,omitempty"`
	SpatialFilter   *SpatialFilter `json:"spatialFilter,omitempty"`
	ReturnedFields  []string       `json:"returnedFields,omitempty"`
	Cursor          *string        `json:"cursor,omitempty"`
}

func (s *SearchService) QueryWithCursor(ctx context.Context, body QueryWithCursorParams) (SearchResponse, error) {
	req, err := s.client.NewRequest("POST", s.endpoint+"/query_with_cursor", body, nil)
	if err != nil {
		return SearchResponse{}, err
	}

	var res SearchResponse
	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return SearchResponse{}, err
	}

	return res, nil
}
