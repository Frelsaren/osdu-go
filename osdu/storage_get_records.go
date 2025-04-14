package osdu

import (
	"context"
	"net/url"
)

type GetRecordsParams struct {
	Records    []string `json:"records"`
	Attributes []string `json:"attributes"`
}

type RecordsResponse struct {
	Records        []Record `json:"records"`
	InvalidRecords []string `json:"invalidRecords"`
	RetryRecords   []string `json:"retryRecords"`
}

func (s *StorageService) GetRecords(ctx context.Context, params GetRecordsParams) (RecordsResponse, error) {
	url, _ := url.Parse(s.endpoint + "/query/records")
	var res RecordsResponse

	req, err := s.client.NewRequest("POST", url.String(), params, nil)
	if err != nil {
		return res, err
	}

	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
