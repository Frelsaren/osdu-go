package osdu

import (
	"context"
)

type AddMemberResponse struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

func (s *EntitlementService) AddMember(ctx context.Context, groupEmail, email, role string) (AddMemberResponse, error) {
	var res AddMemberResponse

	body := AddMemberResponse{
		Email: email,
		Role:  role,
	}

	req, err := s.client.NewRequest("POST", s.endpoint+"/groups/"+groupEmail+"/members", body, nil)
	if err != nil {
		return res, err
	}

	_, err = s.client.Do(ctx, req, &res)

	return res, err
}
