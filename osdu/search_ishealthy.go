package osdu

import (
	"context"
	"fmt"
)

func (s *SearchService) IsHealthy(ctx context.Context) (bool, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("%s/liveness_check", s.endpoint), nil, nil)
	if err != nil {
		return false, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return false, err
	}

	return resp.StatusCode == 200, nil
}
