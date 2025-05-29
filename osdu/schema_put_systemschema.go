package osdu

import (
	"context"
	"fmt"
)

func (s *SchemaService) PutSystemSchema(ctx context.Context, body WriteSchemaBody) (WriteSchemaRespone, error) {
	var res WriteSchemaRespone
	req, err := s.client.NewRequest("PUT", fmt.Sprintf("%s/schema/system", schemaServicePath), body, nil)
	if err != nil {
		return res, err
	}

	httpres, err := s.client.Do(ctx, req, &res.SchemaInfo)

	res.Code = int16(httpres.StatusCode)

	return res, err

}
