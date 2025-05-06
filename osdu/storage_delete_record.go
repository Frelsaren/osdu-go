package osdu

import (
	"context"
	"fmt"
)

func (s *StorageService) DeleteRecord(ctx context.Context, id string) error {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("%s/records/%s:delete", storageServicePath, id), nil, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}
