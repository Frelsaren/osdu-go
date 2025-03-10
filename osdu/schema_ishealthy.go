package osdu

// /api/search/v2/liveness_check

import "context"

func (s *SchemaService) IsHealthy(ctx context.Context) (bool, error) {
	req, err := s.client.NewRequest("GET", s.endpoint+"/liveness_check", nil)
	if err != nil {
		return false, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return false, err
	}

	return resp.StatusCode == 200, nil
}
