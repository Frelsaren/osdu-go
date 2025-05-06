package osdu

import (
	"context"
	"fmt"
)

type GetUploadURLResponse struct {
	FileID   string            `json:"FileID"`
	Location map[string]string `json:"Location"`
}

func (s *FileService) GetUploadURL(ctx context.Context, expiryTime string) (GetUploadURLResponse, error) {
	var resp GetUploadURLResponse

	params := make(map[string]string)
	params["expiryTime"] = expiryTime
	req, err := s.client.NewRequest("GET", fmt.Sprintf("%s/files/uploadURL", fileServicePath), nil, &params)
	if err != nil {
		return resp, err
	}

	_, err = s.client.Do(ctx, req, &resp)

	return resp, err
}
