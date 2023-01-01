package main

import (
	"log"
	"os"

	"context"

	"github.com/google/go-github/github"
	"github.com/urfave/cli/v2"

	"golang.org/x/oauth2"
)

const (
	// Replace this with your GitHub personal access token.
	personalAccessToken = "xxxxxxxx"
	// Replace this with the name of the file you want to create.
	filename = "example.txt"
)

func main() {
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(*cli.Context) error {
			xxx()

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// /Users/panachai/Library/Application Support/Microsoft Edge/Profile 1/Local Storage/leveldb

func xxx() {
	// Create a new GitHub client
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: personalAccessToken})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Create a new Gist
	gist := &github.Gist{
		Description: github.String("My Gist"),
		Public:      github.Bool(true),
		Files: map[github.GistFilename]github.GistFile{
			"my_file.txt": {
				Content: github.String("Hello, World!"),
			},
		},
	}
	// TODO: work on create gist
	// result, _, err := client.Gists.Create(ctx, gist)
	// if err != nil {
	// 	fmt.Printf("Error creating Gist: %v\n", err)
	// 	return
	// }
	// fmt.Printf("Created Gist %v\n", *result.HTMLURL)
}
