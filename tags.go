package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/google/go-github/v72/github"
	"github.com/technicallyty/r2d2/semver"
	"golang.org/x/oauth2"
)

type UpdateType int

const (
	Patch UpdateType = iota
	Minor
	Major

	tagsEnv        = "COSMOS_TAGS"
	commitShaEnv   = "COMMIT_SHA"
	githubTokenEnv = "GITHUB_TOKEN"
)

var (
	ErrTagNotFound = errors.New("tag not found")
)

func readRepoTags() []string {
	tags := os.Getenv(tagsEnv)
	// new lines separate tags. we will read each new line to get the tag.
	tagsSplit := strings.Split(tags, "\n")
	slices.Sort(tagsSplit)
	return tagsSplit
}

func getLatestTagForPkg(pkg string, tags []string) (semver.SemVer, error) {
	// if pkg is empty, user wants the root version of the repo.
	// to do this, we match on the `v` prefix.
	if pkg == "" {
		pkg = "v"
	}
	latest := ""
	for _, tag := range tags {
		if strings.HasPrefix(tag, pkg) {
			latest = tag
		} else {
			if latest != "" {
				break
			}
		}
	}
	if latest != "" && latest[0] != 'v' {
		split := strings.Split(latest, "/")
		latest = split[len(split)-1]
	}
	if latest == "" {
		return semver.SemVer{}, ErrTagNotFound
	}
	return semver.Parse(latest)
}

func updateTag(pkg Package, owner, repo string) error {
	commitSHA := os.Getenv(commitShaEnv)
	ctx := context.Background()
	client := getGithubClient()
	ref := &github.Reference{
		Ref:    github.Ptr("refs/tags/" + pkg.String()),
		Object: &github.GitObject{SHA: github.Ptr(commitSHA), Type: github.Ptr("commit")},
	}
	_, _, err := client.Git.CreateRef(ctx, owner, repo, ref)
	if err != nil {
		return fmt.Errorf("unable to create tag: %w", err)
	}
	return nil
}

func getGithubClient() *github.Client {
	token := os.Getenv(githubTokenEnv)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client
}
