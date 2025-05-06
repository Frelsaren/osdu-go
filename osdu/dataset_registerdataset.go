package osdu

import (
	"context"
	"fmt"
)

type RegisterDatasetRequest struct {
	DatasetRegistries []Record `json:"datasetRegistries"`
}

func (s *DatasetService) RegisterDataset(ctx context.Context, requestBody RegisterDatasetRequest) (*RegisterDatasetRequest, error) {
	req, err := s.client.NewRequest("PUT", fmt.Sprintf("%s/registerDataset", datasetServicePath), requestBody, nil)
	if err != nil {
		return nil, err
	}

	var response RegisterDatasetRequest
	_, err = s.client.Do(ctx, req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
