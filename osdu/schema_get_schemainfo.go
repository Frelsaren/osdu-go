package osdu

import (
	"context"
	"encoding/json"
)

type GetSchemaParams struct {
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

type Response struct {
	SchemaInfos []SchemaInfo `json:"schemaInfos"`
	Offset      int          `json:"offset"`
	Count       int          `json:"count"`
	TotalCount  int          `json:"totalCount"`
}

func (s *SchemaService) GetSchemaInfo(ctx context.Context, params GetSchemaParams) (Response, error) {
	var response Response
	paramsAsMap, err := structToMap(params)
	if err != nil {
		return response, err
	}

	req, err := s.client.NewRequest("GET", s.endpoint+"/schema", nil, &paramsAsMap)
	if err != nil {
		return response, err
	}

	_, err = s.client.Do(ctx, req, &response)

	return response, err

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
