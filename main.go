package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/google/go-github/v72/github"
	"github.com/technicallyty/r2d2/semver"
)

const (
	orgEnv  = "REPO_OWNER"
	repoEnv = "REPO_NAME"

	commentOnlyEnv = "R2D2_COMMENT_ONLY"
	prNumberEnv    = "R2D2_PR_NUMBER"
)

type Package struct {
	Ver  semver.SemVer
	Name string
}

func (p Package) String() string {
	if p.Name == "" {
		return p.Ver.String()
	}
	return p.Name + "/" + p.Ver.String()
}

func main() {
	pkg := getPkgName()

	// get the requested version to tag for this package.
	requestedVersion, err := getRequestedTag()
	if err != nil {
		log.Fatal(err)
	}
	requestedPackage := Package{
		Ver:  requestedVersion,
		Name: pkg,
	}

	// get the latest tagged version of the package.
	latestTag, err := getLatestTagForPkg(pkg, readRepoTags())
	if err != nil && !errors.Is(err, ErrTagNotFound) {
		log.Fatal(err)
	}
	latestPackage := Package{
		Ver:  latestTag,
		Name: pkg,
	}

	var commentOnly bool
	commentOnlyStr := os.Getenv(commentOnlyEnv)
	if commentOnlyStr != "" {
		commentOnly, err = strconv.ParseBool(commentOnlyStr)
		if err != nil {
			log.Fatal(err)
		}
	}

	if commentOnly {
		client := getGithubClient()
		prNumber, err := strconv.Atoi(os.Getenv(prNumberEnv))
		if err != nil {
			log.Fatal(err)
		}

		comment := fmt.Sprintf("This PR will update `%s` to `%s`", latestPackage.String(), requestedPackage.String())
		_, _, err = client.Issues.CreateComment(
			context.Background(),
			os.Getenv(orgEnv),
			os.Getenv(repoEnv),
			prNumber,
			&github.IssueComment{Body: github.Ptr(comment)},
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
