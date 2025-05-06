package osdu

import (
	"context"
	"fmt"
)

type createFileMetadataResponse struct {
	Id string `json:"id"`
}

func (s *FileService) CreateFileMetadata(ctx context.Context, id string, v *Record) (string, error) {
	req, err := s.client.NewRequest("DELETE", fmt.Sprintf("%s/files/metadata", fileServicePath), &v, nil)
	if err != nil {
		return "", err
	}

	var response createFileMetadataResponse
	_, err = s.client.Do(ctx, req, &response)
	if err != nil {
		return "", err
	}
	return response.Id, err
}
