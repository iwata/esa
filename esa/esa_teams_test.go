package esa

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	null "gopkg.in/guregu/null.v3"
)

func TestTeamsService_ListAll(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/teams", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{})
		fmt.Fprint(w, `{
  "teams": [
    {
      "name": "docs",
      "privacy": "open",
      "description": "esa.io official documents",
      "icon": "https://img.esa.io/uploads/production/teams/105/icon/thumb_m_0537ab827c4b0c18b60af6cdd94f239c.png",
      "url": "https://docs.esa.io/"
    }
  ],
  "prev_page": null,
  "next_page": null,
  "total_count": 1,
  "page": 1,
  "per_page": 20,
  "max_per_page": 100
}`)
	})

	list, _, err := client.Teams.List(context.Background())
	if err != nil {
		t.Errorf("Teams.List returned error: %v", err)
	}

	want := &TeamList{
		Teams: []*Team{
			{
				Name:        "docs",
				Privacy:     "open",
				Description: "esa.io official documents",
				Icon:        "https://img.esa.io/uploads/production/teams/105/icon/thumb_m_0537ab827c4b0c18b60af6cdd94f239c.png",
				URL:         "https://docs.esa.io/",
			},
		},
		PrevPage:   null.NewInt(0, false),
		NextPage:   null.NewInt(0, false),
		TotalCount: 1,
		Page:       1,
		PerPage:    20,
		MaxPerPage: 100,
	}
	if !reflect.DeepEqual(list, want) {
		t.Errorf("TeamsService.List returned %+v, want %+v", list, want)
	}
}

func TestTeamsService_ListAll_ErrorStatus(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/teams", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Forbidden", http.StatusForbidden)
	})

	_, resp, err := client.Teams.List(context.Background())
	if err == nil {
		t.Error("Expected error to be returned.")
	}

	if resp == nil {
		t.Error("TeamsService.List returned Reponse, too")
	}
}
