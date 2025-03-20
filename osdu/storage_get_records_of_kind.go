package osdu

import (
	"context"
	"net/url"
)

type GetRecordOfKindParams struct {
	Cursor *string
	Limit  *int32
	Kind   string
}

type GetRecordOfKindReturn struct {
	Cursor  string   `json:"cursor"`
	Results []string `json:"results"`
}

func (s *StorageService) GetRecordsOfKind(ctx context.Context, params GetRecordOfKindParams) (GetRecordOfKindReturn, error) {
	url, _ := url.Parse(s.endpoint + "/query/records?kind=" + params.Kind)
	queryparams := make(map[string]string)
	res := GetRecordOfKindReturn{}

	queryparams["kind"] = params.Kind
	if params.Cursor != nil {
		queryparams["cursor"] = *params.Cursor
	}
	if params.Limit != nil {
		queryparams["limit"] = string(*params.Limit)
	}

	req, err := s.client.NewRequest("GET", url.String(), nil, &queryparams)
	if err != nil {
		return res, err
	}

	_, err = s.client.Do(ctx, req, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
