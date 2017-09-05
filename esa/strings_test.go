// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// copied from https://github.com/google/go-github/blob/master/github/strings.go

package esa

import (
	"fmt"
	"testing"
	"time"

	null "gopkg.in/guregu/null.v3"
)

func TestStringify(t *testing.T) {
	var tests = []struct {
		in  interface{}
		out string
	}{
		// basic types
		{"foo", `"foo"`},
		{123, `123`},
		{1.5, `1.5`},
		{false, `false`},
		{
			[]string{"a", "b"},
			`["a" "b"]`,
		},
		{
			struct {
				A []string
			}{nil},
			// nil slice is skipped
			`{}`,
		},
		{
			struct {
				A string
			}{"foo"},
			// structs not of a named type get no prefix
			`{A:"foo"}`,
		},

		// actual structs
		{
			Team{Name: "hoge", Privacy: "open", Description: "desc", Icon: "https://img.esa.io/", URL: "https://esa.io/"},
			`esa.Team{Name:"hoge", Privacy:"open", Description:"desc", Icon:"https://img.esa.io/", URL:"https://esa.io/"}`,
		},
	}

	for i, tt := range tests {
		s := Stringify(tt.in)
		if s != tt.out {
			t.Errorf("%d. Stringify(%q) => %q, want %q", i, tt.in, s, tt.out)
		}
	}
}

func TestString(t *testing.T) {
	var tests = []struct {
		in  interface{}
		out string
	}{
		{Team{Name: "hoge"}, `esa.Team{Name:"hoge", Privacy:"", Description:"", Icon:"", URL:""}`},
		{TeamList{
			Teams:      []*Team{{Name: "hoge"}},
			PrevPage:   null.NewInt(1, false),
			NextPage:   null.NewInt(3, false),
			TotalCount: 10,
			Page:       2,
			PerPage:    20,
			MaxPerPage: 100,
		}, `esa.TeamList{Teams:[esa.Team{Name:"hoge", Privacy:"", Description:"", Icon:"", URL:""}], PrevPage:null.Int{NullInt64:sql.NullInt64{Int64:1, Valid:false}}, NextPage:null.Int{NullInt64:sql.NullInt64{Int64:3, Valid:false}}, TotalCount:10, Page:2, PerPage:20, MaxPerPage:100}`},
		{Rate{Limit: 75, Remaining: 73, Reset: Timestamp{Time: time.Date(2017, 9, 5, 10, 0, 0, 0, time.UTC)}}, `esa.Rate{Limit:75, Remaining:73, Reset:esa.Timestamp{2017-09-05 10:00:00 +0000 UTC}, err:}`},
		{TeamStats{Members: 100}, `esa.TeamStats{Members:100, Posts:0, PostsWIP:0, PostsShipped:0, Comments:0, Stars:0, DailyActiveUsers:0, WeeklyActiveUsers:0, MonthlyActiveUsers:0}`},
		{InvitationURL{URL: "https://docs.esa.io/team/invitations/member-c05d112fa34870998ab4da1e98846ae3"}, `esa.InvitationURL{URL:"https://docs.esa.io/team/invitations/member-c05d112fa34870998ab4da1e98846ae3"}`},
	}

	for i, tt := range tests {
		s := tt.in.(fmt.Stringer).String()
		if s != tt.out {
			t.Errorf("%d. String() => %q, want %q", i, tt.in, tt.out)
		}
	}
}
