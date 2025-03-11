package osdu

import "context"

func (s *StorageService) IsHealthy(ctx context.Context) (bool, error) {
	req, err := s.client.NewRequest("GET", s.endpoint+"/liveness_check", nil, nil)
	if err != nil {
		return false, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return false, err
	}

	return resp.StatusCode == 200, nil
}
