package osdu

import (
	"context"
	"fmt"
)

func (s *StorageService) DeleteRecords(ctx context.Context, ids []string) error {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("%s/records/delete", s.endpoint), ids, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}
