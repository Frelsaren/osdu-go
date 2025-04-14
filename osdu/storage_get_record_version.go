package osdu

import "context"

func (s *StorageService) GetRecordVersion(ctx context.Context, id string, version string, v *Record) error {
	req, err := s.client.NewRequest("GET", s.endpoint+"/records/"+id+"/"+version, nil, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, &v)
	if err != nil {
		return err
	}

	return nil
}
