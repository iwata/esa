package esa

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
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
