package osdu

import (
	"context"
	"fmt"
)

type DatasetRegistry struct {
	DatasetRegistries []Record `json:"datasetRegistries"`
}

type getDatasetRegistriesBody struct {
	DatasetRegistryIds []string `json:"datasetRegistryIds"`
}

func (s *DatasetService) GetDatasetRegistries(ctx context.Context, datasetIDs []string) (*DatasetRegistry, error) {
	body := getDatasetRegistriesBody{
		DatasetRegistryIds: datasetIDs,
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("%s/getDatasetRegistry", s.endpoint), body, nil)
	if err != nil {
		return nil, err
	}

	datasetRegistry := &DatasetRegistry{}
	resp, err := s.client.Do(ctx, req, datasetRegistry)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error: %s", resp.Status)
	}

	return datasetRegistry, nil
}
