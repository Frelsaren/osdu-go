package osdu

import (
	"context"
	"fmt"
)

func (s *SchemaService) GetSchema(ctx context.Context, id string, v *interface{}) error {

	req, err := s.client.NewRequest("GET", fmt.Sprintf("%s/schema/%s", s.endpoint, id), nil, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, &v)

	return err

}
