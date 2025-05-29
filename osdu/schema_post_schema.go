package osdu

import (
	"context"
	"fmt"
)

func (s *SchemaService) PostSchema(ctx context.Context, body WriteSchemaBody) (WriteSchemaRespone, error) {
	var res WriteSchemaRespone
	req, err := s.client.NewRequest("POST", fmt.Sprintf("%s/schema", schemaServicePath), body, nil)
	if err != nil {
		return res, err
	}

	httpres, err := s.client.Do(ctx, req, &res.SchemaInfo)

	res.Code = int16(httpres.StatusCode)

	return res, err

}
