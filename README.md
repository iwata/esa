# go-esa

[![Build Status](https://travis-ci.org/iwata/go-esa.svg?branch=master)](https://travis-ci.org/iwata/go-esa)
[![Coverage Status](https://coveralls.io/repos/github/iwata/go-esa/badge.svg?branch=master)](https://coveralls.io/github/iwata/go-esa?branch=master)

`go-esa` is a client library for esa.io API v1.

## Installation

```sh
go get github.com/iwata/go-esa
```

## Sample code

```go
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

	team := "esa-team"
	emails := []string{"hoge@example.com", "fuga@example.com"}

	// Send invitations
	// ref. https://docs.esa.io/posts/102#13-1-0
	resList, _, err := client.Invitations.SendToMember(ctx, team, &esa.InvitationMember{
		Member: &esa.InvitationEmails{Emails: emails},
	})
	if err != nil {
		log.Panic(err)
	}
	fmt.Sprintf("%v", reslist)

	// Fetch pending invitations
	// ref. https://docs.esa.io/posts/102#13-2-0
	list, _, err := esaClient.Invitations.GetList(ctx, team)
	if err != nil {
		log.Panic(err)
	}
	fmt.Sprintf("%v", list)
}
```
