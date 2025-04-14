package osdu

import "context"

type RegisterDatasetRequest struct {
	DatasetRegistries []Record `json:"datasetRegistries"`
}

func (s *DatasetService) RegisterDataset(ctx context.Context, requestBody RegisterDatasetRequest) (*RegisterDatasetRequest, error) {
	req, err := s.client.NewRequest("PUT", s.endpoint+"/registerDataset", requestBody, nil)
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
