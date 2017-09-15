# go-esa

[![Build Status](https://travis-ci.org/iwata/go-esa.svg?branch=master)](https://travis-ci.org/iwata/go-esa)
[![Coverage Status](https://coveralls.io/repos/github/iwata/go-esa/badge.svg?branch=master)](https://coveralls.io/github/iwata/go-esa?branch=master)

`go-esa` is a client library for esa.io API v1.

## Requirements

- Go 1.7+

## Installation

```sh
go get github.com/iwata/go-esa
```

## Sample code

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/iwata/go-esa/esa"
	"golang.org/x/oauth2"
)


func main() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("ESA_TOKEN")},
	)

	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)
	client := esa.NewClient(tc)

	// Fetch joining teams
	// ref. https://docs.esa.io/posts/102#4-1-0
	teamList, _, err := client.Teams.List(ctx)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("--- Team List ---")
	fmt.Printf("%v\n", teamList)

	team := os.Getenv("ESA_TEAM")

	// Get URL for invitation
	// ref. https://docs.esa.io/posts/102#12-1-0
	url, _, err := client.Invitations.GetURL(ctx, team)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("--- Invitation URL ---")
	fmt.Printf("%v\n", url)
	
	// Send invitations
	// ref. https://docs.esa.io/posts/102#13-1-0
	emails := []string{"hoge@example.com", "fuga@example.com"}
	resList, _, err := client.Invitations.SendToMember(ctx, team, &esa.InvitationMember{
		Member: &esa.InvitationEmails{Emails: emails},
	})
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("--- Invited List ---")
	fmt.Printf("%v\n", resList)

	// Send invitations
	// ref. https://docs.esa.io/posts/102#13-1-0
	emails := []string{"hoge@example.com", "fuga@example.com"}
	resList, _, err := client.Invitations.SendToMember(ctx, team, &esa.InvitationMember{
		Member: &esa.InvitationEmails{Emails: emails},
	})
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("--- Invited List ---")
	fmt.Printf("%v\n", resList)

	// Fetch pending invitations
	// ref. https://docs.esa.io/posts/102#13-2-0
	list, _, err := client.Invitations.PendingInvitations(ctx, team)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("--- Pending Invitations ---")
	fmt.Printf("%v\n", list)
}
```
