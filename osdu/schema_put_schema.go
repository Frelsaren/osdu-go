package osdu

import (
	"context"
	"fmt"
)

func (s *SchemaService) PutSchema(ctx context.Context, body WriteSchemaBody) (WriteSchemaRespone, error) {
	var res WriteSchemaRespone
	req, err := s.client.NewRequest("PUT", fmt.Sprintf("%s/schema", schemaServicePath), body, nil)
	if err != nil {
		return res, err
	}

	httpRes, err := s.client.Do(ctx, req, &res.SchemaInfo)

	res.Code = int16(httpRes.StatusCode)

	return res, err

}
