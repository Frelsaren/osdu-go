package osdu

import (
	"context"
)

func (s *SchemaService) PutSchema(ctx context.Context, body WriteSchemaBody) (WriteSchemaRespone, error) {
	var res WriteSchemaRespone
	req, err := s.client.NewRequest("PUT", s.endpoint+"/schema", body, nil)
	if err != nil {
		return res, err
	}

	httpres, err := s.client.Do(ctx, req, &res.Res)

	res.Code = int16(httpres.StatusCode)

	return res, err

}
