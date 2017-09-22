/*
Package esa is a client library for esa.io API v1.


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
}

For a full guide visit https://github.com/iwata/go-esa
*/
package esa
