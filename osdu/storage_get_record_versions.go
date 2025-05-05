package osdu

import (
	"context"
	"fmt"
)

type RecordVersionsResponse struct {
	RecordID string  `json:"recordId"`
	Versions []int64 `json:"versions"`
}

func (s *StorageService) GetRecordVersions(ctx context.Context, id string) (RecordVersionsResponse, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("%s/records/versions/%s", s.endpoint, id), nil, nil)
	if err != nil {
		return RecordVersionsResponse{}, err
	}

	resp := RecordVersionsResponse{}
	_, err = s.client.Do(ctx, req, &resp)
	if err != nil {
		return RecordVersionsResponse{}, err
	}

	return resp, nil
}
