package osdu

import "context"

func (s *StorageService) DeleteRecord(ctx context.Context, id string) error {
	req, err := s.client.NewRequest("POST", s.endpoint+"/"+id+":delete", nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}
