package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

const (
	gistURL       = "https://api.github.com/gists"
	oneTabDataKey = "oneTab.groups"
)

type Gist struct {
	Description string                 `json:"description"`
	Public      bool                   `json:"public"`
	Files       map[string]GistFile    `json:"files"`
}

type GistFile struct {
	Content string `json:"content"`
}

func main() {
	// Load the OneTab data from local storage
	oneTabData, err := getOneTabData()
	if err != nil {
		log.Fatalf("Error loading OneTab data: %v", err)
	}

	// Create a new private gist with the OneTab data
	gist, err := createGist(oneTabData)
	if err != nil {
		log.Fatalf("Error creating gist: %v", err)
	}
	fmt.Println("Gist created:", gist.URL)
}

func getOneTabData() (string, error) {
	// Use the local storage API to retrieve the OneTab data
	// ...
}

func createGist(oneTabData string) (*Gist, error) {
	// Set up the HTTP client with an OAuth2 token
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	// Create the gist
	gist := &Gist{
		Description: "OneTab data",
		Public:      false,
		Files: map[string]GistFile{
			"oneTab.json": {Content: oneTabData},
		},
	}
	body, err := json.Marshal(gist)
	if err != nil {
		return nil, err
	}
	resp, err := tc.Post(gistURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the response
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Unexpected status code: %d
