package osdu

import (
	"context"
	"fmt"
)

type downloadUrlResponse struct {
	SignedUrl string `json:"SignedUrl"`
}

func (s *FileService) GetDownloadURL(ctx context.Context, id, expiryTime string) (string, error) {
	params := make(map[string]string)
	params["expiryTime"] = expiryTime
	req, err := s.client.NewRequest("GET", fmt.Sprintf("%s/files/%s/downloadURL", fileServicePath, id), nil, &params)
	if err != nil {
		return "", err
	}

	var resp downloadUrlResponse
	_, err = s.client.Do(ctx, req, &resp)
	if err != nil {
		return "", err
	}

	return resp.SignedUrl, nil
}
