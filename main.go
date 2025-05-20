package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/google/go-github/v72/github"
)

const (
	orgEnv  = "REPO_OWNER"
	repoEnv = "REPO_NAME"

	commentOnlyEnv = "R2D2_COMMENT_ONLY"
	prNumberEnv    = "R2D2_PR_NUMBER"
)

func main() {
	pkg := getPkgName()
	requestedVersion, err := getRequestedTag()
	if err != nil {
		log.Fatal(err)
	}

	latestTag, err := getLatestTagForPkg(pkg, readRepoTags())
	if err != nil && !errors.Is(err, ErrTagNotFound) {
		log.Fatal(err)
	}

	commentOnly, err := strconv.ParseBool(os.Getenv(commentOnlyEnv))
	if err != nil {
		log.Fatal(err)
	}
	if commentOnly {
		client := getGithubClient()
		prNumber, err := strconv.Atoi(os.Getenv(prNumberEnv))
		if err != nil {
			log.Fatal(err)
		}
		comment := fmt.Sprintf("This PR will update package %q from %s to %s", pkg, latestTag.String(), requestedVersion.String())
		_, _, err = client.Issues.CreateComment(
			context.Background(),
			os.Getenv(orgEnv),
			os.Getenv(repoEnv),
			prNumber,
			&github.IssueComment{Body: github.String(comment)},
		)
		if err != nil {
			log.Fatalf("failed to create PR comment: %v", err)
		}
	} else {
		err = updateTag(pkg, requestedVersion, os.Getenv(orgEnv), os.Getenv(repoEnv))
		if err != nil {
			log.Fatal(err)
		}
	}
}
