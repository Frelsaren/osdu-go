package osdu

import (
	"context"
	"encoding/json"
	"fmt"
)

type GetSchemaInfoParams struct {
	Authority          *string `json:"authority,omitempty"`
	Source             *string `json:"source,omitempty"`
	EntityType         *string `json:"entityType,omitempty"`
	SchemaVersionMajor *string `json:"schemaVersionMajor,omitempty"`
	SchemaVersionMinor *string `json:"schemaVersionMinor,omitempty"`
	SchemaVersionPatch *string `json:"schemaVersionPatch,omitempty"`
	Status             *string `json:"status,omitempty"`
	Scope              *string `json:"scope,omitempty"`
	LatestVersion      *bool   `json:"latestVersion,omitempty"`
	Limit              *int32  `json:"limit,omitempty"`
	Offset             *int32  `json:"offset,omitempty"`
}

type SchemaInfoResponse struct {
	SchemaInfos []SchemaInfo `json:"schemaInfos"`
	Offset      int          `json:"offset"`
	Count       int          `json:"count"`
	TotalCount  int          `json:"totalCount"`
}

func (s *SchemaService) GetSchemaInfo(ctx context.Context, params GetSchemaInfoParams) (SchemaInfoResponse, error) {
	var res SchemaInfoResponse
	paramsAsMap, err := structToMap(params)
	if err != nil {
		return res, err
	}

	req, err := s.client.NewRequest("GET", fmt.Sprintf("%s/schema", s.endpoint), nil, &paramsAsMap)
	if err != nil {
		return res, err
	}

	_, err = s.client.Do(ctx, req, &res)

	return res, err

}

func structToMap(input interface{}) (map[string]string, error) {
	var res map[string]string
	marshaled, err := json.Marshal(input)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(marshaled, &res)
	return res, err
}
