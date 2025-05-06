package osdu

import (
	"context"
	"fmt"
)

func (s *EntitlementService) DeleteGroup(ctx context.Context, groupEmail string) error {
	req, err := s.client.NewRequest("DELETE", fmt.Sprintf("%s/groups/%s", entitlementServicePath, groupEmail), nil, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)

	return err
}
