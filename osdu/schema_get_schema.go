package osdu

import (
	"context"
)

func (s *SchemaService) GetSchema(ctx context.Context, id string, v *interface{}) error {

	req, err := s.client.NewRequest("GET", s.endpoint+"/schema/"+id, nil, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, &v)

	return err

}
