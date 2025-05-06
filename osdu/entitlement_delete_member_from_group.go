package osdu

import (
	"context"
	"fmt"
)

func (s *EntitlementService) DeleteMemberFromGroup(ctx context.Context, memberEmail, groupEmail string) error {
	req, err := s.client.NewRequest("DELETE", fmt.Sprintf("%s/groups/%s/members/%s", entitlementServicePath, groupEmail, memberEmail), nil, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)

	return err
}
