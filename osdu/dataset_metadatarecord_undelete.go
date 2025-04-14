package osdu

import (
	"context"
	"fmt"
)

func (s *DatasetService) MetadataUnDelete(ctx context.Context, id string) error {

	req, _ := s.client.NewRequest("POST", s.endpoint+"/metadataRecord/"+id+"/unDelete", nil, nil)

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("error: %s", resp.Status)
	}

	return nil
}
