package osdu

import (
	"context"
	"fmt"
	"strings"
)

func (s *StorageService) GetRecordVersion(ctx context.Context, id string, version int64, v *Record, attributes []string) error {
	params := make(map[string]string)

	if len(attributes) > 0 {
		params["attributes"] = strings.Join(attributes, ",")
	}

	req, err := s.client.NewRequest("GET", fmt.Sprintf("%s/records/%s/%d", storageServicePath, id, version), nil, &params)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, &v)
	if err != nil {
		return err
	}

	return nil
}
