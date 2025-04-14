package osdu

import "context"

type StorageInstructions struct {
	StorageLocation map[string]interface{} `json:"storageLocation"`
	ProviderKey     string                 `json:"providerKey"`
}

func (s *DatasetService) GenerateStorageInstructions(ctx context.Context, kindSubtype, expiryTime string) (*StorageInstructions, error) {
	params := map[string]string{
		expiryTime:  expiryTime,
		kindSubtype: kindSubtype,
	}

	req, err := s.client.NewRequest("GET", s.endpoint+"/storageInstructions", nil, &params)

	if err != nil {
		return nil, err
	}

	response := &StorageInstructions{}

	_, err = s.client.Do(ctx, req, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
