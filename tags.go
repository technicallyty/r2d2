package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"
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

	tagLineRegex = regexp.MustCompile(`^## \[(v\d+\.\d+\.\d+)\]\([^)]+\) - \d{4}-\d{2}-\d{2}$`)
)

func readRepoTags() []string {
	tags := os.Getenv(tagsEnv)
	// new lines separate tags. we will read each new line to get the tag.
	return strings.Split(tags, "\n")
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

func updateTag(pkg string, ver semver.SemVer, owner, repo string) error {
	commitSHA := os.Getenv(commitShaEnv)
	token := os.Getenv(githubTokenEnv)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	ref := &github.Reference{
		Ref:    github.Ptr("refs/tags/" + pkg + "/" + ver.String()),
		Object: &github.GitObject{SHA: github.Ptr(commitSHA), Type: github.Ptr("commit")},
	}
	_, _, err := client.Git.CreateRef(ctx, owner, repo, ref)
	if err != nil {
		return fmt.Errorf("unable to create tag: %w", err)
	}
	return nil
}
