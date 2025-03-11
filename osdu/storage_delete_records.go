package osdu

import "context"

func (s *StorageService) DeleteRecords(ctx context.Context, ids []string) error {
	req, err := s.client.NewRequest("POST", s.endpoint+"/records/delete", ids, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}
