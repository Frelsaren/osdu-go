package osdu

import (
	"context"
	"fmt"
)

type CreateGroupBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateGroupResponse struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (s *EntitlementService) CreateGroup(ctx context.Context, body CreateGroupBody) (CreateGroupResponse, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("%s/groups", entitlementServicePath), body, nil)
	if err != nil {
		return CreateGroupResponse{}, err
	}

	var res CreateGroupResponse
	_, err = s.client.Do(ctx, req, &res)

	if err != nil {
		return CreateGroupResponse{}, err
	}

	return res, nil
}
