package esa

import (
	"context"
	"fmt"

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

// TeamStats represents a statistics of esa team.
type TeamStats struct {
	Members            int `json:"members"`
	Posts              int `json:"posts"`
	PostsWIP           int `json:"posts_wip"`
	PostsShipped       int `json:"posts_shipped"`
	Comments           int `json:"comments"`
	Stars              int `json:"stars"`
	DailyActiveUsers   int `json:"daily_active_users"`
	WeeklyActiveUsers  int `json:"weekly_active_users"`
	MonthlyActiveUsers int `json:"monthly_active_users"`
}

func (s TeamStats) String() string {
	return Stringify(s)
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

// Get fetches a team by name.
//
// API docs: https://docs.esa.io/posts/102#4-2-0
func (s *TeamsService) Get(ctx context.Context, team string) (*Team, *Response, error) {
	u := fmt.Sprintf("teams/%s", team)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	t := &Team{}
	resp, err := s.client.Do(ctx, req, t)
	if err != nil {
		return nil, resp, err
	}
	return t, resp, nil
}

// GetStats fetches a statistics of team by name.
//
// API docs: https://docs.esa.io/posts/102#5-1-0
func (s *TeamsService) GetStats(ctx context.Context, team string) (*TeamStats, *Response, error) {
	u := fmt.Sprintf("teams/%s/stats", team)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	st := &TeamStats{}
	resp, err := s.client.Do(ctx, req, st)
	if err != nil {
		return nil, resp, err
	}
	return st, resp, nil
}
