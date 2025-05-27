package osdu

import (
	"context"
	"fmt"
)

func (s *DatasetService) RevokeURL(ctx context.Context, kindSubtype string, body map[string]string) error {
	params := map[string]string{
		kindSubtype: kindSubtype,
	}

	req, err := s.client.NewRequest("GET", fmt.Sprintf("%s/storageInstructions", datasetServicePath), body, &params)

	if err != nil {
		return err
	}

	response := &StorageInstructions{}

	res, err := s.client.Do(ctx, req, &response)
	if err != nil {
		return err
	}

	if res.StatusCode != 204 {
		return fmt.Errorf("failed to revoke url: %d", res.StatusCode)
	}

	return nil
}
