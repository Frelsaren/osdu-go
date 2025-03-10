package osdu

import (
	"context"
	"net/url"
	"strings"
)

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
