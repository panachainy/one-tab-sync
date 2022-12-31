package main

import (
    "fmt"
    "log"
    "os"

    "github.com/urfave/cli/v2"
    "context"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/jwt"

	"github.com/google/go-github/github"
)
const (
	// Replace this with your GitHub personal access token.
	personalAccessToken = "YOUR_TOKEN_HERE"
	// Replace this with the name of the file you want to create.
	filename = "example.txt"
	// Replace this with the contents of the file you want to create.
	fileContents = "This is an example file."
)

func main() {
    app := &cli.App{
        Name:  "boom",
        Usage: "make an explosive entrance",
        Action: func(*cli.Context) error {
            ss()

            return nil
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}

func getOneTabData() (string, error) {
	// Use the local storage API to retrieve the OneTab data
	// ...
    return '',nil;
}

func ss() {
	// Set up the OAuth2 client.
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: personalAccessToken},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	// Set up the GitHub client.
	client := github.NewClient(httpClient)

	// Create the private gist.
	gist := &github.Gist{
		Public:   github.Bool(false),
		Filename: &filename,
		Files: map[github.GistFilename]github.GistFile{
			filename: github.GistFile{Content: &fileContents},
		},
	}
	result, _, err := client.Gists.Create(context.Background(), gist)
	if err != nil {
		fmt.Printf("Error creating gist: %v\n", err)
		return
	}

	fmt.Printf("Gist created successfully: %v\n", *result.HTMLURL)
}
