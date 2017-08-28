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
	}

	for i, tt := range tests {
		s := tt.in.(fmt.Stringer).String()
		if s != tt.out {
			t.Errorf("%d. String() => %q, want %q", i, tt.in, tt.out)
		}
	}
}
