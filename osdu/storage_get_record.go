package osdu

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

func (s *StorageService) GetRecord(ctx context.Context, id string, v *Record, attributes []string) error {
	url, _ := url.Parse(fmt.Sprintf("%s/records/%s", s.endpoint, id))
	params := make(map[string]string)

	if len(attributes) > 0 {
		params["attributes"] = strings.Join(attributes, ",")
	}

	req, err := s.client.NewRequest("GET", url.String(), nil, &params)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, &v)
	if err != nil {
		return err
	}

	return nil
}
