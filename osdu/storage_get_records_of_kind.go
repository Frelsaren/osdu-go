package osdu

import (
	"context"
	"fmt"
	"net/url"
)

type GetRecordOfKindParams struct {
	Cursor *string
	Limit  *int32
	Kind   string
}

type RecordsOfKindResponse struct {
	Cursor  string   `json:"cursor"`
	Results []string `json:"results"`
}

func (s *StorageService) GetRecordsOfKind(ctx context.Context, params GetRecordOfKindParams) (RecordsOfKindResponse, error) {
	url, _ := url.Parse(fmt.Sprintf("%s/records", storageServicePath))
	queryparams := make(map[string]string)
	res := RecordsOfKindResponse{}

	queryparams["kind"] = params.Kind
	if params.Cursor != nil {
		queryparams["cursor"] = *params.Cursor
	}
	if params.Limit != nil {
		queryparams["limit"] = string(*params.Limit)
	}

	req, err := s.client.NewRequest("GET", url.String(), nil, &queryparams)
	urlString := req.URL.String()
	fmt.Println("Request URL:", urlString)
	if err != nil {
		return res, err
	}

	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
