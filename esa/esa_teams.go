package esa

import (
	"context"

	null "gopkg.in/guregu/null.v3"
)

// TeamsService provides access to the team related functions
// in the esa API.
//
// API docs: https://docs.esa.io/posts/102#4-0-0
type TeamsService service

// Team represents a esa team.
type Team struct {
	Name        string `json:"name"`
	Privacy     string `json:"privacy"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	URL         string `json:"url"`
}

func (t Team) String() string {
	return Stringify(t)
}

// TeamList represents a list of esa teams.
type TeamList struct {
	Teams      []*Team  `json:"teams"`
	PrevPage   null.Int `json:"prev_page"`
	NextPage   null.Int `json:"next_page"`
	TotalCount int      `json:"total_count"`
	Page       int      `json:"page"`
	PerPage    int      `json:"per_page"`
	MaxPerPage int      `json:"max_per_page"`
}

func (l TeamList) String() string {
	return Stringify(l)
}

// List lists all teams
//
// API docs: https://docs.esa.io/posts/102#4-1-0
func (s *TeamsService) List(ctx context.Context) (*TeamList, *Response, error) {
	req, err := s.client.NewRequest("GET", "teams", nil)
	if err != nil {
		return nil, nil, err
	}

	list := &TeamList{}
	resp, err := s.client.Do(ctx, req, list)
	if err != nil {
		return nil, resp, err
	}
	return list, resp, nil
}
