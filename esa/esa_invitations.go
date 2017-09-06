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

// Invitation represents an invitation to esa team.
type Invitation struct {
	Email     string    `json:"email"`
	Code      string    `json:"code"`
	ExpiresAt Timestamp `json:"expires_at"`
	URL       string    `json:"url"`
}

func (i Invitation) String() string {
	return Stringify(i)
}

// InvitationList represents a list of invitations.
type InvitationList struct {
	Invitations []*Invitation `json:"invitations"`
	PrevPage    int           `json:"prev_page,omitempty"`
	NextPage    int           `json:"next_page,omitempty"`
	TotalCount  int           `json:"total_count,omitempty"`
	Page        int           `json:"page,omitempty"`
	PerPage     int           `json:"per_page,omitempty"`
	MaxPerPage  int           `json:"max_per_page,omitempty"`
}

func (l InvitationList) String() string {
	return Stringify(l)
}

// InvitationMember represents members that wanna send invitations
type InvitationMember struct {
	Member *InvitationEmails `json:"member"`
}

func (m InvitationMember) String() string {
	return Stringify(m)
}

// InvitationEmails represents emails that wanna send invitations
type InvitationEmails struct {
	Emails []string `json:"emails"`
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

// SendToMember send invitation emails
//
// API docs: https://docs.esa.io/posts/102#12-1-0
func (s *InvitationsService) SendToMember(ctx context.Context, team string, member *InvitationMember) (*InvitationList, *Response, error) {
	u := fmt.Sprintf("teams/%s/invitations", team)
	req, err := s.client.NewRequest("POST", u, member)
	if err != nil {
		return nil, nil, err
	}

	l := &InvitationList{}
	resp, err := s.client.Do(ctx, req, l)
	if err != nil {
		return nil, resp, err
	}
	return l, resp, nil
}

// GetList fetches a list of invitations.
//
// API docs: https://docs.esa.io/posts/102#13-2-0
func (s *InvitationsService) GetList(ctx context.Context, team string) (*InvitationList, *Response, error) {
	u := fmt.Sprintf("teams/%s/invitations", team)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	l := &InvitationList{}
	resp, err := s.client.Do(ctx, req, l)
	if err != nil {
		return nil, resp, err
	}
	return l, resp, nil
}

// Cancel deletes an invitation by an invitation code.
//
// API docs: https://docs.esa.io/posts/102#13-3-0
func (s *InvitationsService) Cancel(ctx context.Context, team string, code string) (*Response, error) {
	u := fmt.Sprintf("teams/%s/invitations/%s", team, code)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
