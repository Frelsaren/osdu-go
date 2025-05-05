package osdu

import (
	"context"
	"fmt"
)

func (s *EntitlementService) DeleteMember(ctx context.Context, memberEmail string) error {
	req, err := s.client.NewRequest("DELETE", fmt.Sprintf("%s/members/%s", s.endpoint, memberEmail), nil, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)

	return err
}
