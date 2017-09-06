package esa

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestInvitationsService_GetURL(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/teams/hoge/invitation", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{})
		fmt.Fprint(w, `{
  "url": "https://docs.esa.io/team/invitations/member-c05d112fa34870998ab4da1e98846ae3"
}`)
	})

	u, _, err := client.Invitations.GetURL(context.Background(), "hoge")
	if err != nil {
		t.Errorf("Invitations.GetURL returned error: %v", err)
	}

	want := &InvitationURL{
		URL: "https://docs.esa.io/team/invitations/member-c05d112fa34870998ab4da1e98846ae3",
	}
	if !reflect.DeepEqual(u, want) {
		t.Errorf("InvitationsService.GetURL returned %+v, want %+v", u, want)
	}
}

func TestInvitationsService_GetURL_ErrorStatus(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/teams/hoge/invitation", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Forbidden", http.StatusForbidden)
	})

	_, resp, err := client.Invitations.GetURL(context.Background(), "hoge")
	if err == nil {
		t.Error("Expected error to be returned.")
	}

	if resp == nil {
		t.Error("InvitationsService.GetURL returned Reponse, too")
	}
}

func TestInvitationsService_RegenerateURL(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/teams/hoge/invitation_regenerator", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{})
		fmt.Fprint(w, `{
  "url": "https://docs.esa.io/team/invitations/member-58891f72edcbb8ac22f1e5548b0128d9"
}`)
	})

	u, _, err := client.Invitations.RegenerateURL(context.Background(), "hoge")
	if err != nil {
		t.Errorf("Invitations.RegenerateURL returned error: %v", err)
	}

	want := &InvitationURL{
		URL: "https://docs.esa.io/team/invitations/member-58891f72edcbb8ac22f1e5548b0128d9",
	}
	if !reflect.DeepEqual(u, want) {
		t.Errorf("InvitationsService.RegenerateURL returned %+v, want %+v", u, want)
	}
}

func TestInvitationsService_RegenerateURL_ErrorStatus(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/teams/hoge/invitation_regenerator", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Forbidden", http.StatusForbidden)
	})

	_, resp, err := client.Invitations.RegenerateURL(context.Background(), "hoge")
	if err == nil {
		t.Error("Expected error to be returned.")
	}

	if resp == nil {
		t.Error("InvitationsService.RegenerateURL returned Reponse, too")
	}
}

func TestInvitationsService_SendToMember(t *testing.T) {
	setup()
	defer teardown()

	input := &InvitationMember{
		Member: &InvitationEmails{
			Emails: []string{"for@example.com", "bar@example.com"},
		},
	}

	mux.HandleFunc("/v1/teams/hoge/invitations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(InvitationMember)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{
    "invitations": [
        {
            "email": "foo@example.com",
            "code": "mee93383edf699b525e01842d34078e28",
            "expires_at": "2017-08-17T12:00:41+09:00",
            "url": "https://docs.esa.io/team/invitations/mee93383edf699b525e01842d34078e28/join"
        },
        {
            "email": "bar@example.com",
            "code": "mc542eed211a8e4f1db6ccccb14fcda9d",
            "expires_at": "2017-08-17T12:00:44+09:00",
            "url": "https://docs.esa.io/team/invitations/mc542eed211a8e4f1db6ccccb14fcda9d/join"
        }
    ]
	}`)
	})

	l, _, err := client.Invitations.SendToMember(context.Background(), "hoge", input)
	if err != nil {
		t.Errorf("Invitations.SendToMember returned error: %v", err)
	}

	want := &InvitationList{
		Invitations: []*Invitation{
			{
				Email:     "foo@example.com",
				Code:      "mee93383edf699b525e01842d34078e28",
				ExpiresAt: Timestamp{time.Date(2017, 8, 17, 12, 0, 41, 0, jst).Local()},
				URL:       "https://docs.esa.io/team/invitations/mee93383edf699b525e01842d34078e28/join",
			},
			{
				Email:     "bar@example.com",
				Code:      "mc542eed211a8e4f1db6ccccb14fcda9d",
				ExpiresAt: Timestamp{time.Date(2017, 8, 17, 12, 0, 44, 0, jst).Local()},
				URL:       "https://docs.esa.io/team/invitations/mc542eed211a8e4f1db6ccccb14fcda9d/join",
			},
		},
	}
	if !reflect.DeepEqual(l, want) {
		t.Errorf("InvitationsService.SendToMember returned %+v, want %+v", l, want)
	}
}

func TestInvitationsService_SendToMember_ErrorStatus(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/teams/hoge/invitations", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Forbidden", http.StatusForbidden)
	})

	_, resp, err := client.Invitations.SendToMember(context.Background(), "hoge", &InvitationMember{})
	if err == nil {
		t.Error("Expected error to be returned.")
	}

	if resp == nil {
		t.Error("InvitationsService.SendToMember returned Reponse, too")
	}
}

func TestInvitationsService_GetList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/teams/hoge/invitations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{})
		fmt.Fprint(w, `{
    "invitations": [
        {
            "email": "foo@example.com",
            "code": "mee93383edf699b525e01842d34078e28",
            "expires_at": "2017-08-17T12:00:41+09:00",
            "url": "https://docs.esa.io/team/invitations/mee93383edf699b525e01842d34078e28/join"
        },
        {
            "email": "bar@example.com",
            "code": "mc542eed211a8e4f1db6ccccb14fcda9d",
            "expires_at": "2017-08-17T12:00:44+09:00",
            "url": "https://docs.esa.io/team/invitations/mc542eed211a8e4f1db6ccccb14fcda9d/join"
        }
    ],
    "prev_page": null,
    "next_page": null,
    "total_count": 2,
    "page": 1,
    "per_page": 20,
    "max_per_page": 100
	}`)
	})

	l, _, err := client.Invitations.GetList(context.Background(), "hoge")
	if err != nil {
		t.Errorf("Invitations.GetList returned error: %v", err)
	}

	want := &InvitationList{
		Invitations: []*Invitation{
			{
				Email:     "foo@example.com",
				Code:      "mee93383edf699b525e01842d34078e28",
				ExpiresAt: Timestamp{time.Date(2017, 8, 17, 12, 0, 41, 0, jst).Local()},
				URL:       "https://docs.esa.io/team/invitations/mee93383edf699b525e01842d34078e28/join",
			},
			{
				Email:     "bar@example.com",
				Code:      "mc542eed211a8e4f1db6ccccb14fcda9d",
				ExpiresAt: Timestamp{time.Date(2017, 8, 17, 12, 0, 44, 0, jst).Local()},
				URL:       "https://docs.esa.io/team/invitations/mc542eed211a8e4f1db6ccccb14fcda9d/join",
			},
		},
		PrevPage:   0,
		NextPage:   0,
		TotalCount: 2,
		Page:       1,
		PerPage:    20,
		MaxPerPage: 100,
	}
	if !reflect.DeepEqual(l, want) {
		t.Errorf("InvitationsService.GetList returned %+v, want %+v", l, want)
	}
}

func TestInvitationsService_GetList_ErrorStatus(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/teams/hoge/invitations", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Forbidden", http.StatusForbidden)
	})

	_, resp, err := client.Invitations.GetList(context.Background(), "hoge")
	if err == nil {
		t.Error("Expected error to be returned.")
	}

	if resp == nil {
		t.Error("InvitationsService.GetList returned Reponse, too")
	}
}

func TestInvitationsService_Cancel(t *testing.T) {
	setup()
	defer teardown()

	input := "mee93383edf699b525e01842d34078e28"

	mux.HandleFunc("/v1/teams/hoge/invitations/"+input, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testFormValues(t, r, values{})
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.Invitations.Cancel(context.Background(), "hoge", input)
	if err != nil {
		t.Errorf("Invitations.Cancel returned error: %v", err)
	}

	if resp == nil {
		t.Error("Invitations.Cancel returned Response.")
	}
}

func TestInvitationsService_Cancel_ErrorStatus(t *testing.T) {
	setup()
	defer teardown()

	input := "mee93383edf699b525e01842d34078e28"
	mux.HandleFunc("/v1/teams/hoge/invitations"+input, func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Forbidden", http.StatusForbidden)
	})

	resp, err := client.Invitations.Cancel(context.Background(), "hoge", input)
	if err == nil {
		t.Error("Expected error to be returned.")
	}

	if resp == nil {
		t.Error("InvitationsService.Cancel returned Reponse.")
	}
}
