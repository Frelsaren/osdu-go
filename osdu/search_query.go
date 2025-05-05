package osdu

import (
	"context"
	"fmt"
)

func (s *SearchService) Query(ctx context.Context, body QueryParams) (SearchResponse, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("%s/query", s.endpoint), body, nil)
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
