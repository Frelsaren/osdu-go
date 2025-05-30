package osdu

import (
	"context"
	"fmt"
)

func (s *DatasetService) IsHealthy(ctx context.Context) (bool, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("%s/liveness_check", datasetServicePath), nil, nil)
	if err != nil {
		return false, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return false, err
	}

	return resp.StatusCode == 200, nil
}
