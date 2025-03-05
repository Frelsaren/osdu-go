package osdu

import (
	"context"
	"net/url"
	"strings"
)

// StorageService provides access to the Storage API
type StorageService service

func (s *StorageService) IsHealthy(ctx context.Context) (bool, error) {
	req, err := s.client.NewRequest("GET", s.endpoint+"/health/readiness_check", nil)
	if err != nil {
		return false, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return false, err
	}

	return resp.StatusCode == 200, nil
}

func (s *StorageService) GetRecord(ctx context.Context, id string, v *interface{}, attributes []string) error {
	url, _ := url.Parse(s.endpoint + "/" + id)

	if len(attributes) > 0 {
		q := url.Query()
		q.Set("attributes", strings.Join(attributes, ","))
		url.RawQuery = q.Encode()
	}

	req, err := s.client.NewRequest("GET", url.String(), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, &v)
	if err != nil {
		return err
	}

	return nil
}

func (s *StorageService) GetRecordVersion(ctx context.Context, id string, version string, v *interface{}) error {
	req, err := s.client.NewRequest("GET", s.endpoint+"/"+id+"/"+version, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, &v)
	if err != nil {
		return err
	}

	return nil
}

type RecordVersionsResponse struct {
	RecordID string  `json:"recordId"`
	Versions []int64 `json:"versions"`
}

func (s *StorageService) GetRecordVersions(ctx context.Context, id string) (RecordVersionsResponse, error) {
	req, err := s.client.NewRequest("GET", s.endpoint+"/versions/"+id, nil)
	if err != nil {
		return RecordVersionsResponse{}, err
	}

	resp := RecordVersionsResponse{}
	_, err = s.client.Do(ctx, req, &resp)
	if err != nil {
		return RecordVersionsResponse{}, err
	}

	return resp, nil
}

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

func (s *StorageService) DeleteRecords(ctx context.Context, ids []string) error {
	req, err := s.client.NewRequest("POST", s.endpoint+"/delete", ids)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}
