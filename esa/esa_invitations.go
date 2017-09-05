package esa

import (
	"context"
	"fmt"
)

// InvitationsService provides access to invitations related functions
// in the esa API.
//
// API docs:
//  - https://docs.esa.io/posts/102#12-0-0
//  - https://docs.esa.io/posts/102#13-0-0
type InvitationsService service

// InvitationURL represents an invitation URL.
type InvitationURL struct {
	URL string `json:"url"`
}

func (u InvitationURL) String() string {
	return Stringify(u)
}

// GetURL fetches a team by name.
//
// API docs: https://docs.esa.io/posts/102#12-1-0
func (s *InvitationsService) GetURL(ctx context.Context, team string) (*InvitationURL, *Response, error) {
	u := fmt.Sprintf("teams/%s/invitation", team)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	url := &InvitationURL{}
	resp, err := s.client.Do(ctx, req, url)
	if err != nil {
		return nil, resp, err
	}
	return url, resp, nil
}

// RegenerateURL regenerates an invitation URL
//
// API docs: https://docs.esa.io/posts/102#12-2-0
func (s *InvitationsService) RegenerateURL(ctx context.Context, team string) (*InvitationURL, *Response, error) {
	u := fmt.Sprintf("teams/%s/invitation_regenerator", team)
	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, nil, err
	}

	url := &InvitationURL{}
	resp, err := s.client.Do(ctx, req, url)
	if err != nil {
		return nil, resp, err
	}
	return url, resp, nil
}
