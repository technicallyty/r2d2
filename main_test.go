package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func TestSomething(t *testing.T) {
	// Clone the repo into memory (alternatively, use git.PlainOpen for local repos)
	repo, err := git.PlainClone("/tmp/repo", false, &git.CloneOptions{
		URL:      "https://github.com/cosmos/cosmos-sdk.git",
		Progress: nil,
	})
	if err != nil {
		log.Fatal(err)
	}

	// List all tag references
	tags, err := repo.Tags()
	if err != nil {
		log.Fatal(err)
	}

	err = tags.ForEach(func(reference *plumbing.Reference) error {
		ref := reference.Name().Short()
		fmt.Println(ref)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
