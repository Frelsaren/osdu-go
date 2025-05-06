package osdu

import (
	"context"
	"fmt"
)

func (s *FileService) DeleteFileMetadata(ctx context.Context, id string, v *Record) error {
	req, err := s.client.NewRequest("DELETE", fmt.Sprintf("%s/files/%s/metadata", fileServicePath, id), nil, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, &v)
	return err
}
