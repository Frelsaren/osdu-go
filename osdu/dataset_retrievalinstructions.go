package osdu

import "context"

type retrievalInstructionsBody struct {
	DatasetRegistryIds []string `json:"datasetRegistryIds"`
}

type RetrievalInstructions struct {
	Datasets []struct {
		DatasetRegistryId   string                 `json:"datasetRegistryId"`
		RetrievalProperties map[string]interface{} `json:"retrievalProperties"`
		ProviderKey         string                 `json:"providerKey"`
	} `json:"datasets"`
}

func (s *DatasetService) GetRetrievalInstructions(ctx context.Context, ids []string, expiryTime string) (*RetrievalInstructions, error) {
	params := map[string]string{
		expiryTime: expiryTime,
	}
	body := retrievalInstructionsBody{
		DatasetRegistryIds: ids,
	}
	req, err := s.client.NewRequest("GET", s.endpoint+"/retrievalInstructions", body, &params)

	if err != nil {
		return nil, err
	}

	response := &RetrievalInstructions{}

	_, err = s.client.Do(ctx, req, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
